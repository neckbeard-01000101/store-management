package main

import (
	"api/server/handlers"
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
	// router.HandleFunc("POST /initialize", handlers.HandleInitilizeStorage)
	router.HandleFunc("POST /updateAmount/", handlers.HandleAmountUpdate)
	router.HandleFunc("/add", handlers.HandleAddToStorage)
	router.HandleFunc("/deleteDocument/", handlers.DeleteDocument)
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Server is running on http://localhost%s", PORT)
	log.Fatal(server.ListenAndServe())
}
