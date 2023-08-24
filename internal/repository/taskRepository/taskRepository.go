package taskRepository

import (
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

func (r *taskRepository) Get(res *[]taskDomain.Task) error {
	return r.db.Model(&model.Task{}).
		Find(&res).Error
}

func (r *taskRepository) GetByID(res *taskDomain.Task, taskID uuid.UUID) error {
	return r.db.Model(&model.Task{}).
		Where("id = ?", taskID).
		Find(&res).Error
}

func (r *taskRepository) Create(req *taskDomain.TaskRequest) (res *taskDomain.Task, err error) {
	if err := r.db.Create(&model.Task{
		Name: req.Name,
	}).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *taskRepository) Update(req *taskDomain.TaskRequest, taskID uuid.UUID) (res *taskDomain.Task, err error) {
	if err := r.db.Where("id = ?", taskID).
		Updates(&model.Task{
			Name: req.Name,
		}).
		Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *taskRepository) Delete(taskID uuid.UUID) error {
	return r.db.Where("id = ?", taskID).
		Delete(&model.Task{}).Error
}
