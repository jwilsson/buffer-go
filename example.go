package main

import (
    "./buffer"

    "fmt"
    "os"
)

func main() {
    client := buffer.NewClient("1/ba8351857c1936233f25158026d4d5c5", nil)

    shares, err := client.Links.GetShares("http://bufferapp.com")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    }

    fmt.Println("Shares: ", shares)
}
