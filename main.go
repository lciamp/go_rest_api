package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!\n")
}
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

// function to add events
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter data with the event title and description only in order to update.")
	}
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)

}

func main() {
	// index page
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}