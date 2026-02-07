package main

import (
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Msg string `json:"msg"`
}


func main() {
	http.HandleFunc("/provider-details", detials_handler)
	http.HandleFunc("/resync", resync_handler)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
