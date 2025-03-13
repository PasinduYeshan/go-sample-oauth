package main

import (
	"github.com/PasinduYeshan/go-sample-oauth/internal/ads"
	"github.com/PasinduYeshan/go-sample-oauth/internal/merchantservices"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ads", ads.GetAllAds)

	http.HandleFunc("/services", merchantservices.GetAllServices)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
