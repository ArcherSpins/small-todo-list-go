package main

import (
	"log"
	"net/http"
	"todo-list/handlers"
	"todo-list/models"

	"github.com/gorilla/mux"
)

func main() {
	// Подключение к базе данных PostgreSQL
	connStr := "host=localhost port=8080 user=root password=secret dbname=todo_list sslmode=disable"
	err := models.InitDB(connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Создаем новый маршрутизатор
	r := mux.NewRouter()

	// Определяем маршруты для обработки HTTP-запросов
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":2001", r))
}
