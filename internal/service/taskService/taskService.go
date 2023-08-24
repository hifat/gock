package taskService

import (
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"github.com/hifat/gock/internal/utils/ernos"
	zlog "github.com/hifat/gock/pkg"
)

var NewTaskServiceSet = wire.NewSet(NewTaskService)

type taskService struct {
	taskRepo taskDomain.TaskRepository
}

func NewTaskService(taskRepo taskDomain.TaskRepository) taskDomain.TaskService {
	return &taskService{taskRepo}
}

func (r *taskService) Get(res *[]taskDomain.Task) error {
	err := r.taskRepo.Get(res)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}

func (r *taskService) GetByID(res *taskDomain.Task, taskID uuid.UUID) error {
	err := r.taskRepo.GetByID(res, taskID)
	if err != nil {
		if err.Error() == ernos.M.RECORD_NOTFOUND {
			return ernos.NotFound()
		}

		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}

func (r *taskService) Create(req *taskDomain.TaskRequest) (*taskDomain.Task, error) {
	res, err := r.taskRepo.Create(req)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	return res, nil
}

func (r *taskService) Update(req *taskDomain.TaskRequest, taskID uuid.UUID) (*taskDomain.Task, error) {
	res, err := r.taskRepo.Update(req, taskID)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	return res, nil
}

func (r *taskService) Delete(taskID uuid.UUID) error {
	err := r.taskRepo.Delete(taskID)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}
