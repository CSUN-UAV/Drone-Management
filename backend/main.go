package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"flag"
	"strconv"

	drone_mongo "github.com/CSUN-UAV/Drone-Management/backend/Drone_mongo"

	drone_asynq "github.com/CSUN-UAV/Drone-Management/backend/Drone_asynq"
	websocket "github.com/CSUN-UAV/Drone-Management/backend/Websocket"
	"github.com/gorilla/mux"
)

var addr = flag.String("addr", "0.0.0.0:1200", "http service address")


func handleWs(w http.ResponseWriter, r *http.Request) {
	swp := r.Header.Get("Sec-Websocket-Protocol");
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	fmt.Print(ws)
	fmt.Print(swp)
}

func handleSSHLogs(w http.ResponseWriter, r *http.Request) {
	idx := mux.Vars(r)["idx"]
	index, err := strconv.Atoi(idx)
	if err != nil {
		fmt.Println(err)
		return
	}
	logs := mux.Vars(r)["log_type"]
	params := mux.Vars(r)
	fmt.Println(params)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var wg sync.WaitGroup
	wg.Add(1)

	// drone_mongo.NewGetDocumentsTask(params["idx"], params["log_type"], w, &wg)
	tobj := drone_mongo.NewGetDocumentsTask(index, logs, w, &wg) //notice the pointer to the wait group
	drone_asynq.TaskQueue <- tobj
	wg.Wait()
	fmt.Println("Wait Group Finished Success...")
}

func main() {
	drone_asynq.StartTaskDispatcher(9)
	flag.Parse()
	log.SetFlags(0)	

	r := mux.NewRouter()

	// ws
	r.HandleFunc("/ws", handleWs)

	// rest
	r.HandleFunc("/api/drone/logs/ssh/{idx:[0-9]+}/{log_type}", handleSSHLogs)

	http.ListenAndServeTLS(*addr, "/etc/letsencrypt/live/csunuav.me/cert.pem", "/etc/letsencrypt/live/csunuav.me/privkey.pem", r)
}
