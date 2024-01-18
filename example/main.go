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
	a.Data.Print()

	p, err := client.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.Data.Print()

	s, err := client.Set("foo", "bar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.Data.Print()

	g, err := client.Get("foo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	g.Data.Print()

	k, err := client.Keys()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	k.Data.Print()

	d, err := client.Del("foo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d.Data.Print()

	push, err := client.Push("foo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	push.Data.Print()

	pop, err := client.Pop()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pop.Data.Print()

	e, err := client.Enque("foo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	e.Data.Print()

	de, err := client.Deque()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	de.Data.Print()
}
