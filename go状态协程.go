package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
在前面的例子中，我们用互斥锁进行了明确的锁定来让共享的state跨多个Go协程同步访问。
另一个选择是使用内置的Go协程和通道的同步特性来达到同样的效果。这个基于通道的方法和Go通过通信来共享内存，
以及确保每块数据被单独的Go协程所拥有的思路是一致的。
 */

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64 = 0
	var writeOps uint64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <-state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <-true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <-read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 用相同的方法启动10个写操作。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <-write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让Go协程们跑1s。
	time.Sleep(time.Second)

	// 最后，获取并报告 `ops` 值。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}
