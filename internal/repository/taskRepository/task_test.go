package taskRepository_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"github.com/hifat/gock/internal/repository/taskRepository"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testTaskRepoSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock

	taskRepo taskRepository.ITaskRepository
}

func TestUserRepo(t *testing.T) {
	suite.Run(t, &testTaskRepoSuite{})
}

func (s *testTaskRepoSuite) SetupSuite() {
	dbMock, mock, err := sqlmock.New()
	s.Require().NoError(err)
	dialector := postgres.New(postgres.Config{
		Conn:       dbMock,
		DriverName: "postgres",
	})

	gormMock, err := gorm.Open(dialector, &gorm.Config{})
	s.Require().NoError(err)

	s.db = gormMock
	s.mock = mock
	s.taskRepo = taskRepository.New(gormMock)
}

func (s *testTaskRepoSuite) AfterTest(_, _ string) {
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (u *testTaskRepoSuite) TestUserRepo_GetByID() {
	u.Run("success - get task by id", func() {
		ctx := context.Background()

		var (
			taskID    = uuid.New()
			name      = "get a task"
			done      = false
			createdAt = time.Now()
		)

		expectedSQL := `SELECT .+ FROM "tasks" WHERE id = .+`
		row := sqlmock.NewRows([]string{"id", "name", "done", "created_at"}).
			AddRow(taskID, name, done, createdAt)
		u.mock.ExpectQuery(expectedSQL).
			WithArgs(taskID).
			WillReturnRows(row)

		var res taskDomain.Task
		err := u.taskRepo.GetByID(ctx, &res, taskID)
		u.Require().NoError(err)
		u.Require().Equal(res.ID, taskID)
		u.Require().Equal(res.Name, name)
		u.Require().False(res.Done)
		u.Require().NotEmpty(res.CreatedAt)
	})
}
