package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Msg string `json:"msg"`
}


func noop(w http.ResponseWriter, r *http.Request) {
	response := Response{"This endpoint does nothing."}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/provider-details", noop)
	http.HandleFunc("/resync", noop)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
