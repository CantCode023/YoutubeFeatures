Check out the main [documentation](../README.md)

<br>

# Table of Contents

- [Channel Model](#channel)
- [Video Model](#video)
- [Comment Model](#comment)

<br>

# Channel
- ## ChannelViewer
    - LikesGiven `integer`
    - DislikesGiven `integer`
    - Comments `integer`
    - ChannelName `string`

<br>

- ## ChannelCreator
    - Videos `[]Video`, a list of model [Video](#video)
    - Subscribers `integer`
    - ChannelName `string`

<br>

# Video
- Title `string`
- Description `string`
- Likes `integer`
- Dislikes `integer`
- Comments `[]Comment`, a list of model [Comment](#comment)

<br>

# Comment
- Author [ChannelViewer](#channelviewer)
- Message `string`
- Likes `integer`
- Dislikes `integer`
- Replies `[]Comment`, a list of model [Comment](#comment)