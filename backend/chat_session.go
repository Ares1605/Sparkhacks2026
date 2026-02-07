package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"x/db"
	"x/llm"
)

func history_handler(w http.ResponseWriter, r *http.Request) {
	sessionUuid := r.URL.Query().Get("session")
	if sessionUuid == "" {
		writeError(w, http.StatusBadRequest, "Missing 'session' UUID in search query")
		return
	}

	history, err := database.GetChatHistory(r.Context(), sessionUuid)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("An error occurred fetching the chat history from session (%s)", sessionUuid))
		return
	}

	writeResponse(w, http.StatusOK, history)
}

func create_session_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Expects GET request method, found %s", r.Method))
		return
	}

	sessionUuid := uuid.New().String()
	err := database.InsertChatSession(r.Context(), sessionUuid)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate chat session"))
		return
	}

	writeResponse(w, http.StatusOK, sessionUuid)
}

func ask_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Expects POST request method, found %s", r.Method))
		return
	}

	sessionUuid := r.URL.Query().Get("session")
	if sessionUuid == "" {
		writeError(w, http.StatusBadRequest, "Missing 'session' UUID in search query")
		return
	}

	var body struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "Failed to parse request body. Ensure data is json")
		return
	}

	if body.Message == "" {
		writeError(w, http.StatusBadRequest, "Missing property 'message'")
		return
	}

	userMessageDate := time.Now().UTC()
	userChatMessage := db.NewChatMessage(db.UserMessage, body.Message, sessionUuid, userMessageDate)

	orders, err := database.GetAllOrder(r.Context())
	if err != nil {
		panic(":(")
	}

	buf, err := json.Marshal(&orders)
	ordersStringified := string(buf)

	output, err := llm.Call([]llm.Message{
		llm.NewMessage(fmt.Sprintf(
			`You are Precog, a proactive shopping assistant. Talk affirmatively, and provide useful feedback.
			
			Included below is a list of the users previous order history in a structural format. You must use this order
			history to assist yourself in responding best to the users needs. For example, if the user is asking about what
			portafilter they should buy, and you see the particular espresso machine in their order history, ensure your
			recommendations are compatible with the particular espresso machine they had bought. Point out previous orders from their
			order history when relevant. Users love that. Here is their order history:
			%s`, ordersStringified), llm.SystemMessage),
		llm.NewMessage(body.Message, llm.UserMessage),
	}, []llm.Tool{})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get LLM response.")
		return
	}

	serverMessageDate := time.Now().UTC()
	serverChatMessage := db.NewChatMessage(db.ServerMessage, output, sessionUuid, serverMessageDate)
	database.InsertChatMessage(r.Context(), userChatMessage)
	database.InsertChatMessage(r.Context(), serverChatMessage)

	w.WriteHeader(http.StatusOK)
	writeResponse(w, http.StatusOK, output)
}
