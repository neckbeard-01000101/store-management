package handlers

import (
	"backend/server/database"
	"backend/server/middleware"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func ToggleOrderState(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)

	pathSegments := strings.Split(r.URL.Path, "/")
	orderID := pathSegments[len(pathSegments)-1]
	newState := r.URL.Query().Get("newState")
	collectionName := r.URL.Query().Get("collectionName")

	objectId, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		panic(err)
	}

	collection := database.Client.Database("store").Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println(objectId)
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"order-state": newState}}
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
