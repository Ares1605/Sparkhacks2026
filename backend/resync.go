package main

import (
	"encoding/json"
	"net/http"
)

func resync_handler(w http.ResponseWriter, r *http.Request) {
	response := Response{"This endpoint does nothing."}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		panic(err)
	}
}
