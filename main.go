package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message string

type requestBody struct {
	Message string `json:"message"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	message = body.Message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message received: %s", message)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if message == "" {
		http.Error(w, "No message available", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s", message)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/post", postHandler).Methods("POST")
	router.HandleFunc("/api/get", getHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
