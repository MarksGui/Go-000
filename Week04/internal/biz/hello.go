package biz

import (
	"Go-000/Week04/internal/data"
	"Go-000/Week04/internal/entity"
)

type HelloRepo interface {
	GetHello(id int64) (*entity.Hello, error)
}

type HelloBIZ struct {
	repo HelloRepo
}

func NewHelloBIZ(repo *data.HelloRepository) *HelloBIZ {
	return &HelloBIZ{repo: repo}
}

func (b *HelloBIZ) GetHello(id int64) (*entity.Hello, error) {
	return b.repo.GetHello(id)
}
