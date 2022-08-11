package repository

import (
	"context"
)

type FriendsRepo interface {
	Request(ctx context.Context, userID, friendID string) error

	// TODO
	//Accept(ctx context.Context, userID, friendID string, relation int) error
	//Reject(ctx context.Context, userID, friendID string) error
	//ListOfFriend(ctx context.Context, userID string) ([]model.Friends, error)
	//Get(ctx context.Context, userID, friendID string) (model.Friends, error)
}
