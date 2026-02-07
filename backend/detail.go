package main

import (
	"net/http"
)

func detials_handler(w http.ResponseWriter, r *http.Request) {
	providers, err := database.GetAllProviders(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type ProviderDetails struct {
		Name       string `json:"name"`
		LastSynced string `json:"last_synced"`
	}

	var details = []ProviderDetails{}

	for _, provider := range providers {
		details = append(details, ProviderDetails{
			Name:       provider.Name,
			LastSynced: provider.LastSync.String(),
		})
	}

	writeResponse(w, http.StatusOK, details)
}
