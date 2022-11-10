package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// handleBasicError handles basis errors when coding and decoding JSON data.
func handleBasicError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// customerNotExists gives a 404 Not Found error back when a customer is not found by ID.
func customerNotExists(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)

	_, err := fmt.Fprintf(w, "No user matches the ID specified..\n")
	handleBasicError(err)
}
