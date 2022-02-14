package service

import (
	goweb "github.com/oleg5896/go-web"
	"github.com/oleg5896/go-web/pkg/repository"
)

type AddItem interface {
	AddFile(file goweb.File) (int, error)
}

type GetItem interface {
}

type GetList interface {
}

type Service struct {
	AddItem
	GetItem
	GetList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AddItem: NewAddService(repos.AddItem),
	}
}
