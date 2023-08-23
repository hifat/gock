package taskService

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
)

var NewTaskServiceSet = wire.NewSet(NewTaskService)

type taskService struct {
	taskRepo taskDomain.TaskRepository
}

func NewTaskService(taskRepo taskDomain.TaskRepository) taskDomain.TaskService {
	return &taskService{taskRepo}
}

func (s *taskService) Get() string {
	return s.taskRepo.Get()
}
