package model

type DBItem struct {
	Id        int    `db:"id"`
	ShortLink string `db:"short_link"`
	FullLink  string `db:"full_link"`
}
