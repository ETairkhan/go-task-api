package handler

import (
	"encoding/json"
	"net/http"
	"task/service"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	id := h.service.CreateTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"task_id": id})
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, found := h.service.GetTask(id)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
