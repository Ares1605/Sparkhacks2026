package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func resync_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	scriptPath, scriptDir, err := findSyncScript()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	pythonPath, err := exec.LookPath("python3")
	if err != nil {
		pythonPath, err = exec.LookPath("python")
	}
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Python executable not found in PATH.")
		return
	}

	if err := os.MkdirAll(filepath.Join(scriptDir, "history"), 0o755); err != nil {
		writeJSON(w, http.StatusInternalServerError, "Unable to prepare history directory: "+err.Error())
		return
	}

	cmd := exec.CommandContext(r.Context(), pythonPath, scriptPath)
	cmd.Dir = scriptDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		msg := "Resync failed: " + err.Error()
		trimmed := strings.TrimSpace(string(output))
		if trimmed != "" {
			msg += "; output: " + trimOutput(trimmed, 2000)
		}
		writeJSON(w, http.StatusInternalServerError, msg)
		return
	}

	writeJSON(w, http.StatusOK, "Resync complete. Order history updated.")
}

func findSyncScript() (string, string, error) {
	candidate := filepath.Join("scripts", "sync-amazon-data.py")

	absPath, err := filepath.Abs(candidate)
	if err != nil {
		return "", "", fmt.Errorf("unable to resolve sync script path: %w", err)
	}
	return absPath, filepath.Dir(absPath), nil
}

func writeJSON(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(&Response{Msg: msg})
}

func trimOutput(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}
