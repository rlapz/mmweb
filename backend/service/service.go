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

	repoAuth repo.Auth
	repoTodo repo.Todo
	repoUser repo.User
}

func New(cfg *config.Config, authRepo repo.Auth, todoRepo repo.Todo, userRepo repo.User) *Service {
	return &Service{
		signMethod: cfg.JwtSignMethod,
		signKey:    cfg.JwtSignatureKey,
		loginExp:   cfg.LoginExp,
		repoAuth:   authRepo,
		repoTodo:   todoRepo,
		repoUser:   userRepo,
	}
}
