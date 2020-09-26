package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	//我们将使用一个无符号整型数来表示（永远是正整数）这个计数器。
	var ops uint64 = 0

	// 为了模拟并发更新，我们启动50个Go协程，对计数
	// 器进行一次加一操作。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	//等待一秒，让ops的自增操作执行一会
	time.Sleep(time.Second)

	// 为了在计数器还在被其它 Go 协程更新时，安全的使用它，
	// 我们通过 `LoadUint64` 将当前值的拷贝提取到 `opsFinal`
	// 中。和上面一样，我们需要给这个函数所取值的内存地址 `&ops`
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops", opsFinal)
}
