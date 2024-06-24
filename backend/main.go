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
	router.HandleFunc("/get/", handlers.HandleGetData)
	router.HandleFunc("/collections", handlers.HandleGetCollections)
	router.HandleFunc("/toggleOrderState/", handlers.ToggleOrderState)
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Server is running on http://localhost%s", PORT)
	log.Fatal(server.ListenAndServe())
}
