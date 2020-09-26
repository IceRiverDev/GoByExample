package main

import (
	"fmt"
)

/*
常规的通过通道发送和接收数据是阻塞的。
然而，我们可以 使用带一个default子句的select来实现非阻塞的发送、接收，
甚至是非阻塞的多路select。
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <-msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//我们可以实现一个多路的非阻塞的选择器，这样即使每一路没有成功发送或接收数据
	//也不会出现死锁，而是会执行default分支语句，避免了死锁的情况
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
