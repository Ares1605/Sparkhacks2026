package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"x/db"
	"x/llm"
	"github.com/google/uuid"
)

func history_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sessionUuid := r.URL.Query().Get("session")
	if sessionUuid == "" {
		if err := json.NewEncoder(w).Encode(map[string]string{
      "error_message": "Missing 'session' UUID in search query",
    }); err != nil {
    	log.Printf("Failed to JSON encode error message with error: %v", err)
    }
		return
	}

	history, err := database.GetChatHistory(r.Context(), sessionUuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("An error occurred fetching the chat history from session (%s)", sessionUuid), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(history); err != nil {
    log.Printf("Failed to JSON encode history data with error: %v", err)
	}
}

func create_session_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string {
			"error_message": fmt.Sprintf("Expects GET request method, found %s", r.Method),
		})
		return
	}

	sessionUuid := uuid.New().String()
	err := database.InsertChatSession(r.Context(), sessionUuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string {
			"error_message": "Failed to generate chat session",
		})
		log.Printf("Error when inserting chat session with session UUID (%s): %v", sessionUuid, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {
		"session_uuid": sessionUuid,
	})
}

type askBody struct {
	Message string `json:"message"`
}
type askResponse struct {
	Response string `json:"response"`
}

func ask_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error_message": fmt.Sprintf("Expects POST request method, found %s", r.Method),
		})
		return
	}

	sessionUuid := r.URL.Query().Get("session")
	if sessionUuid == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
      "error_message": "Missing 'session' UUID in search query",
    })
		return
	}

	var body askBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error_message": "Invalid JSON body, ensure you're passing a message of type string.",
		})
		return
	}

	userMessageDate := time.Now().UTC()
	userChatMessage := db.NewChatMessage(db.UserMessage, body.Message, sessionUuid, userMessageDate)

	output, err := llm.Call([]llm.Message{
		llm.NewMessage("Talk affirmatively, and provide useful feedback", llm.SystemMessage),
		llm.NewMessage(body.Message, llm.UserMessage),
	}, []llm.Tool{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error_message": "Failed to get LLM response.",
		})
		return
	}

	serverMessageDate := time.Now().UTC()
	serverChatMessage := db.NewChatMessage(db.ServerMessage, output, sessionUuid, serverMessageDate)
	database.InsertChatMessage(r.Context(), userChatMessage)
	database.InsertChatMessage(r.Context(), serverChatMessage)

	response := askResponse{
		Response: output,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
    log.Printf("Failed to JSON encode history data with error: %v", err)
	}
}
