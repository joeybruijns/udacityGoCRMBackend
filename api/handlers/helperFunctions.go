package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func customerNotExists(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)

	_, err := fmt.Fprintf(w, "No user matches the ID specified..\n")
	handleError(err)
}
