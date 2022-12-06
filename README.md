<h1>! IMPORTANT !</h1>
This is a beginner mini project I made to understand channels, functions, pointers, etc. in depth. If you see any stuffs that could be improved, feel free to let me know on my Discord Blue Duck#8344 as I'm still learning GoLang and I'm looking forward to improve my skills. Thanks!

<br>

<h1>Table Of Contents</h1>

- [Setup](#setup)
- [Creating Channels](#creating-channels)
- [Subscribing and Unsubscribing](#subscribing-and-unsubscribing)
- [Uploading Video](#uploading-video)
- [Liking and Disliking](#liking-and-disliking)
- [Commenting and Replying](#commenting-and-replying)
- [Liking, Disliking Comments and Replies](#liking-disliking-comments-and-replies)

<br>

# Setup
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    fmt.Println("Welcome!");
    models.TestModel();
}
```

<br>

# Creating Channels

## Create a Viewer Account
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{
        LikesGiven: 0,
        DislikesGiven: 0,
        Comments: 0,
        ChannelName: "Channel Name",
    }
    // But for simplicity, we'll do a one-liner. Learn more about the models from the documentation.
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    fmt.Println(*viewerChannel);
}
```
We used `&models.ChannelViewer` to set `viewerChannel` as a pointer.
You can learn more about `ChannelViewer` in the [documentation](./models/README.md#channelviewer)

<br>

## Create a Creator Account
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name"};
    creatorChannel := &models.ChannelCreator{[]models.Video{}, 0, "Channel Name"}
    fmt.Println(*viewerChannel);
    fmt.Println(*creatorChannel);
}
```
You can learn more about `ChannelCreator` in the [documentation](./models/README.md#channelcreator)

<br>

# Subscribing and Unsubscribing
To do this, you must use a viewer channel to subscribe to a creator channel.

## Subscribe
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")
    fmt.Println(*viewerChannel);
    fmt.Println(*creatorChannel);
    notification := viewerChannel.Subscribe(creatorChannel);
}
```
The `Subscribe` method will return a string channel that can later be used as a notification whenever the subscribed creator uploads a new video.<br>
You can learn more about channels [here.](https://gobyexample.com/channels)

<br>

## Unsubscribe
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")
    fmt.Println(*viewerChannel);
    fmt.Println(*creatorChannel);
    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);
    viewerChannel.Unsubscribe(creatorChannel, notification);
    fmt.Println("Unsubscribed " + creatorChannel.channelName);
}
```
With the `Unsubscribe` method, you have to pass in notification as an argument to `viewerChannel.Unsubscribe`, this is to close the notification channel so no further notification will be received.<br>
You can learn more about closing channels [here.](https://gobyexample.com/closing-channels)

<br>

# Uploading Video
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")

    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
}
```
`creatorChannel.UploadVideo` accepts three arguments: `title`, `description`, and the notification channel so that it can send pings or notifications to the receiver.<br>
The receiver will be handled internally (it will print the received value) so no additional code is required.<br>
`creatorChannel.UploadVideo` will return a [Video](./models/README.md#video) type. Learn more about channels [here.](https://gobyexample.com/channels)

<br>

# Liking and Disliking
This method can only be called within a [Viewer Channel](./models/README.md#channelviewer) and can only like or dislike [Videos.](./models/README.md#video)

## Like
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")

    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
    viewerChannel.LikeVideo(firstVideo);
}
```
`viewerChannel.LikeVideo` will return an integer, which is the video's amount of likes.

<br>

## Dislike
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")

    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
    viewerChannel.DislikeVideo(firstVideo);
}
```
`viewerChannel.DislikeVideo` will return an integer, which is the video's amount of dislikes.

<br>

# Commenting and Replying
This method can only be called within a [Viewer Channel](./models/README.md#channelviewer) and can only comment on [Videos](./models/README.md#video) and reply to a [Comment.](./models/README.md#comment)

## Comment
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")

    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
    viewerChannel.LikeVideo(firstVideo);

    commentOne := viewerChannel.CommentVideo(firstVideo, "Comment Message")
}
```
`viewerChannel.CommentVideo` will return a [Comment Object.](./models/README.md#comment)

<br>

## Reply

```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")

    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
    viewerChannel.likeVideo(firstVideo);

    commentOne := viewerChannel.CommentVideo(firstVideo, "Comment Message");

    replyOne := viewerChannel.ReplyComment(commentOne, "Reply Message");
}
```
`viewerChannel.ReplyComment` will return a [Comment Object.](./models/README.md#comment)<br>
Note that reply type is [Comment](./models/README.md#comment) which means you can stack reply with more replies. 

<br>

# Liking and Disliking Comments or Replies

## Liking Comments or Replies
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")

    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
    viewerChannel.likeVideo(firstVideo);

    commentOne := viewerChannel.CommentVideo(firstVideo, "Comment Message");

    replyOne := viewerChannel.ReplyComment(commentOne, "Reply Message");
    viewerChannel.LikeCommentOrReply(commentOne);
    viewerChannel.LikeCommentOrReply(replyOne);
    fmt.Println(commentOne.Likes, replyOne.Likes);
}
```
`viewerChannel.LikeCommentOrReply` will return the comment or reply passed in from the argument.

## Disliking Comments or Replies
```go
package main

import (
    "fmt"
    "example.com/pkgs/models"
)

func main() {
    viewerChannel := &models.ChannelViewer{0,0,0,"Channel Name Here"};
    creatorChannel := &models.ChannelCreator([]models.Video{}, 0, "Channel Name Here")

    notification := viewerChannel.Subscribe(creatorChannel);
    fmt.Println("Subscribed to " + creatorChannel.channelName);

    firstVideo := creatorChannel.UploadVideo("Title", "Description", notification);
    viewerChannel.likeVideo(firstVideo);

    commentOne := viewerChannel.CommentVideo(firstVideo, "Comment Message");

    replyOne := viewerChannel.ReplyComment(commentOne, "Reply Message");
    viewerChannel.DislikeCommentOrReply(commentOne);
    viewerChannel.DislikeCommentOrReply(replyOne);
    fmt.Println(commentOne.Dislikes, replyOne.Dislikes);
}
```
`viewerChannel.DislikeCommentOrReply` will return the comment or reply passed in from the argument.

# Conclusion

Congratulations! You have read the documentation all the way through. Now explore the source code and please tell me what I could improve.