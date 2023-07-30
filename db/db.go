package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDb()(*mongo.Collection, context.Context, error) {
	var collection *mongo.Collection
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("onesignal").Collection("users")
	fmt.Println(collection)

	return collection,ctx,err
}
