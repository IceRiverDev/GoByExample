package main

import (
	"fmt"
	"time"
)

func main()  {
	p := fmt.Println

	//获取当前时间
	now := time.Now()
	p(now) //2020-09-23 22:56:07.423575 +0800 CST m=+0.000046539

	then := time.Date(2020, 10, 8, 20, 34, 50, 651387237, time.UTC)
	p(then) //2020-10-08 20:34:50.651387237 +0000 UTC

	//可以提取时间对象的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Date())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	/*
	output:
	2020
	October
	2020 October 8
	20
	34
	50
	651387237
	UTC
	 */

	//打印星期一到星期日(Weekday)
	p(then.Weekday())//Thursday

	//时间与时间之间也可以进行比较
	p(then.Before(now)) // false
	p(then.After(now)) // true
	p(then.Equal(now)) // false

	//方法Sub用来返回一个Duration来表示两个时间之间的间隔时间
	diff := then.Sub(now)
	p(diff)

	//可以计算不同单位下的时间长度值
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//可以使用Add来将时间后移一个时间间隔
	//或者使用一个-来将时间前移一个时间间隔
	p(then.Add(diff)) //2020-10-24 01:35:01.732571474 +0000 UTC
	p(then.Add(-diff)) //2020-09-23 15:34:39.570203 +0000 UTC
}
