package service

import "context"

type FriendService interface {
	Request(ctx context.Context, userID, friendID string) error
}
