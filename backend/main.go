package main

import (
	logs "drone/Controllers/Logs"
	websocket "drone/Websocket"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// need a rest mux router!!!

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ws)
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	// reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	r := mux.NewRouter()
	r.HandleFunc("/api/drone/logs/ssh", logs.SshLogHandler).Methods("POST")

	go func() {
		http.ListenAndServe(":8082", r)
	}()
	// http.ListenAndServe(":8082", r)

	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
