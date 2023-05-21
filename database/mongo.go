package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMgoConnection() *mongo.Client {
	// var (
	// 	MGO_HOST     = config.ENV.MgoHost
	// 	MGO_PASSWORD = config.ENV.MgoPassword
	// )
	// uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.fuxzenu.mongodb.net/?retryWrites=true&w=majority")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI("mongodb+srv://naufalhakim:OOnSx5j0Q0qNeB5o@cluster0.fuxzenu.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected Successfully")
	return client
}

func MgoCollection(coll string, client *mongo.Client) *mongo.Collection {
	return client.Database("intelitalk").Collection(coll)
}
