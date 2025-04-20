package storage

import (
	"sync"
	"task/model"
)

type MemoryStorage struct {
	mu    sync.RWMutex
	tasks map[string]*model.Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[string]*model.Task),
	}
}

func (s *MemoryStorage) Save(task *model.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.ID] = task
}

func (s *MemoryStorage) Get(id string) (*model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, exists := s.tasks[id]
	return task, exists
}
