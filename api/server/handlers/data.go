package handlers

import (
	"api/server/database"
	"api/server/middleware"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleGetData(w http.ResponseWriter, r *http.Request) {
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

func HandleOrderState(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
	newState := r.URL.Query().Get("newState")
	orderID := r.URL.Query().Get("orderId")
	collectionName := r.URL.Query().Get("collectionName")

	objectId, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(w, "Invalid order ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	collection := database.Client.Database("store").Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var order bson.M
	if err := collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&order); err != nil {
		http.Error(w, "Error fetching order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if order["order-state"] == newState {
		http.Error(w, "Error: The new state is the same as the current state, so it was not updated", http.StatusBadRequest)
		return
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"order-state": newState}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error updating order state: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "No order found with the given ID", http.StatusNotFound)
		return
	}

	fmt.Printf("Matched %d documents and modified %d documents\n", result.MatchedCount, result.ModifiedCount)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"message": "Order state updated successfully"})
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	pathSegments := strings.Split(r.URL.Path, "/")
	collectionName := pathSegments[len(pathSegments)-2]
	documentID := pathSegments[len(pathSegments)-1]

	objectId, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		http.Error(w, "Invalid document ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(objectId)
	collection := database.Client.Database("store").Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		http.Error(w, "Error deleting document: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "No document found with the given ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"message": "Document deleted successfully"})
}

func ToggleIsDone(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := pathSegments[2]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	collectionName := r.URL.Query().Get("collectionName")
    
	collection := database.Client.Database("store").Collection(collectionName)

	var order bson.M
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&order)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	orderNumStr := r.URL.Query().Get("order-num")
	orderNum, err := strconv.Atoi(orderNumStr)
	if err != nil {
		http.Error(w, "Invalid order number format", http.StatusBadRequest)
		return
	}
	newIsDoneStatusStr := r.URL.Query().Get("new-state")
    
    newIsDoneStatus, err := strconv.ParseBool(newIsDoneStatusStr)

    if err != nil {
		http.Error(w, "Error converting the state into a bool", http.StatusBadRequest)
        return 
    }

	updateResult, err := collection.UpdateMany(
		context.TODO(),
		bson.M{"order-num": orderNum},
		bson.M{"$set": bson.M{"is-done": newIsDoneStatus}},
	)
	if err != nil {
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Toggled 'is-done' status for %d documents\n", updateResult.ModifiedCount)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"message": "Order 'is-done' status toggled successfully"})
}
