package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID   uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	Name string    `gorm:"varchar(150)"`
	Done bool      `gorm:"boolean; default:false"`

	CreatedAt *time.Time      `gorm:"timestamp"`
	UpdatedAt *time.Time      `gorm:"timestamp"`
	DeletedAt *gorm.DeletedAt `gorm:"timestamp"`
}
