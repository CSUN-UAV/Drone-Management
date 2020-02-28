package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	drone_mongo "github.com/CSUN-UAV/Drone-Management/backend/Drone_mongo"

	websocket "github.com/CSUN-UAV/Drone-Management/backend/Websocket"
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

func handleUI(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var wg sync.WaitGroup
	wg.Add(1)

	drone_mongo.NewGetDocumentsTask("logs", w, &wg)
	tobj := drone_mongo.NewGetDocumentsTask("logs", w, &wg) //notice the pointer to the wait group

	fmt.Println(tobj)
	wg.Wait()
	fmt.Println("Wait Group Finished Success...")
}

func main() {
	fmt.Println("Chat App v0.01")
	r := mux.NewRouter()
	r.HandleFunc("/api/drone/logs/ssh", handleUI)

	// go func() {
	// 	http.ListenAndServe(":8082", r)
	// }()
	// http.ListenAndServe(":8082", r)
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
