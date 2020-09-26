package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("workding...")
	time.Sleep(2 * time.Second)
	fmt.Println("done...")
	done <-true
}

func main() {
	done := make(chan bool)
	go worker(done)
	if <-done {
		fmt.Println("main over...")
	}
}
