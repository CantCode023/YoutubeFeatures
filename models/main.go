package models

import "fmt"

func TestModel() {
	fmt.Println("Model working!");
}

// // TODO: subscribe, unsubscribe

func (viewer *ChannelViewer) Subcribe(creator *ChannelCreator) (chan string) {
	creator.Subscribers++;
	notification := make(chan string, 1);
	return notification;
}

func (viewer *ChannelViewer) Unsubscribe(creator *ChannelCreator, notification chan string) {
	creator.Subscribers--;
	close(notification);
}

// // TODO: upload video

func (creator *ChannelCreator) UploadVideo(title, description string, notification chan string) *Video {
	newVideo := &Video{title, description, 0, 0, []Comment{}};
	creator.Videos = append(creator.Videos, *newVideo)
	notification <- fmt.Sprintf("%s uploaded a new video called \"%s\"", creator.ChannelName, newVideo.Title);
	fmt.Println(<-notification);
	return newVideo;
}

// // TODO: like video, dislike video, comment video

func (viewer *ChannelViewer) LikeVideo(video *Video) int {
	video.Likes++;
	viewer.LikesGiven++;
	return video.Likes;
}

func (viewer *ChannelViewer) DislikeVideo(video *Video) int {
	video.Dislikes++;
	viewer.DislikesGiven++;
	return video.Dislikes;
}

func (viewer *ChannelViewer) CommentVideo(video *Video, commentMessage string) *Comment {
	newComment := &Comment{*viewer, commentMessage, 0, 0, []Comment{}};
	video.Comments = append(video.Comments, *newComment);
	return newComment
}

func (viewer *ChannelViewer) ReplyComment(comment *Comment, replyMessage string) *Comment {
	newReply := &Comment{*viewer, replyMessage, 0, 0, []Comment{}};
	(*comment).Replies = append((*comment).Replies, *newReply);
	return newReply
}

// // TODO: like comment, dislike comment

func (viewer *ChannelViewer) LikeCommentOrReply(comment *Comment) *Comment {
	comment.Likes++;
	return comment;
}

func (viewer *ChannelViewer) DislikeCommentOrReply(comment *Comment) *Comment {
	comment.Dislikes++;
	return comment;
}