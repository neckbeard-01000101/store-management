package handlers

import (
	"api/server/database"
	"api/server/middleware"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func HandleProfits(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)

	month := r.URL.Query().Get("month")
	state := r.URL.Query().Get("state")
	cost, _ := strconv.Atoi(r.URL.Query().Get("cost"))
	profit, _ := strconv.Atoi(r.URL.Query().Get("profit"))
	numOfPieces, _ := strconv.Atoi(r.URL.Query().Get("pieces"))

	collection := database.Client.Database("store").Collection("months-profits")

	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"month": month}).Decode(&result)
	if err != nil {
		_, err := collection.InsertOne(context.TODO(), bson.M{
			"month":           month,
			"returned":        0,
			"profits":         0,
			"pieces_returned": 0,
			"pieces_sold":     0,
		})
		if err != nil {
			http.Error(w, "Failed to insert document", http.StatusInternalServerError)
			return
		}
	}

	update := bson.M{}
	switch state {
	case "delivered":
		update = bson.M{"$inc": bson.M{"profits": profit, "pieces_sold": numOfPieces}}
	case "returned":
		update = bson.M{"$inc": bson.M{"returned": cost, "pieces_returned": numOfPieces}}
	case "pending":
		return
	default:
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"month": month}, update)
	if err != nil {
		http.Error(w, "Failed to update document", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Document updated successfully")
}
