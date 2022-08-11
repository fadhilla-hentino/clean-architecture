package model

// constant for friendship relation
const (
	// NotFriendStatus is a status for user that doesn't have friendship relation
	NotFriendStatus = iota // 0

	// OutgoingStatus is a status for outgoing users
	OutgoingStatus // 1

	// IncomingStatus is a status for incoming users
	IncomingStatus // 2

	// FriendsStatus is a status for two users that already have friends relationship
	FriendsStatus // 3
)

// Friends representation Friends table
type Friends struct {
	UserID       string
	FriendID     string
	Relationship int
}
