package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DATABASE_URL")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB!")
}
