package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		log.Fatal("NODE_ID not set")
	}

	fmt.Printf("[Node-%s] Starting...\n", nodeID)

	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		fmt.Printf("[Node-%s] Heartbeat sent at %s\n", nodeID, time.Now().Format(time.RFC3339))
	}
}