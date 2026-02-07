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

	http.HandleFunc("/provider-details", detials_handler)
	http.HandleFunc("/resync", resync_handler)
	http.HandleFunc("/session/history", history_handler)
	http.HandleFunc("/session/ask", ask_handler)
	http.HandleFunc("/session/create", create_session_handler)

	http.HandleFunc("/test-connection", test_connection_handler)

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
