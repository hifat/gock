package taskService

import (
	"context"

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

func (r *taskService) Get(ctx context.Context, res *[]taskDomain.Task) error {
	err := r.taskRepo.Get(ctx, res)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}

func (r *taskService) GetByID(ctx context.Context, res *taskDomain.Task, taskID uuid.UUID) error {
	err := r.taskRepo.GetByID(ctx, res, taskID)
	if err != nil {
		if err.Error() == ernos.M.RECORD_NOTFOUND {
			return ernos.NotFound()
		}

		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}

func (r *taskService) Create(ctx context.Context, req *taskDomain.TaskRequest) (*taskDomain.Task, error) {
	res, err := r.taskRepo.Create(ctx, req)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	return res, nil
}

func (r *taskService) Update(ctx context.Context, req *taskDomain.TaskRequest, taskID uuid.UUID) (*taskDomain.Task, error) {
	res, err := r.taskRepo.Update(ctx, req, taskID)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	return res, nil
}

func (r *taskService) Delete(ctx context.Context, taskID uuid.UUID) error {
	err := r.taskRepo.Delete(ctx, taskID)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}
