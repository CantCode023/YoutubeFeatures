<h1>Table Of Contents</h1>

- [Setup](#setup)
- [Channels](#channels)

<br>

# Setup
```go
package main

import (
    "fmt"
    "example.com/pkgs/model"
)

func main() {
    fmt.Println("Welcome!");
    model.TestModel();
}
```

# Channels

## Create a viewer account
```go
package main

import (
    "fmt"
    "example.com/pkgs/model"
)

func main() {
    viewerChannel := &model.ChannelViewer{0,0,0,"Channel Name Here"};
    fmt.Println(viewerChannel);
}
```
We used `&model.ChannelViewer` to set `viewerChannel` as a pointer.
You can learn more about `ChannelViewer` in the [documentation](./models/README.md#table-of-contents)

<br>

- ### Create a