package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	nodeID := os.Getenv("NODE_ID")
	controllerURL := os.Getenv("CONTROLLER_URL")

	if nodeID == "" || controllerURL == "" {
		log.Fatal("NODE_ID or CONTROLLER_URL not set")
	}

	fmt.Printf("[Node-%s] Starting...\n", nodeID)

	ticker := time.NewTicker(2 * time.Second)

	for range ticker.C {
		url := fmt.Sprintf("%s/heartbeat?id=%s", controllerURL, nodeID)
		_, err := http.Get(url)
		if err != nil {
			fmt.Printf("[Node-%s] Failed to send heartbeat: %v\n", nodeID, err)
			continue
		}
		fmt.Printf("[Node-%s] Heartbeat sent\n", nodeID)
	}
}