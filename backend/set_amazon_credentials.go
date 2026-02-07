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
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Expected POST request method.")
		return
	}

	var body setAmazonCredentialsBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest,  "Invalid JSON body. Expected username and password.")
		return
	}

	body.Username = strings.TrimSpace(body.Username)
	if body.Username == "" || strings.TrimSpace(body.Password) == "" {
		writeError(w, http.StatusBadRequest,  "username and password are required.")
		return
	}

	if err := database.UpsertProviderCredentials(
		r.Context(),
		amazonProviderRowID,
		amazonProviderName,
		body.Username,
		body.Password,
	); err != nil {

		writeError(w, http.StatusInternalServerError, "Failed to store Amazon credentials.")
		return
	}

	writeResponse(w, http.StatusOK,  "Successfully set amazon credentials!")
}
