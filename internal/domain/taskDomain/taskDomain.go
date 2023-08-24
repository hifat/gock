package taskDomain

import (
	"context"

	"github.com/google/uuid"
)

type TaskService interface {
	Get(ctx context.Context, res *[]Task) error
	GetByID(ctx context.Context, res *Task, taskID uuid.UUID) error
	Create(ctx context.Context, req *TaskRequest) (*Task, error)
	Update(ctx context.Context, req *TaskRequest, taskID uuid.UUID) (*Task, error)
	Delete(ctx context.Context, taskID uuid.UUID) error
}

type TaskRepository interface {
	Get(ctx context.Context, res *[]Task) error
	GetByID(ctx context.Context, res *Task, taskID uuid.UUID) error
	Create(ctx context.Context, req *TaskRequest) (*Task, error)
	Update(ctx context.Context, req *TaskRequest, taskID uuid.UUID) (*Task, error)
	Delete(ctx context.Context, taskID uuid.UUID) error
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
