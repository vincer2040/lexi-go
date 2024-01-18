package main

import (
	"fmt"
	"os"

	lexigo "github.com/vincer2040/lexi-go/pkg/lexi-go"
)

func main() {
    client := lexigo.New("127.0.0.1:6969")
    err := client.Connect()
    defer client.Close()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    a, err := client.Auth("root", "root")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(a)

    p, err := client.Ping()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(p)
}
