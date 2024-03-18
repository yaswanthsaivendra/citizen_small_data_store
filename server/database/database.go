package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	Citizens *mongo.Collection
)

func InitDB(uri string) error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	localclient, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		return err
	}
	client = localclient

	Citizens = client.Database("small_data_store").Collection("citizens")

	err = client.Ping(context.Background(), nil)

	return err
}

func Close() error {
	return client.Disconnect(context.Background())
}
