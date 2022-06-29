package repository

import (
	strgen "Ozon_fintech/pkg/string_generator"
	"github.com/jmoiron/sqlx"
)

type LinkerPostgres struct {
	db  *sqlx.DB
	gen *strgen.StringGenerator
}

func (l *LinkerPostgres) GetFullLink(shortLink string) (string, error) {
	panic(any("implement me"))
}

func (l *LinkerPostgres) SetShortLink(fullLink string) (string, error) {
	panic(any("implement me"))
}

func NewLinkerPostgres(db *sqlx.DB, gen *strgen.StringGenerator) *LinkerPostgres {
	return &LinkerPostgres{
		db:  db,
		gen: gen,
	}
}
