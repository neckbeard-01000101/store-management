package handlers

import (
	"backend/server/database"
	"backend/server/middleware"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"time"
)

func HandleSendingData(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
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
	collection := database.Client.Database("store").Collection(collectionName)

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

func ToggleOrderState(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)

	pathSegments := strings.Split(r.URL.Path, "/")
	orderID := pathSegments[len(pathSegments)-1]
	newState := r.URL.Query().Get("newState")
	collectionName := r.URL.Query().Get("collectionName")
	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	if newState == "" {
		http.Error(w, "New state is required", http.StatusBadRequest)
		return
	}

	objectId, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		fmt.Println("Invalid id")
	}

	collection := database.Client.Database("store").Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"order-state": newState}}
	println(orderID)
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error updating order state: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Matched %d documents and modified %d documents\n", result.MatchedCount, result.ModifiedCount)

	if result.MatchedCount == 0 {
		http.Error(w, "No order found with the given ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"message": "Order state updated successfully"})
}