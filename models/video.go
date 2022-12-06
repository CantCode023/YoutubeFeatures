package models

type Video struct {
	Title string
	Description string
	Likes int
	Dislikes int
	Comments []Comment
}