package taskService_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"github.com/hifat/gock/internal/repository/taskRepository/mockTaskRepository"
	"github.com/hifat/gock/internal/service/taskService"
	"github.com/hifat/gock/internal/utils/ernos"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testUserServiceSuite struct {
	suite.Suite

	mockTaskRepo *mockTaskRepository.MockITaskRepository
	underTest    taskService.ITaskService
}

func (s *testUserServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.mockTaskRepo = mockTaskRepository.NewMockITaskRepository(ctrl)
	s.underTest = taskService.New(s.mockTaskRepo)
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, &testUserServiceSuite{})
}

func (s *testUserServiceSuite) TestUserService_Get() {
	s.Run("success - get", func() {
		ctx := context.Background()

		expectedTasks := []taskDomain.Task{
			{
				ID:        uuid.New(),
				Name:      "get task 1",
				Done:      false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        uuid.New(),
				Name:      "get task 2",
				Done:      false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		s.mockTaskRepo.EXPECT().
			Get(ctx, gomock.Any()).
			DoAndReturn(func(ctx context.Context, res *[]taskDomain.Task) error {
				*res = expectedTasks
				return nil
			})

		tasks := []taskDomain.Task{}
		err := s.underTest.Get(ctx, &tasks)
		s.Require().NoError(err)
		for i, task := range tasks {
			s.Require().Equal(expectedTasks[i].Name, task.Name)
			s.Require().NotNil(task.Done)
			s.Require().NotEmpty(task.CreatedAt)
			s.Require().NotEmpty(task.UpdatedAt)
		}
	})
	s.Run("success - empty slice", func() {
		ctx := context.Background()

		expectedTasks := []taskDomain.Task{}
		s.mockTaskRepo.EXPECT().
			Get(ctx, gomock.Any()).
			DoAndReturn(func(ctx context.Context, res *[]taskDomain.Task) error {
				*res = expectedTasks
				return nil
			})

		tasks := []taskDomain.Task{}
		err := s.underTest.Get(ctx, &tasks)
		s.Require().NoError(err)
		s.Require().NotNil(tasks)
		s.Require().Empty(tasks)
	})
	s.Run("failed - other error", func() {
		ctx := context.Background()

		s.mockTaskRepo.EXPECT().
			Get(ctx, gomock.Any()).
			Return(errors.New("other error"))

		tasks := []taskDomain.Task{}
		err := s.underTest.Get(ctx, &tasks)
		s.Require().Error(err)
		s.Require().Equal(ernos.InternalServerError(), err)
	})
}

func (s *testUserServiceSuite) TestUserService_GetByID() {
	s.Run("success - get by ID", func() {
		ctx := context.Background()

		expectedTask := taskDomain.Task{
			ID:        uuid.New(),
			Name:      "get task detail",
			Done:      false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		s.mockTaskRepo.EXPECT().
			GetByID(ctx, gomock.Any(), expectedTask.ID).
			DoAndReturn(func(ctx context.Context, res *taskDomain.Task, taskID uuid.UUID) error {
				*res = expectedTask
				return nil
			})

		var res taskDomain.Task
		err := s.underTest.GetByID(ctx, &res, expectedTask.ID)
		s.Require().NoError(err)
		s.Require().Equal(expectedTask.ID, res.ID)
		s.Require().Equal(expectedTask.Name, res.Name)
		s.Require().Equal(expectedTask.Done, res.Done)
		s.Require().Equal(expectedTask.CreatedAt, res.CreatedAt)
		s.Require().Equal(expectedTask.UpdatedAt, res.UpdatedAt)
	})
	s.Run("failed - record not found", func() {
		ctx := context.Background()

		taskID := uuid.New()
		s.mockTaskRepo.EXPECT().
			GetByID(ctx, gomock.Any(), taskID).
			Return(errors.New(ernos.M.RECORD_NOT_FOUND))

		var res taskDomain.Task
		err := s.underTest.GetByID(ctx, &res, taskID)
		s.Require().Error(err)
		s.Require().Equal(ernos.NotFound(), err)
	})
	s.Run("failed - other error", func() {
		ctx := context.Background()

		taskID := uuid.New()
		s.mockTaskRepo.EXPECT().
			GetByID(ctx, gomock.Any(), taskID).
			Return(errors.New("other error"))

		var res taskDomain.Task
		err := s.underTest.GetByID(ctx, &res, taskID)
		s.Require().Error(err)
		s.Require().Equal(ernos.InternalServerError(), err)
	})
}

func (s *testUserServiceSuite) TestUserService_Create() {
	s.Run("success - create task", func() {
		ctx := context.TODO()

		newTask := taskDomain.TaskRequest{
			Name: "create some task",
		}

		res := taskDomain.Task{
			ID:        uuid.New(),
			Name:      newTask.Name,
			Done:      false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		s.mockTaskRepo.EXPECT().Create(ctx, &newTask).Return(&res, nil)

		task, err := s.underTest.Create(ctx, &newTask)
		s.Require().NoError(err)
		s.Require().Equal(newTask.Name, task.Name)
		s.Require().False(task.Done)
		s.Require().NotEmpty(task.CreatedAt)
		s.Require().NotEmpty(task.UpdatedAt)
	})
	s.Run("failed - other error", func() {
		ctx := context.TODO()

		s.mockTaskRepo.EXPECT().
			Create(ctx, gomock.Any()).
			Return(nil, errors.New("other error"))

		var req taskDomain.TaskRequest
		_, err := s.underTest.Create(ctx, &req)
		s.Require().Error(err)
		s.Require().Equal(ernos.InternalServerError(), err)
	})
}

func (s *testUserServiceSuite) TestUserService_Update() {
	s.Run("success - update", func() {
		ctx := context.Background()

		taskID := uuid.New()
		originTask := taskDomain.Task{
			ID:   taskID,
			Name: "this is original task",
			Done: false,
		}

		updatedTask := taskDomain.TaskRequest{
			Name: "this is updated task",
		}

		expectedTask := taskDomain.Task{
			ID:   taskID,
			Name: "this is updated task",
			Done: false,
		}

		s.mockTaskRepo.EXPECT().
			Update(ctx, &updatedTask, taskID).
			DoAndReturn(func(ctx context.Context, req *taskDomain.TaskRequest, taskID uuid.UUID) (*taskDomain.Task, error) {
				originTask.Name = req.Name
				return &originTask, nil
			})

		task, err := s.underTest.Update(ctx, &updatedTask, taskID)
		s.Require().NoError(err)
		s.Require().Equal(expectedTask.ID, task.ID)
		s.Require().Equal(expectedTask.Name, task.Name)
		s.Require().Equal(expectedTask.Done, task.Done)
	})
	s.Run("failed - other error", func() {
		ctx := context.Background()

		s.mockTaskRepo.EXPECT().
			Update(ctx, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("other error"))

		var req taskDomain.TaskRequest
		task, err := s.underTest.Update(ctx, &req, uuid.New())
		s.Require().Error(err)
		s.Require().Equal(ernos.InternalServerError(), err)
		s.Require().Nil(task)
	})
}

func (s *testUserServiceSuite) TestUserService_Delete() {
	s.Run("success - delete", func() {
		ctx := context.Background()

		s.mockTaskRepo.EXPECT().
			Delete(ctx, gomock.Any()).
			Return(nil)

		err := s.underTest.Delete(ctx, uuid.New())
		s.Require().NoError(err)
	})
	s.Run("failed - other error", func() {
		ctx := context.Background()

		s.mockTaskRepo.EXPECT().
			Delete(ctx, gomock.Any()).
			Return(errors.New("other error"))

		err := s.underTest.Delete(ctx, uuid.New())
		s.Require().Error(err)
		s.Require().Equal(ernos.InternalServerError(), err)
	})
}
