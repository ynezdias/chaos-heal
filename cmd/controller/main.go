package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[Controller] Starting heartbeat monitor...")

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		fmt.Printf("[Controller] Checking heartbeats at %s\n", time.Now().Format(time.RFC3339))
	}
}