package service

import "Ozon_fintech/pkg/repository"

type LinkerService struct {
	repo *repository.Repository
}

func (l *LinkerService) GetFullLink(shortLink string) (string, error) {
	return l.repo.GetFullLink(shortLink)
}

func (l *LinkerService) SetShortLink(fullLink string) (string, error) {
	return l.repo.SetShortLink(fullLink)
}

func NewLinkerService(repo *repository.Repository) *LinkerService {
	return &LinkerService{repo: repo}
}
