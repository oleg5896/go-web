package service

import "github.com/oleg5896/go-web/pkg/repository"

type WebItem interface{

}

type WebList interface{

}

type Service struct {
	WebItem
	WebList
}

func newService(repos *repository.Repository) *Service{
	return &Service{}
}