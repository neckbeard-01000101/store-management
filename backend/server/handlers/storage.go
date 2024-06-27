package handlers

import (
	"backend/server/database"
	"backend/server/middleware"
	"backend/server/models"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	if !emptyOrNotExist {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Collection exists and is not empty"})
		return
	}
	data := []models.Storage{
		{TypeOfProduct: "Hoodie", Quantity: 0},
		{TypeOfProduct: "Slim_T-shirt", Quantity: 0},
		{TypeOfProduct: "Hav_T-shirt", Quantity: 0},
	}

	err = initializeCollection(database.Client, "store", "storage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Collection initialized with data"})

}

func HandleAmountUpdate(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	collection := database.Client.Database("store").Collection("storage")
	var newQuantity int
	product := r.URL.Query().Get("type")
	amountStr := r.URL.Query().Get("amount")
	operation := r.URL.Query().Get("oper")
	if product == "" || amountStr == "" {
		http.Error(w, "Missing 'type' or 'amount' query parameter", http.StatusBadRequest)
		return
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		http.Error(w, "Invalid 'amount' query parameter", http.StatusBadRequest)
		return
	}

	filter := bson.D{{"type_of_product", product}}

	var result models.Storage
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if operation != "add" && operation != "sub" {
		http.Error(w, "Invalid 'oper' query parameter", http.StatusBadRequest)
		return
	}
	if operation == "add" {
		newQuantity = result.Quantity + amount
	}
	if operation == "sub" {

		newQuantity = result.Quantity - amount
	}

	update := bson.D{
		{"$set", bson.D{
			{"quantity", newQuantity},
		}},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"type_of_product": product,
		"new_quantity":    newQuantity,
	}
	json.NewEncoder(w).Encode(response)
}
