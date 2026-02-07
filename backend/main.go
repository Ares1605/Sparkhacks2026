package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"x/db"
)

var database db.Database
var python_executable string

func main() {
	process_args()

	var err error
	if database, err = db.Open("primary.db"); err != nil {
		panic(err)
	}

	http.Handle("/provider-details", withCORS(http.HandlerFunc(detials_handler)))
	http.Handle("/resync", withCORS(http.HandlerFunc(resync_handler)))
	http.Handle("/session/history", withCORS(http.HandlerFunc(history_handler)))
	http.Handle("/session/ask", withCORS(http.HandlerFunc(ask_handler)))
	http.Handle("/session/create", withCORS(http.HandlerFunc(create_session_handler)))

	http.Handle("/test-connection", withCORS(http.HandlerFunc(test_connection_handler)))

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func process_args() {
	python := flag.String("python", "python3", "Python executable")
	flag.Parse()

	python_executable = *python
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
