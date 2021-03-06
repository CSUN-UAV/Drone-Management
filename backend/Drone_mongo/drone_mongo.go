package drone_mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	drone_config "github.com/CSUN-UAV/Drone-Management/backend/Drone_config"
	models "github.com/CSUN-UAV/Drone-Management/backend/Models"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type key string

const (
	HostKey     = key("hostKey")
	UsernameKey = key("usernameKey")
	PasswordKey = key("passwordKey")
	DatabaseKey = key("databaseKey")
)

var ctx context.Context
var client *mongo.Client

func init() {
	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, HostKey, drone_config.MongoHost)
	ctx = context.WithValue(ctx, UsernameKey, drone_config.MongoUser)
	ctx = context.WithValue(ctx, PasswordKey, drone_config.MongoPassword)
	ctx = context.WithValue(ctx, DatabaseKey, drone_config.MongoDb)

	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value(UsernameKey).(string),
		ctx.Value(PasswordKey).(string),
		ctx.Value(HostKey).(string),
		ctx.Value(DatabaseKey).(string),
	)

	clientOptions := options.Client().ApplyURI(uri)

	var err error

	client, err = mongo.Connect(ctx, clientOptions)

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mongo Connected")
	}
}

type GetDocumentsTask struct {
	idx			int
	collection string
	w          http.ResponseWriter
	wg         *sync.WaitGroup
}

func NewGetDocumentsTask(idx int,collection string, w http.ResponseWriter, wg *sync.WaitGroup) *GetDocumentsTask {
	return &GetDocumentsTask{idx, collection, w, wg}
}

func (t *GetDocumentsTask) Perform() {
	defer t.wg.Done()
	switch t.collection {
	case "logs":
		// var xdoc map[string]interface{}
		var xdoc []*models.DroneCommandLogs
		collection := client.Database("logs").Collection("main")
		findOptions := options.Find()
		findOptions.SetSkip(int64(t.idx))
		findOptions.SetSort(bson.D{{"_id", -1}})
		findOptions.SetLimit(20)
		// filter := bson.D{{"_id", bson.D{{"&lt", 2}}}}

		cur, err := collection.Find(ctx, bson.D{}, findOptions)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			for cur.Next(ctx) {
				// var doc map[string]interface{}
				var log models.DroneCommandLogs
				err := cur.Decode(&log)
				if err != nil {
					fmt.Println("error")
				}
				xdoc = append(xdoc, &log)
			}
			for i, j := 0, len(xdoc) -1; i < j; i, j = i+1, j-1 {
				xdoc[i], xdoc[j] = xdoc[j], xdoc[i]
			}
			// cur.Decode(&xdoc)
			json.NewEncoder(t.w).Encode(xdoc)
		}
		break
	}
}

type AddLogToDbTask struct {
	// Data		string
	// Data		interface{}
	Data 		json.RawMessage
	ws 			*websocket.Conn
}

func NewAddLogToDbTask(data json.RawMessage, ws *websocket.Conn) *AddLogToDbTask {
// func NewAddLogToDbTask(data interface{}, ws *websocket.Conn) *AddLogToDbTask {
// func NewAddLogToDbTask(data string, ws *websocket.Conn) *AddLogToDbTask{
	return &AddLogToDbTask{data, ws}
}

func (t *AddLogToDbTask) Perform() {
	log := models.DroneCommandLogs{}

	err := json.Unmarshal(t.Data, &log);

	// fmt.Println("init3")
	// fmt.Println(log)
	// err := json.Unmarshal([]byte(t.Data), &log);

	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println("init4")
	collection := client.Database("logs").Collection("main")
	collection.InsertOne(ctx, log)
}