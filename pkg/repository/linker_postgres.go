package repository

import (
	"Ozon_fintech/pkg/model"
	"Ozon_fintech/pkg/storage"
	strgen "Ozon_fintech/pkg/string_generator"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LinkerPostgres struct {
	db  *sqlx.DB
	gen *strgen.StringGenerator
}

// GetFullLink - get full link from Postgres by short link
func (l *LinkerPostgres) GetFullLink(shortLink string) (string, error) {
	var count int
	tx, err := l.db.Begin()
	if err != nil {
		return "", err
	}

	querySelect := fmt.Sprintf("SELECT Count(id) FROM %s WHERE short_link=$1", linksTable)
	row := tx.QueryRow(querySelect, shortLink)
	err = row.Scan(&count)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	if count > 0 {
		var item model.DBItem
		query := fmt.Sprintf("SELECT * FROM %s WHERE short_link=$1 LIMIT 1", linksTable)
		row := tx.QueryRow(query, shortLink)
		err = row.Scan(&item.Id, &item.ShortLink, &item.FullLink)

		if err != nil {
			tx.Rollback()
			return "", err
		}

		return item.FullLink, tx.Commit()
	} else {
		tx.Rollback()
		return "", errors.New(storage.KeyError)
	}

}

// SetShortLink - insert full link into Postgres and return short link
func (l *LinkerPostgres) SetShortLink(fullLink string) (string, error) {
	var count int
	tx, err := l.db.Begin()
	if err != nil {
		return "", err
	}

	querySelect := fmt.Sprintf("SELECT Count(id) FROM %s WHERE full_link=$1", linksTable)
	row := tx.QueryRow(querySelect, fullLink)
	err = row.Scan(&count)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	if count > 0 {
		var item model.DBItem
		query := fmt.Sprintf("SELECT * FROM %s WHERE full_link=$1 LIMIT 1", linksTable)
		row := tx.QueryRow(query, fullLink)
		err = row.Scan(&item.Id, &item.ShortLink, &item.FullLink)

		if err != nil {
			tx.Rollback()
			return "", err
		}

		err = tx.Commit()
		if err != nil {
			return "", err
		}

		return item.ShortLink, errors.New(storage.KeyExistError)
	} else {
		for {
			generateString := l.gen.GenerateString()
			query := fmt.Sprintf("INSERT INTO %s (short_link, full_link) values ($1, $2)",
				linksTable)

			_, err := tx.Exec(query, generateString, fullLink)

			if err == nil {
				err = tx.Commit()
				if err != nil {
					return "", err
				} else {
					return generateString, nil
				}
			}

			querySelect := fmt.Sprintf("SELECT Count(id) FROM %s WHERE full_link=$1", linksTable)
			row := tx.QueryRow(querySelect, fullLink)
			err = row.Scan(&count)

			if err != nil {
				tx.Rollback()
				return "", err
			}

			if count > 0 {
				var item model.DBItem
				query := fmt.Sprintf("SELECT * FROM %s WHERE full_link=$1 LIMIT 1", linksTable)
				row := tx.QueryRow(query, fullLink)
				err = row.Scan(&item.Id, &item.ShortLink, &item.FullLink)

				if err != nil {
					tx.Rollback()
					return "", err
				}

				err = tx.Commit()
				if err != nil {
					return "", err
				}

				return item.ShortLink, errors.New(storage.KeyExistError)
			}
		}
	}
}

func NewLinkerPostgres(db *sqlx.DB, gen *strgen.StringGenerator) *LinkerPostgres {
	return &LinkerPostgres{
		db:  db,
		gen: gen,
	}
}
