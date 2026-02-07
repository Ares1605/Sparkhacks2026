package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type setAmazonCredentialsBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func set_amazon_credentials_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Err: "Expected POST request method.",
		})
		return
	}

	var body setAmazonCredentialsBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Err: "Invalid JSON body. Expected username and password.",
		})
		return
	}

	body.Username = strings.TrimSpace(body.Username)
	if body.Username == "" || strings.TrimSpace(body.Password) == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Err: "username and password are required.",
		})
		return
	}

	if err := database.UpsertProviderCredentials(
		r.Context(),
		amazonProviderRowID,
		amazonProviderName,
		body.Username,
		body.Password,
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			Err: "Failed to store Amazon credentials.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&MessageResponse{
		Message: "Successfully set amazon credentials!",
	})
}
