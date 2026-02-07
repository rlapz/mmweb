package service

import "github.com/rlapz/mmweb/repo"

type Service struct {
	repo repo.Repo
}

func New(r repo.Repo) *Service {
	return &Service{
		repo: r,
	}
}
