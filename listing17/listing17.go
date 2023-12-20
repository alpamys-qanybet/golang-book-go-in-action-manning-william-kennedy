package main

import (
	"log"
	"net/http"

	"books/go-in-action/listing17/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
