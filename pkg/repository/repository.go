package repository

import (
	"github.com/jmoiron/sqlx"
	goweb "github.com/oleg5896/go-web"
)

type AddItem interface {
	AddFile(file goweb.File) (int, error)
}

type GetItem interface {
}

type GetList interface {
}

type Repository struct {
	AddItem
	GetItem
	GetList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AddItem: NewItemPg(db),
	}
}
