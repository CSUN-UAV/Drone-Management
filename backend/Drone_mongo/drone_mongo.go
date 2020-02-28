package drone_mongo

import (
	"context"
	"fmt"

	drone_config "github.com/CSUN-UAV/Drone-Management/backend/Drone_config"
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
