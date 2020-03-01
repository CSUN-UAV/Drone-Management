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

func handleWs(w http.ResponseWriter, r *http.Request) {
	// swp := r.Header.Get("Sec-Websocket-Protocol");
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	fmt.Print(ws)
}

func handleSSHLogs(w http.ResponseWriter, r *http.Request) {
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

	// ws
	r.HandleFunc("/ws", handleWs)

	// rest
	r.HandleFunc("/api/drone/logs/ssh", handleSSHLogs)

	http.ListenAndServe(":8080", nil)
}
