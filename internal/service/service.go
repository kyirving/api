package service

import "api/internal/repository"

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Ping() error {
	return s.repo.Ping()
}
