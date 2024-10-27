package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Text string `json:"message"`
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

	newMessage := Message{Text: body.Text}
	DB.Create(&newMessage)

	// message = body.Message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message received: %s", body.Text)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	DB.Find(&messages)

	if len(messages) == 0 {
		http.Error(w, "No message available", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
	// fmt.Fprintf(w, "Hello, %s", message)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	// Создаем новый маршрутизатор
	router := mux.NewRouter()
	router.HandleFunc("/api/post", postHandler).Methods("POST")
	router.HandleFunc("/api/get", getHandler).Methods("GET")

	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", router)
}
