package taskService

import (
	"context"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"github.com/hifat/gock/internal/repository/taskRepository"
	"github.com/hifat/gock/internal/utils/ernos"
	zlog "github.com/hifat/gock/pkg"
)

//go:generate mockgen -source=./task.go -destination=./mockTaskService/mockTaskService.go -package=mockUserService
type ITaskService interface {
	Get(ctx context.Context, res *[]taskDomain.Task) error
	GetByID(ctx context.Context, res *taskDomain.Task, taskID uuid.UUID) error
	Create(ctx context.Context, req *taskDomain.TaskRequest) (*taskDomain.Task, error)
	Update(ctx context.Context, req *taskDomain.TaskRequest, taskID uuid.UUID) (*taskDomain.Task, error)
	Delete(ctx context.Context, taskID uuid.UUID) error
}

var NewTaskServiceSet = wire.NewSet(New)

type taskService struct {
	taskRepo taskRepository.ITaskRepository
}

func New(taskRepo taskRepository.ITaskRepository) ITaskService {
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
		if err.Error() == ernos.M.RECORD_NOT_FOUND {
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
