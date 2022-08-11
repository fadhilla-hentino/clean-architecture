package postgres

import "context"

type repo struct {
}

func NewRepo() *repo {
	return &repo{}
}

func (repo *repo) Request(ctx context.Context, userID, friendID string) error {
	return nil
}
