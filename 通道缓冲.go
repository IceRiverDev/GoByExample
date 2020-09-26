package main

import "fmt"

/*
默认情况下，通道是无缓冲的，这意味着只有对应的接收(<-chan)通道准备好接收时，才允许进行发送(chan <-)。
也就是说，非缓冲通道是同步的(发送一个数据，接收一个数据)。如果通道发送一方的数据没有被读取，那么它会阻塞等待接收方数据的读取，直到数据被读取成功才会解除阻塞
而通道接收方在接收到数据之前也是一直处于阻塞状态，直到发送方发送数据以后才会解除阻塞。
可缓存通道与普通通道的区别在于:
1. 发送方不会等待接收方是否接收，它会先在缓冲通道发送一部分数据，缓冲通道被数据填满，那么它会处于阻塞状态，直到缓冲通道
再次有多余空间时，才会继续发送数据。
2. 接收方不会等待发送方一次性把数据发送完毕才进行接收，它会先去缓冲通道读取数据，直到缓冲通道里面的数据都被读取完毕，无数据可读时
它会处于阻塞状态，直到缓冲通道中继续有数据时，才会解除阻塞。
*/

func main() {
	//创建一个缓冲通道
	messages := make(chan string, 2)

	//由于通道是缓冲的，因此我们可以将这些值发送到通道中
	//而不需要相应的并发接收
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//使用非缓冲通道则不允许这种操作，会造成死锁
	messages1 := make(chan string)
	messages1 <- "ping"
	messages1 <- "pong"

	//fatal error: all goroutines are asleep - deadlock!
	//
	fmt.Println(<-messages1)
	fmt.Println(<-messages1)
}
