# 17Live Go Client

A Go client for interacting with the 17Live API, providing functionality to send messages, send pokes, share reactions, and follow and unfollow streamers.

## Installation

To install the 17Live Go client, use:

```bash
go get github.com/voyagen/17live
```

## Usage

Below is an example of how to use the api client to send a message:

```go
package main

import (
    "fmt"
    "time"

    "github.com/voyagen/17live/api"
)

func main() {
    // Initialize the client with a username and password
    apiclient, err := api.NewClient(
        "your-username",
        "your-password",
    )
    if err != nil {
        fmt.Println("Error creating client:", err)
        return
    }

    // Send poke request with roomid and message
    err = apiclient.SendMessage(12345678, "Hello World!")
    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }
    fmt.Println("Message sent successfully")
}
```

## Features

- Send messages
- Send poke requests
- Share reactions (Facebook, 17Live)
- Follow and Unfollow streamers

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
