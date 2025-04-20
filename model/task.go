package model

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
)

type Task struct {
	ID     string     `json:"id"`
	Result string     `json:"result,omitempty"`
	Status TaskStatus `json:"status"`
}
