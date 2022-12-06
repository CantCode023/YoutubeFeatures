package models

type Comment struct {
	author ChannelViewer
	message string
	likes int
	dislikes int
	replies []Comment
}