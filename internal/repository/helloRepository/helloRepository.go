package helloRepository

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/helloDomain"
	"gorm.io/gorm"
)

var NewHelloRepoSet = wire.NewSet(NewHelloRepository)

type helloRepository struct {
	db *gorm.DB
}

func NewHelloRepository(db *gorm.DB) helloDomain.HelloRepository {
	return &helloRepository{db}
}

func (s *helloRepository) GetHello() string {
	return "hi gock!"
}
