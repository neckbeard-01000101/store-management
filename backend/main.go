package main

import (
	"backend/server/handlers"
	"log"
	"net/http"
)

func main() {
	const PORT string = ":8000"

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Listening to incoming requests..."))

	})
	router.HandleFunc("/send", handlers.HandlePostForm)
	router.HandleFunc("/get/", handlers.HandleGetData)
	router.HandleFunc("/collections", handlers.HandleGetCollections)
	router.HandleFunc("/toggleOrderState/", handlers.ToggleOrderState)
	router.HandleFunc("POST /initialize", handlers.HandleInitilizeStorage)
	router.HandleFunc("PATCH /add/", handlers.HandleAdd)
	router.HandleFunc("PATCH /sub/", handlers.HandleSub)
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Server is running on http://localhost%s", PORT)
	log.Fatal(server.ListenAndServe())
}
