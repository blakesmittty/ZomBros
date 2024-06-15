package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // allowing all requests
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%v sent: %v\n", ws.RemoteAddr(), msg)
	}
}

func main() {
	fmt.Print("save the import")
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleConnections)

	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string {"http://localhost:3000"}))(mux)

	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", corsHandler)
}