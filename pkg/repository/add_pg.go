package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	goweb "github.com/oleg5896/go-web"
)

type ItemPg struct {
	db *sqlx.DB
}

func NewItemPg(db *sqlx.DB) *ItemPg {
	return &ItemPg{db: db}
}

func (r *ItemPg) AddFile(file goweb.File) (int, error) {
	var id int
	querySelect := fmt.Sprintf("select id from %s where path=$1", filesTable)
	row := r.db.QueryRow(querySelect, file.Path)
	if err := row.Scan(&id); err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	} else {
		return id, nil
	}
	query := fmt.Sprintf("INSERT INTO %s (path) values ($1) returning id", filesTable)
	row = r.db.QueryRow(query, file.Path)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
