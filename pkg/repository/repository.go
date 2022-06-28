package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
}

func NewRepositoryPostgres(db *sqlx.DB) *Repository {
	return &Repository{}
}

func NewRepositoryCustom() *Repository {
	return &Repository{}
}
