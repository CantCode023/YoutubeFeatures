package models

import "fmt"

func TestModel() {
	fmt.Println("Model working!");
}

// // TODO: subscribe, unsubscribe

func (viewer *ChannelViewer) subcribe(creator ChannelCreator) (notification chan string) {
	creator.subscribers++;
	return notification;
}

func (viewer *ChannelViewer) unsubscribe(creator ChannelCreator, notification chan string) {
	creator.subscribers--;
	close(notification);
}

// // TODO: upload video

func (creator *ChannelCreator) uploadVideo(title, description string, notification chan string) *Video {
	newVideo := &Video{title, description, 0, 0, []Comment{}};
	notification <- fmt.Sprintf("%s uploaded a new video called \"%s\"", creator.channelName, newVideo.title);
	fmt.Println(<-notification)
	return newVideo;
}

// // TODO: like video, dislike video, comment video

func (viewer *ChannelViewer) likeVideo(video Video) int {
	video.likes++;
	viewer.likesGiven++;
	return video.likes;
}

func (viewer *ChannelViewer) dislikeVideo(video Video) int {
	video.dislikes++;
	viewer.dislikesGiven++;
	return video.dislikes;
}

func (viewer *ChannelViewer) commentVideo(video Video, commentMessage string) []Comment {
	newComment := &Comment{*viewer, commentMessage, 0, 0, []Comment{}};
	video.comments = append(video.comments, *newComment);
	return video.comments;
}

func (viewer *ChannelViewer) replyComment(comment Comment, replyMessage string) []interface{} {
	newReply := &Comment{*viewer, replyMessage, 0, 0, []Comment{}};
	comment.replies = append(comment.replies, *newReply);
	return []interface{}{comment.message, comment.replies}
}

// // TODO: like comment, dislike comment

func (viewer *ChannelViewer) likeCommentOrReply(comment Comment) Comment {
	comment.likes++;
	return comment;
}

func (viewer *ChannelViewer) dislikeCommentOrReply(comment Comment) Comment {
	comment.dislikes++;
	return comment;
}