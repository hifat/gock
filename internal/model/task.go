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
	gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
