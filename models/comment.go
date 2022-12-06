package models

type Comment struct {
	Author ChannelViewer
	Message string
	Likes int
	Dislikes int
	Replies []Comment
}