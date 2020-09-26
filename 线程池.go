package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <- chan int, results chan <- int)  {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//创建5个goroutine
	//由于一开始通道里面没有值，所以管道处于阻塞状态
	for w := 1; w <= 5; w++ {
		go workers(w, jobs, results)
	}

	//由于是非阻塞通道，一次性向通道写入9个值，然后关闭通道
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	//循环读取管道中的值
	for a := 1; a <= 9; a++ {
		<-results
	}
}
