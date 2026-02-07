package main

import (
	"net/http"
)

func test_connection_handler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "OK")
}
