package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type FormData struct {
	OrderNum      string `json:"order-num"`
	OrderState    string `json:"order-state"`
	CustomerName  string `json:"customer-name"`
	CustomerCity  string `json:"customer-city"`
	CustomerPhone string `json:"customer-phone"`
	SellerName    string `json:"seller-name"`
	TotalCost     string `json:"total-cost"`
	SellerProfit  string `json:"seller-profit"`
	DeliveryFee   string `json:"delivery-fee"`
	Size          string `json:"size"`
	Color         string `json:"color"`
	ClothesType   string `json:"clothes-type"`
	ProductCost   string `json:"cost-of-product"`
}

func handlePostForm(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	/*
	   add total profit => total cost - seler profit - delivery fee - cost of product

	*/
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var formData FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("store").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"order-num":       formData.OrderNum,
		"order-state":     formData.OrderState,
		"customer-name":   formData.CustomerName,
		"customer-city":   formData.CustomerCity,
		"customer-phone":  formData.CustomerPhone,
		"seller-name":     formData.SellerName,
		"total-cost":      formData.TotalCost,
		"seller-profit":   formData.SellerProfit,
		"delivery-fee":    formData.DeliveryFee,
		"size":            formData.Size,
		"color":           formData.Color,
		"clothes-type":    formData.ClothesType,
		"cost-of-product": formData.ProductCost,
	})
	if err != nil {
		http.Error(w, "Error saving form data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Inserted document ID: %s", result.InsertedID)
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

	router := http.NewServeMux()
	router.HandleFunc("/send", handlePostForm)

	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Starting server on port %s", PORT)
	log.Fatal(server.ListenAndServe())
}
