package logs

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/CSUN-UAV/Drone-Management/backend/Models"
	"github.com/google/uuid"
)

func SshLogHandler(resp http.ResponseWriter, req *http.Request) {
	var sshLog models.DroneCommandLogs
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&sshLog)

	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}
	sshLog.UUID = id.String()
	if err != nil {
		fmt.Println("fml")
		return
	}
	json.NewEncoder(resp).Encode(sshLog)
}
