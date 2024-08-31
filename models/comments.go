package models

type Comment struct {
	PostID      string `json:"postID"`
	ID          string `json:"id"`
	Comment     string `json:"comment"`
	UserID      string `json:"userID"`
	UserName    string `json:"userName"`
	CommentTime string `json:"commentTime"`
}
