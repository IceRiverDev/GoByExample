package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。
对于更加复杂的情况，我们可以使用一个互斥锁来在go协程间安全的访问数据。
 */

func main()  {
	var state = make(map[int]int)

	var mtx = &sync.Mutex{}

	var ops uint64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mtx.Lock()
				total += state[key]
				mtx.Unlock()
				atomic.AddUint64(&ops, 1)
			}
			runtime.Gosched()
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mtx.Lock()
				state[key] = val
				mtx.Unlock()
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

	mtx.Lock()
	fmt.Println("state:", state)
	mtx.Unlock()
}
