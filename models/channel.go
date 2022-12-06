package models

type ChannelViewer struct {
	likesGiven       int
	dislikesGiven    int
	comments         int
	channelName string
}

type ChannelCreator struct {
	videos []Video
	subscribers int
	channelName string
}