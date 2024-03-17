package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var people []Person

func main() {
	log.Println("server started")
	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/health", healthCheckHandler)
	err := http.ListenAndServe("localhost:8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPost:
		postPeople(w, r)
	default:
		http.Error(w, "Invalid http method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		http.Error(w, "Something wrong with getting people", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "get people")
}

func postPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Something wrong with this person", http.StatusBadRequest)
		return
	}
	people = append(people, person)
	fmt.Fprintf(w, "post new person %v", person)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "http server works correctly")
}
