package main

import (
	"fmt"
	"log"
	"net/http"
	"x/db"
)

type Response struct {
	Msg string `json:"msg"`
}

var database db.Database


func main() {
	var err error
	if database, err = db.Open("primary.db"); err != nil {
		panic(err)
	}

	http.HandleFunc("/provider-details", detials_handler)
	http.HandleFunc("/resync", resync_handler)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
