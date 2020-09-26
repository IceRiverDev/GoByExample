package main

import (
	"fmt"
	"time"
)

//我们可以在程序中定义超时机制来处理一些超出预期执行时间的行为

func main()  {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	//2秒以后这里才会接收到发送过来的数据
	case res := <-c1:
		fmt.Println(res)
		//定义一个超时，超时时间为1秒
		//由于超时分支比接收数据分支的时间要短
		//因此这里超时分支会被提前执行
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	select {
	//1秒以后这里接收到发送过来的数据
	case res := <-c1:
		fmt.Println(res)
		//定义一个超时，超时时间为3秒
		//由于超时分支比接收数据分支的时间要长
		//因此这里超时分支不会被执行
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 3")
	}
	//output:
	//timeout 1
	//result 1
}

