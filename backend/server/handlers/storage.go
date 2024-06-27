package handlers

import (
	"backend/server/database"
	"backend/server/middleware"
	"backend/server/models"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func collectionNotExistOrEmpty(client *mongo.Client, dbName, collectionName string) (bool, error) {
	db := client.Database(dbName)
	collections, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}

	exists := false
	for _, coll := range collections {
		if coll == collectionName {
			exists = true
			break
		}
	}

	if !exists {
		return true, nil
	}

	collection := db.Collection(collectionName)
	count, err := collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func initializeCollection(client *mongo.Client, dbName, collectionName string, data []models.Storage) error {
	collection := client.Database(dbName).Collection(collectionName)
	var documents []interface{}
	for _, item := range data {
		documents = append(documents, item)
	}
	_, err := collection.InsertMany(context.TODO(), documents)
	return err
}

func HandleInitilizeStorage(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)

	emptyOrNotExist, err := collectionNotExistOrEmpty(database.Client, "store", "storage")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if emptyOrNotExist {
		data := []models.Storage{
			{TypeOfProduct: "Hoodie", Quantity: 0},
			{TypeOfProduct: "Slim T-shirt", Quantity: 0},
			{TypeOfProduct: "Hav T-shirt", Quantity: 0},
		}

		err := initializeCollection(database.Client, "store", "storage", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Collection initialized with data"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Collection exists and is not empty"})
	}
}
