package handlers

import (
	"backend/server/database"
	"backend/server/middleware"
	"backend/server/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandlePostForm(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var formData models.FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	formData.TotalCost *= formData.NumOfPieces
	formData.TotalProfit = formData.TotalCost - formData.SellerProfit - formData.DeliveryFee - formData.ProductCost
	now := time.Now()
	formData.SentDate = now.Format("02-01")
	collectionName := now.Format("01-2006")
	formData.Month = collectionName
	collection := database.Client.Database("store").Collection(collectionName)
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
		"collection-name": formData.Month,
		"pieces-num":      formData.NumOfPieces,
	})
	if err != nil {
		http.Error(w, "Error saving form data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	monthsProfitsCollection := database.Client.Database("store").Collection("months-profits")

	filter := bson.M{"month": collectionName}
	update := bson.M{"$inc": bson.M{"profit": formData.TotalProfit}}
	opts := options.Update().SetUpsert(true)

	_, err = monthsProfitsCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		http.Error(w, "Failed to update month's profit: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Inserted document ID: %s", result.InsertedID)
}
