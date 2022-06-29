package repository

import (
	"Ozon_fintech/pkg/storage"
)

type LinkerStorage struct {
	store *storage.Storage
}

func (l *LinkerStorage) GetFullLink(shortLink string) (string, error) {
	return l.store.Get(shortLink)
}

func (l *LinkerStorage) SetShortLink(fullLink string) (string, error) {
	return l.store.Post(fullLink)
}

func NewLinkerStorage(store *storage.Storage) *LinkerStorage {
	return &LinkerStorage{store: store}
}
