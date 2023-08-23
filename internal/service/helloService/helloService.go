package helloService

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/helloDomain"
)

var NewHelloServiceSet = wire.NewSet(NewHelloService)

type HelloService struct {
	helloRepo helloDomain.HelloRepository
}

func NewHelloService(helloRepo helloDomain.HelloRepository) helloDomain.HelloService {
	return &HelloService{helloRepo}
}

func (s *HelloService) GetHello() string {
	return s.helloRepo.GetHello()
}
