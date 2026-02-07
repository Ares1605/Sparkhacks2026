package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

func resync_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := exec.CommandContext(r.Context(), python_executable, "sync-amazon-data.py")
	cmd.Dir = "scripts"

	if _, err := cmd.CombinedOutput(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{Err: "Resync Failed"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&OkResponse{Data: "Resync completed"})
}
