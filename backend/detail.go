package main

import (
	"encoding/json"
	"net/http"
)

func detials_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	providers, err := database.GetAllProviders(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{Err: err.Error()})
		return
	}

	var provider_details = []ProviderDetails{}
	for _, provider := range providers {
		provider_details = append(provider_details, ProviderDetails{
			Name:       provider.Name,
			LastSynced: provider.LastSync.String(),
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&OkResponse{Data: provider_details})
}
