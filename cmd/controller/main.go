package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	pb "chaos-heal/proto" // adjust if your module path differs

	"google.golang.org/grpc"
)

type NodeStatus string

const (
	Alive     NodeStatus = "ALIVE"
	Suspected NodeStatus = "SUSPECTED"
	Dead      NodeStatus = "DEAD"
)

type NodeInfo struct {
	LastHeartbeat time.Time
	Status        NodeStatus
}

type server struct {
	pb.UnimplementedHeartbeatServiceServer
}

var (
	nodes = make(map[string]*NodeInfo)
	mu    sync.Mutex
)

func (s *server) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	nodeID := req.NodeId

	info, exists := nodes[nodeID]
	if !exists {
		info = &NodeInfo{}
		nodes[nodeID] = info
		log.Printf("[Controller] New node registered: %s\n", nodeID)
	}

	previousStatus := info.Status

	info.LastHeartbeat = time.Now()
	info.Status = Alive

	if previousStatus != Alive {
		log.Printf("[Controller] %s recovered -> ALIVE\n", nodeID)
	} else {
		log.Printf("[Controller] Heartbeat received from %s\n", nodeID)
	}

	return &pb.HeartbeatResponse{Ack: true}, nil
}

func startFailureDetector() {
	go func() {
		for {
			time.Sleep(2 * time.Second)

			mu.Lock()
			for nodeID, info := range nodes {
				elapsed := time.Since(info.LastHeartbeat)

				switch {
				case elapsed > 10*time.Second && info.Status != Dead:
					info.Status = Dead
					log.Printf("[Controller] %s marked DEAD\n", nodeID)

				case elapsed > 5*time.Second && info.Status == Alive:
					info.Status = Suspected
					log.Printf("[Controller] %s marked SUSPECTED\n", nodeID)
				}
			}
			mu.Unlock()
		}
	}()
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHeartbeatServiceServer(grpcServer, &server{})

	startFailureDetector()

	log.Println("[Controller] Server started on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}