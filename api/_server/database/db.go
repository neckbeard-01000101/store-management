package database

import (
	"context"
	"log"
	"time"

	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	uri, err := env.MustGet("MONGODB_URI")

	if err != nil {
		log.Fatal(err)
	}
	clientOptions := options.Client().ApplyURI(uri)
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
}
