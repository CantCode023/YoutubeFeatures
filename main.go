package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{
		LikesGiven: 5,
		DislikesGiven: 0,
		Comments: 10,
		ChannelName: "PSGE",
	};
	creatorChannel := &models.ChannelCreator{
		Videos: []models.Video{},
		Subscribers: 0,
		ChannelName: "PSGE_CREATOR",
	};
	fmt.Println(*viewerChannel);
	fmt.Println(*creatorChannel);

	notification := viewerChannel.Subcribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.ChannelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
	fmt.Println(creatorChannel)
    viewerChannel.LikeVideo(firstVideo);
	fmt.Println(firstVideo);

    commentOne := viewerChannel.CommentVideo(firstVideo, "Comment Message");
	fmt.Println(commentOne);

	replyOne := viewerChannel.ReplyComment(commentOne, "Reply Message");
	viewerChannel.LikeCommentOrReply(replyOne);
	viewerChannel.DislikeCommentOrReply(commentOne);
	fmt.Println(replyOne.Likes, commentOne.Dislikes);
	// viewerChannel.Unsubscribe(creatorChannel, notification);

}