package ads

import (
	"encoding/json"
	"net/http"
)

func GetAllAds(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(ads)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Sample data.
var ads = []Ad{
	{ID: 1, Title: "Ad 1", Description: "Description for Ad 1", Price: 100.00},
	{ID: 2, Title: "Ad 2", Description: "Description for Ad 2", Price: 200.00},
	{ID: 3, Title: "Ad 3", Description: "Description for Ad 3", Price: 300.00},
}
