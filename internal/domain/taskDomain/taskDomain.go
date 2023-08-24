package taskDomain

import "github.com/google/uuid"

type TaskService interface {
	Get(res *[]Task) error
	GetByID(res *Task, taskID uuid.UUID) error
	Create(req *TaskRequest) (*Task, error)
	Update(req *TaskRequest, taskID uuid.UUID) (*Task, error)
	Delete(taskID uuid.UUID) error
}

type TaskRepository interface {
	Get(res *[]Task) error
	GetByID(res *Task, taskID uuid.UUID) error
	Create(req *TaskRequest) (*Task, error)
	Update(req *TaskRequest, taskID uuid.UUID) (*Task, error)
	Delete(taskID uuid.UUID) error
}

type Task struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Done bool      `json:"done"`
}

type TaskRequest struct {
	Name string `json:"name" binding:"required"`
	Done bool
}
