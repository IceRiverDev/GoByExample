package main

import "fmt"

type Foo struct {
	C chan string
}

func main() {
	//c1 := make(chan string)
	f := Foo{
		make(chan string),
	}

	go func() {
		f.C <- "hello world"
	}()

	fmt.Println(<-f.C)
}
