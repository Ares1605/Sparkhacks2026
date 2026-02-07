package main

import (
	"encoding/json"
	"net/http"
)

func detials_handler(w http.ResponseWriter, r *http.Request) {
	providers, err := database.GetAllProviders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var provider_details = []ProviderDetails{}
	for _, provider := range providers {
		provider_details = append(provider_details, ProviderDetails{
			Name:       provider.Name,
			LastSynced: provider.LastSync.String(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&provider_details); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
