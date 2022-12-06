package models

type Video struct {
	title string
	description string
	likes int
	dislikes int
	comments []Comment
}