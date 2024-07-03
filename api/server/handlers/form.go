package handlers

import (
	"api/server/database"
	"api/server/middleware"
	"api/server/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	formData.TotalProfit = formData.TotalCost - formData.SellerProfit - formData.DeliveryFee - formData.ProductCost
	now := time.Now()
	formData.SentDate = now.Format("02-01")
	collectionName := now.Format("01-2006")
	formData.Month = collectionName
	collection := database.Client.Database("store").Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var existingDoc bson.M
	err = collection.FindOne(ctx, bson.M{"order-num": formData.OrderNum}).Decode(&existingDoc)
	if err == nil {
		formData.DeliveryFee = 0
		formData.TotalCost = 0
		formData.SellerProfit = 0
		formData.ProductCost = 0
		formData.TotalProfit = 0
	}
	formData.IsDone = false
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
		"is-done":         formData.IsDone,
	})
	if err != nil {
		http.Error(w, "Error saving form data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Inserted document ID: %s", result.InsertedID)
}
