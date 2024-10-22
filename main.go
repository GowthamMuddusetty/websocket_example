package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket endpoint
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade the GET request to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to upgrade to WebSocket: %v", err)
	}
	defer ws.Close()

	// Continuously read messages from the WebSocket
	for {
		// Read message
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		// Print received message
		fmt.Printf("Received: %s\n", msg)

		// Echo message back to the client
		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}

func main() {
	// Configure HTTP server
	http.HandleFunc("/ws", handleConnections)

	// Start WebSocket server
	log.Println("WebSocket server started at :8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
