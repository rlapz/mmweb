package service

import "github.com/rlapz/mmweb/src/repo"

type Service struct {
	repo repo.Repo
}

func New(r repo.Repo) *Service {
	return &Service{
		repo: r,
	}
}
