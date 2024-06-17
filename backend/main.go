package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
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
	OrderNum      int `json:"order-num"`
	OrderState    string `json:"order-state"`
	CustomerName  string `json:"customer-name"`
	CustomerCity  string `json:"customer-city"`
	CustomerPhone int `json:"customer-phone"`
	SellerName    string `json:"seller-name"`
	TotalCost     int `json:"total-cost"`
	SellerProfit  int `json:"seller-profit"`
	DeliveryFee   int `json:"delivery-fee"`
	Size          string `json:"size"`
	Color         string `json:"color"`
	ClothesType   string `json:"clothes-type"`
	ProductCost   int `json:"cost-of-product"`
	TotalProfit int
	SentDate string
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
	formData.TotalProfit = formData.TotalCost - formData.SellerProfit - formData.DeliveryFee - formData.ProductCost
	now := time.Now()
    formData.SentDate = now.Format("02-01")
	collectionName := now.Format("01-2006")
	collection := client.Database("store").Collection(collectionName)
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
		"total-profit":    formData.TotalProfit,
		"sent-date":       formData.SentDate,
	})
	if err != nil {
		http.Error(w, "Error saving form data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Inserted document ID: %s", result.InsertedID)
}


func handleSendingData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    collectionName := strings.TrimPrefix(r.URL.Path, "/get/")
	if collectionName == "" {
        http.Error(w, "Collection name is required", http.StatusBadRequest)
        return
    }
    collection := client.Database("store").Collection(collectionName)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Error fetching documents: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var documents []bson.M
    if err = cursor.All(ctx, &documents); err != nil {
        http.Error(w, "Error decoding documents: "+err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(documents)
}

func handleGetCollections(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    storeDB := client.Database("store")
    collections, err := storeDB.ListCollectionNames(ctx, bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(collections)
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
	router.HandleFunc("/get/", handleSendingData)
	router.HandleFunc("/collections", handleGetCollections)
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Starting server on port %s", PORT)
	log.Fatal(server.ListenAndServe())
}
