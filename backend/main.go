package main

import (
	"fmt"
	"log"
	"net/http"
	"x/db"
)

var database db.Database
var python_executable string

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	var err error
	if python_executable, err = find_python(); err != nil {
		log.Fatal("Failed to find python executable")
	}

	if database, err = db.Open("primary.db"); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/provider-details", detials_handler)
	mux.HandleFunc("/set-amazon-credentials", set_amazon_credentials_handler)
	mux.HandleFunc("/resync", resync_handler)
	mux.HandleFunc("/recommendations", recommendations_handler)
	mux.HandleFunc("/resource/image/{id}", image_resource_handler)
	mux.HandleFunc("/session/history", history_handler)
	mux.HandleFunc("/session/ask", ask_handler)
	mux.HandleFunc("/session/create", create_session_handler)
	mux.HandleFunc("/test-connection", test_connection_handler)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
