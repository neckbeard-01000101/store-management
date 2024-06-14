package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type data struct {
	orderNum      string
	orderState    string
	customerName  string
	customerCity  string
	customerPhone string
	sellerName    string
	totalCost     int
	sellerProfit  int
}

func main() {
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	const PORT string = ":8000"
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fmt.Println("A Request has been made on localhost:8000/send")
	})
	fmt.Printf("Starting server on port %s\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Println("Failed to start the server!")
		return
	}

}
