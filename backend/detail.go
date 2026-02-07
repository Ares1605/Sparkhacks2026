package main

import (
	"encoding/json"
	"net/http"
)

func detials_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	provider, exists, err := database.GetProviderStatusByID(r.Context(), 1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{Err: err.Error()})
		return
	}

	response := ProviderDetailsResponse{
		Amazon: AmazonProviderDetails{
			LoggedIn:   exists,
			LastSynced: nil,
			Username:   nil,
		},
	}
	if exists {
		response.Amazon.LastSynced = provider.LastSync
		response.Amazon.Username = provider.Username
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}
