package main

import (
	"fmt"
	"sync"
	"time"
)

//我们可以通过waitGroup来等待子goroutine全部运行完毕以后
//再退出主goroutine
func Worker(id int, wg *sync.WaitGroup)  {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
	wg.Done()
}

var wg sync.WaitGroup
func main() {
	//循环5次，创建5个goroutine，然后为wg递增waitGroup的计数器
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		//由于每个goroutine会睡眠1秒，因此一开始会暂时阻塞一秒
		//所以暂时只会打印第一条字符串，当所有goroutine苏醒以后
		//开始抢夺资源，因此会随即打印第二条字符串
		go Worker(i, &wg)
	}
	//阻塞，直到waitGroup的计数恢复为0，即所有子goroutine全部完成
	wg.Wait()
}
