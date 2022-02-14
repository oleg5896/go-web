package service

import (
	goweb "github.com/oleg5896/go-web"
	"github.com/oleg5896/go-web/pkg/repository"
)

type AddService struct {
	repo repository.AddItem
}

func NewAddService(repo repository.AddItem) *AddService {
	return &AddService{repo: repo}
}

func (s *AddService) AddFile(file goweb.File) (int, error) {
	return s.repo.AddFile(file)
}
