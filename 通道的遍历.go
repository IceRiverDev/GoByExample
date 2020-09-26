package main

import "fmt"

func main()  {
	queue := make(chan string, 2)

	go func() {
		queue <-"one"
		queue <-"two"
		close(queue)
	}()

	//range会遍历通道queue，当迭代两次以后，range会自动去判断通道是否已被关闭
	//因此我们不用像之前的例子那样使用二值形式来判断通道是否已被关闭
	//当range检测到通道被关闭以后，循环会终止，如果通道没有关闭
	//那么将会继续阻塞迭代执行for循环
	for elem := range queue {
		fmt.Println("received data", elem)
	}
}
