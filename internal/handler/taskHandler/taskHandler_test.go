package taskHandler_test

import (
	"github.com/hifat/gock/internal/handler/taskHandler"
	"github.com/hifat/gock/internal/service/taskService/mockTaskService"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testHandlerServiceSuite struct {
	suite.Suite

	mockTaskServ *mockTaskService.MockITaskService
	underTest    taskHandler.TaskHandler
}

func (s *testHandlerServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.mockTaskServ = mockTaskService.NewMockITaskService(ctrl)
	s.underTest = taskHandler.New(s.mockTaskServ)
}
