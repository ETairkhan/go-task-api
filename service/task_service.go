package service

import (
	"fmt"
	"time"
	"task/model"
	"task/storage"

	"github.com/google/uuid"
)

type TaskService struct {
	storage *storage.MemoryStorage
}

func NewTaskService(store *storage.MemoryStorage) *TaskService {
	return &TaskService{storage: store}
}

func (s *TaskService) CreateTask() string {
	id := uuid.New().String()
	task := &model.Task{
		ID:     id,
		Status: model.StatusPending,
	}
	s.storage.Save(task)

	go s.runTask(task)
	return id
}

func (s *TaskService) runTask(task *model.Task) {
	task.Status = model.StatusRunning
	s.storage.Save(task)

	time.Sleep(20 * time.Second) // имитация долгой операции

	task.Result = fmt.Sprintf("Task %s completed successfully", task.ID)
	task.Status = model.StatusCompleted
	s.storage.Save(task)
}

func (s *TaskService) GetTask(id string) (*model.Task, bool) {
	return s.storage.Get(id)
}
