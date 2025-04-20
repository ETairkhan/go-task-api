package main

import (
	"log"
	"net/http"
	"task/handler"
	"task/service"
	"task/storage"

	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewMemoryStorage()
	svc := service.NewTaskService(store)
	h := handler.NewTaskHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", h.GetTask).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
