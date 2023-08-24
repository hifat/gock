package taskRepository

import (
	"context"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"github.com/hifat/gock/internal/model"
	"gorm.io/gorm"
)

var NewTaskRepoSet = wire.NewSet(NewtaskRepository)

type taskRepository struct {
	db *gorm.DB
}

func NewtaskRepository(db *gorm.DB) taskDomain.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Get(ctx context.Context, res *[]taskDomain.Task) error {
	return r.db.WithContext(ctx).
		Model(&model.Task{}).
		Find(&res).Error
}

func (r *taskRepository) GetByID(ctx context.Context, res *taskDomain.Task, taskID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ?", taskID).
		Find(&res).Error
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
