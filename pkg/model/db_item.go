// Package model contains help model
package model

type DBItem struct {
	Id        int    `db:"id"`
	ShortLink string `db:"short_link"`
	FullLink  string `db:"full_link"`
}

type LinkInfo struct {
	FullLink string `json:"full_link" example:"https://pkg.go.dev/"`
}
