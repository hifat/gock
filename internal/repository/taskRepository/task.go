package taskRepository

import (
	"context"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"github.com/hifat/gock/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source=./task.go -destination=./mockTaskRepository/mockTaskRepository.go -package=mockTaskRepository
type ITaskRepository interface {
	Get(ctx context.Context, res *[]taskDomain.Task) error
	GetByID(ctx context.Context, res *taskDomain.Task, taskID uuid.UUID) error
	Create(ctx context.Context, req *taskDomain.TaskRequest) (*taskDomain.Task, error)
	Update(ctx context.Context, req *taskDomain.TaskRequest, taskID uuid.UUID) (*taskDomain.Task, error)
	Delete(ctx context.Context, taskID uuid.UUID) error
}

var NewTaskRepoSet = wire.NewSet(New)

type taskRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Get(ctx context.Context, res *[]taskDomain.Task) error {
	return r.db.WithContext(ctx).
		Model(&model.Task{}).
		Select("id, name, done, created_at").
		Find(&res).Error
}

func (r *taskRepository) GetByID(ctx context.Context, res *taskDomain.Task, taskID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&model.Task{}).
		Select("id, name, done, created_at").
		Where("id = ?", taskID).
		First(&res).Error
}

func (r *taskRepository) Create(ctx context.Context, req *taskDomain.TaskRequest) (res *taskDomain.Task, err error) {
	if err := r.db.WithContext(ctx).
		Create(&model.Task{
			Name: req.Name,
		}).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *taskRepository) Update(ctx context.Context, req *taskDomain.TaskRequest, taskID uuid.UUID) (res *taskDomain.Task, err error) {
	if err := r.db.WithContext(ctx).
		Where("id = ?", taskID).
		Updates(&model.Task{
			Name: req.Name,
		}).
		Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *taskRepository) Delete(ctx context.Context, taskID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", taskID).
		Delete(&model.Task{}).Error
}
