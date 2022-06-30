// Package service collects another services into one
package service

import "Ozon_fintech/pkg/repository"

const (
	LengthLink = 10
)

// Validator - validate link
type Validator interface {
	ValidateLink(link string) error
}

type Linker interface {
	GetFullLink(shortLink string) (string, error)
	SetShortLink(fullLink string) (string, error)
}

type Service struct {
	Validator
	Linker
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Validator: NewValid(),
		Linker:    NewLinkerService(repo),
	}
}
