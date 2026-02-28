package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type NodeInfo struct {
	LastSeen time.Time
	Status   string
}

var (
	nodes   = make(map[string]*NodeInfo)
	mu      sync.Mutex
	timeout = 5 * time.Second
)

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := nodes[id]; !exists {
		nodes[id] = &NodeInfo{}
	}

	nodes[id].LastSeen = time.Now()
	nodes[id].Status = "Alive"

	fmt.Printf("[Controller] Heartbeat received from Node-%s\n", id)
}

func monitorFailures() {
	for {
		time.Sleep(2 * time.Second)

		mu.Lock()
		for id, info := range nodes {
			if time.Since(info.LastSeen) > timeout {
				if info.Status != "Suspected" {
					info.Status = "Suspected"
					fmt.Printf("[Controller] Node-%s suspected FAILED\n", id)
				}
			}
		}
		mu.Unlock()
	}
}

func main() {
	fmt.Println("[Controller] Starting...")

	http.HandleFunc("/heartbeat", heartbeatHandler)

	go monitorFailures()

	log.Fatal(http.ListenAndServe(":8080", nil))
}