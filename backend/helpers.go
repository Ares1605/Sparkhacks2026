package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"x/llm"
)

func writeResponse(w http.ResponseWriter, status int, v any) {
	type OkResponse struct {
		Data any `json:"data"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&OkResponse{v}); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	type ErrorResponse struct {
		Err string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&ErrorResponse{msg}); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

func build_llm_tools(ctx context.Context) []llm.Tool {
	return []llm.Tool{
		llm.NewTool(
			"get_all_providers",
			"Get a list of all providers of order data on the users account.",
			func(args map[string]any) (string, error) {
				providers, err := database.GetAllProviders(ctx)
				if err != nil {
					return "", err
				}

				buf, err := json.Marshal(&providers)
				return string(buf), err
			},
		),

		llm.NewTool(
			"get_all_orders",
			"Get a list of all orders on the users account.",
			func(args map[string]any) (string, error) {
				orders, err := database.GetAllOrder(ctx)
				if err != nil {
					return "", err
				}

				buf, err := json.Marshal(&orders)
				return string(buf), err
			},
		),
	}
}
