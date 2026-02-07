package main

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"os/exec"
)

func image_resource_handler(w http.ResponseWriter, r *http.Request) {
	// NOTE: handler uses http.Error instead of write header since it is writing binary data

	img_id := r.PathValue("id")
	if img_id == "" {
		http.Error(w, "Missing image resource id in path", http.StatusBadRequest)
		return
	}

	img_path, err := url.PathUnescape(img_id)
	if err != nil {
		http.Error(w, "Invalid image resource id in path", http.StatusBadRequest)
		return
	}

	var stderr bytes.Buffer
	cmd := exec.CommandContext(r.Context(), python_executable, "remove-bg.py", img_path)
	cmd.Dir = "./scripts"
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		log.Println(cmd.String())
		log.Println(err)
		http.Error(w, "An error occured while getting resource", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/*")
	w.Write(output)
}
