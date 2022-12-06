package models

type ChannelViewer struct {
	LikesGiven       int
	DislikesGiven    int
	Comments         int
	ChannelName string
}

type ChannelCreator struct {
	Videos []Video
	Subscribers int
	ChannelName string
}