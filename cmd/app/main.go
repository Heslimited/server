package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body Message
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	DB.Create(&body)

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
}

func patchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var body Message
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var message Message
	if err := DB.First(&message, id).Error; err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	message.Text = body.Text
	DB.Save(&message)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message updated: %s", body.Text)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var body Message
	if err := DB.First(&body, id).Error; err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	DB.Delete(&body)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message deleted")
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
	router.HandleFunc("/api/patch/{id}", patchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", deleteHandler).Methods("DELETE")

	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", router)
}
