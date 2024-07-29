package main

import (
	"api/server/database"
	"api/server/handlers"
	"github.com/gofor-little/env"
	"log"
	"net/http"
)

func main() {
	err := env.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database.Init()
	const PORT string = ":8000"

	router := http.NewServeMux()
	router.HandleFunc("/send", handlers.HandlePostForm)
	router.HandleFunc("/get/", handlers.HandleGetData)
	router.HandleFunc("/collections", handlers.HandleGetCollections)
	router.HandleFunc("/updateState/", handlers.HandleOrderState)
	router.HandleFunc("POST /updateAmount/", handlers.HandleAmountUpdate)
	router.HandleFunc("/add", handlers.HandleAddToStorage)
	router.HandleFunc("/deleteDocument/", handlers.DeleteDocument)
	router.HandleFunc("/profit/", handlers.HandleProfits)
	router.HandleFunc("/is-done/", handlers.ToggleIsDone)
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Server is running on http://localhost%s", PORT)
	log.Fatal(server.ListenAndServe())
}
