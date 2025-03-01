package main

import (
	"log"
	"net/http"

	"github.com/PasinduYeshan/go-sample-oauth/internal/handlers"
)

func main() {
	http.HandleFunc("/ads", handlers.GetAllAds)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
