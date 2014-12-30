package main

import (
    "./buffer"

    "fmt"
    "os"
)

func main() {
    client := buffer.NewClient("ACCESS TOKEN", nil)

    user, err := client.User.Get()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    }

    fmt.Println("User Timezone:", user.Timezone)
}
