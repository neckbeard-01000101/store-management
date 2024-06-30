package handlers

import (
	"api/server/database"
	"api/server/middleware"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func HandleGetCollections(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	storeDB := database.Client.Database("store")
	collections, err := storeDB.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(collections)
}
