package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!\n")
}

func main() {
	// fake db using struck and slice
	type event struct {
		ID 			string `json:"ID"`
		Title 		string `json:"Title"`
		Description string `json:"Description"`

	}

	type allEvents []event

	var events = allEvents {
		{
			ID: 		 "1",
			Title: 		 "Intro to go.",
			Description: "Intro to rest apis using go.",
		},
	}

	// index page
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}