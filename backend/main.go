package main

import (
	"backend/server/handlers"
	"log"
	"net/http"
)

func main() {
	const PORT string = ":8000"

	router := http.NewServeMux()
	router.HandleFunc("/send", handlers.HandlePostForm)
	router.HandleFunc("/get/", handlers.HandleSendingData)
	router.HandleFunc("/collections", handlers.HandleGetCollections)
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Starting server on port %s", PORT)
	log.Fatal(server.ListenAndServe())
}
