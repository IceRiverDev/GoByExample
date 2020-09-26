package main

import (
	"fmt"
	"time"
)

//一般程序会有获取Unix时间的秒数，毫秒数，或者微秒数的需要。
//来看看如何用go来实现。

func main()  {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	//得到毫秒数没有现存的方法
	//需要手动转换一下
	millis := nanos / 1000000
	//millis
	fmt.Println(millis)
	//secs
	fmt.Println(secs)
	//nanos
	fmt.Println(nanos)

	//我们也可以将从1970年1月1号到现在的时间秒数转为相应的时间
	fmt.Println(time.Unix(secs, 0))
}
