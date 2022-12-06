Check out the main [documentation](../README.md)

<br>

# Table of Contents

- [Channel Model](#channel)
- [Video Model](#video)
- [Comment Model](#comment)

<br>

# Channel
- ## ChannelViewer
    - likesGiven integer
    - dislikesGiven integer
    - comments integer
    - channelName string

<br>

- ## ChannelCreator
    - videos []Video, a list of model [Video](#video)
    - subscribers integer
    - channelName string

<br>

# Video
- title string
- description string
- likes integer
- dislikes integer
- comments []Comment, a list of model [Comment](#comment)

<br>

# Comment
- author [ChannelViewer](#channelviewer)
- message string
- likes integer
- dislikes integer
- replies []Comment, a list of model [Comment](#comment)