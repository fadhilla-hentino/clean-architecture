package service

import (
	"context"

	"fadhilla-hentino/clean-architecture/repository"
	"go.uber.org/zap"
)

type service struct {
	repo   repository.FriendsRepo
	logger *zap.Logger
}

func NewService(repo repository.FriendsRepo, logger *zap.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) Request(ctx context.Context, userID, friendID string) error {
	return s.repo.Request(ctx, userID, friendID)
}
