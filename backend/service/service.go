package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/repo"
)

type Service struct {
	signMethod *jwt.SigningMethodHMAC
	signKey    []byte
	loginExp   time.Duration
	repo       repo.Repo
}

func New(cfg *config.Config, r repo.Repo) *Service {
	return &Service{
		signMethod: cfg.JwtSignMethod,
		signKey:    cfg.JwtSignatureKey,
		loginExp:   cfg.LoginExp,
		repo:       r,
	}
}
