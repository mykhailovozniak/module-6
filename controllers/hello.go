package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func HelloController(res http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(res, "Hello world")

	if err != nil {
		log.Fatal("Error", err)
	}

	log.Print("Successfully return html - hello")
}
