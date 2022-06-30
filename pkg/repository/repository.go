// Package repository implement data storing logic
package repository

import (
	"Ozon_fintech/pkg/storage"
	strgen "Ozon_fintech/pkg/string_generator"
	"github.com/jmoiron/sqlx"
)

type Linker interface {
	GetFullLink(shortLink string) (string, error)
	SetShortLink(fullLink string) (string, error)
}

type Repository struct {
	Linker
}

func NewRepositoryPostgres(db *sqlx.DB, gen *strgen.StringGenerator) *Repository {
	return &Repository{
		Linker: NewLinkerPostgres(db, gen),
	}
}

func NewRepositoryStorage(db *storage.Storage) *Repository {
	return &Repository{
		Linker: NewLinkerStorage(db),
	}
}
