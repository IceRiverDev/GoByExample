package main

import (
	"fmt"
	"time"
)

/*
通道默认定义为双向通道，既可以接收数据也可以发送数据
但是我们也可以定义单向通道，只负责接收的通道或者只负责发送通道
但是在实际应用中，单向通道并没有任何意义，因为如果只能发送或者接收
那么就无法让不同的协程之间通过通道进行通信。通常会在定义函数时，
定义一个单向通道的形参，这样当我们传入一个双向通道参数，
那么这个通道就只能进行单向使用.
 */

func ping(pings chan <- string, msg string)  {
	pings <-msg
}

func pong (pongs <- chan string)  {
	fmt.Println(<-pongs + " pong")
}

func hello(h chan <-string, msg string) {
	h <-msg
}

func world(h <-chan string, w chan <- string) {
	msg := <-h
	w <- msg + " world"
}

func main() {
	//管道ch可以让协程ping和pong之间进行通信
	ch := make(chan string)
	go ping(ch, "ping")
	go pong(ch)

	//虽然这里的管道可以正常运行，但是它只是在主线程中运行
	//并没有起到goroutine之间进行通信的作用，因此意义不大
	h := make(chan string, 1)
	w := make(chan string, 1)
	hello(h, "hello")
	world(h, w)
	fmt.Println(<-w)
	time.Sleep(2 * time.Second)
}
