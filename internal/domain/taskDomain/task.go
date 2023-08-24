package taskDomain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Done bool      `json:"done"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type TaskRequest struct {
	Name string `json:"name" binding:"required"`
	Done bool
}
