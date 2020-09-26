package main

import "fmt"

func main()  {
	//创建两个通道，一个通道用来接收发送数据
	//还有一个通道用来告知主goroutine，子gourotine数据全部接收成功
	//然后结束程序
	jobs := make(chan int, 5)
	done := make(chan bool)

	//子goroutine会通过循环来获取管道的数据，直到ok为false时，表明通道已被关闭
	//说明通道中的数据已经全部接收完毕，此时可以发送一个done的信号，
	//然后结束子goroutine
	go func() {
		for {
			//如果 `jobs` 已经关闭了，并且通道中所有的值都已经接收
			//完毕，那么 `more` 的值将是 `false`。当我们完成所有
			j, ok := <-jobs
			if ok {
				fmt.Println("received jobs", j)
			} else {
				fmt.Println("received all jobs")
				done <-true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

/*
输出结果：
sent job 1
sent job 2
sent job 3
sent all jobs
received jobs 1
received jobs 2
received jobs 3
received all jobs
 */
