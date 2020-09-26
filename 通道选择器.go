package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <-"one"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <-"two"
	}()

	count := 0
	//循环两次，来获取两个通道中的数据
	//因为函数2的睡眠时间为2秒，因此它的数据晚于函数1发送
	//所以select优先接收到到函数1的值，在接收到函数2的值
	for {
		count++
		select {
		case msg1 := <-c1:
			fmt.Println("received",msg1)
		case msg2 := <-c2:
			fmt.Println("received",msg2)
		}
		if count < 2 {
			continue
		} else {
			break
		}
	}

	//如果说函数1和函数2的睡眠时间都为1秒，那么谁的数据先被select接收到呢？
	//如果两个通道的数据同时到达，那么select的选择是随机的。
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received",msg1)
		case msg2 := <-c2:
			fmt.Println("received",msg2)
		}
		break
	}
	// 多执行几次函数会发现有的时候是one，有的时候是two，这就是select的选择机制

}
