package main

import (
	"net/http"
)

func detials_handler(w http.ResponseWriter, r *http.Request) {
	provider, exists, err := database.GetProviderStatusByID(r.Context(), 1)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type AmazonProviderDetails struct {
		LoggedIn   bool    `json:"logged_in"`
		LastSynced *string `json:"last_synced"`
		Username   *string `json:"username"`
	}

	type ProviderDetailsResponse struct {
		Amazon AmazonProviderDetails `json:"amazon"`
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

	writeResponse(w, http.StatusOK, response)
}
