package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器对象，时间为2秒
	timer1 := time.NewTimer(time.Second * 2)

	//timer1对象的字段C的值为通道，2秒以后通道的解除阻塞
	<-timer1.C
	fmt.Println("Timer 1 expired")

	//创建timer2对象
	timer2 := time.NewTimer(time.Second)

	//创建一个goroutine，1秒以后读取到timer2管道的值，解除阻塞
	//打印字符串
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	//立即停止定时器，如果停止成功返回true，否则返回false
	//这里返回true，因此字符串会被打印
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
