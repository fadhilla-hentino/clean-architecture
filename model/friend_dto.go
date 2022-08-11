package model

type FriendRequestReq struct {
	UserID   string `json:"userID"`
	FriendID string `json:"friendID"`
}
