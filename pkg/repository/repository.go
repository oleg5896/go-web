package repository

import "github.com/jmoiron/sqlx"

type WebItem interface {
}

type WebList interface {
}

type Repository struct {
	WebItem
	WebList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
