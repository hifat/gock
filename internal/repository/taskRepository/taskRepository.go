package taskRepository

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"gorm.io/gorm"
)

var NewTaskRepoSet = wire.NewSet(NewtaskRepository)

type taskRepository struct {
	db *gorm.DB
}

func NewtaskRepository(db *gorm.DB) taskDomain.TaskRepository {
	return &taskRepository{db}
}

func (s *taskRepository) Get() string {
	return "hi gock!"
}
