package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
有时候，我们希望Go能智能的处理Unix信号。
例如，我们希望当服务器接收到一个SIGTERM信号时能够自动关机，
或者一个命令行工具在接收到一个SIGINT信号 时停止处理输入信息。
这里讲的就是在Go中如何通过通道来处理信号。
 */

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//注册这个给定的通道用于接收特定的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//子goroutine会执行一个阻塞信号等待的啊哦做，当它得到一个信号时
	//解除阻塞继续往下执行，然后发送一个bool值给通道done
	go func() {
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	//等待子goroutine发送的值，一旦接收到值就解除阻塞退出程序
	<-done
	fmt.Println("exiting")
}
