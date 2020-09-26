package main

import (
	"fmt"
	"time"
)

/*
定时器是当你想要在未来某一刻执行一次时使用的
打点器则是当你想要在固定的时间间隔重复执行准备的。
这里是一个打点器的例子，它将定时的执行，直到我们将它停止。
 */
func main()  {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	//定时器和打点器的本质区别就是定时器规定了某一个时间点执行相应的行为
	//打点器则是周期性的执行相应的行为，间隔为我们定义的时间
}
