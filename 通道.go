package main

import "fmt"

/*
“不要通过共享内存来通信，而应该通过通信来共享内存” 这是一句风靡golang社区的经典语。

Go语言中，要传递某个数据给另一个goroutine(协程)，可以把这个数据封装成一个对象，然后把这个对象的指针传入某个channel中，
另外一个goroutine从这个channel中读出这个指针，并处理其指向的内存对象。Go从语言层面保证同一个时间只有一个goroutine能够访问channel里面的数据，
为开发者提供了一种优雅简单的工具，所以Go的做法就是使用channel来通信，通过通信来传递内存数据，使得内存数据在不同的goroutine中传递，而不是使用共享内存来通信。

通道的注意点
Channel通道在使用的时候，有以下几个注意点：

1.用于goroutine，传递消息的。

2.通道，每个都有相关联的数据类型, nil chan，不能使用，类似于nil map，不能直接存储键值对

3.阻塞：发送数据：chan <- data,是阻塞的，直到另一条goroutine，读取数据来解除阻塞 读取数据：data <- chan,也是阻塞的。直到另一条goroutine，写出数据解除阻塞。

4.本身channel就是同步的，意味着同一时间，只能有一条goroutine来操作。

最后：通道是goroutine之间的连接，所以通道的发送和接收必须处在不同的goroutine中。
*/

func main()  {
	//在go中使用make来创建一个通道
	//格式为`make(chan val-type)
	//通道的值必须指定类型
	messages := make(chan string)

	//使用 channel <- data语法来将一个数据传入通道
	//创建一个子goroutine
	go func(msg string) {
		messages <- msg
	}("ping")

	//我们在主goroutine里面接收子goroutine传递过来的值
	receivedMsg := <- messages
	fmt.Println(receivedMsg + " pong")//ping pong
}
