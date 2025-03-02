package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PasinduYeshan/go-sample-oauth/internal/models"
)

func GetAllServices(w http.ResponseWriter, r *http.Request) {

	services := []models.Service{
		{
			ID:          1,
			Name:        "Car Oil Change",
			Description: "Complete engine oil change and filter replacement",
			Price:       59.99,
			Category:    "Car Repair",
		},
		{
			ID:          2,
			Name:        "Bike Tune-up",
			Description: "Full bicycle tune-up including gear and brake adjustment",
			Price:       45.00,
			Category:    "Bike Repair",
		},
		{
			ID:          3,
			Name:        "Car Brake Repair",
			Description: "Brake pad replacement and system check",
			Price:       129.99,
			Category:    "Car Repair",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}
