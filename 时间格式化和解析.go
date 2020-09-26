package main

import (
	"fmt"
	"time"
)

//Go 支持通过基于描述模板的时间格式化和解析。
func main()  {
	p := fmt.Println

	// 这里是一个基本的按照 RFC3339 进行格式化的例子，使用对应模式常量。
	t := time.Now()
	tRFC3399 := t.Format(time.RFC3339)
	p(tRFC3399)

	//也可以根据RFC3399来解析一个时间值
	t1, _ := time.Parse(time.RFC3339, tRFC3399)
	p(t1)

	p(t.Format("3:04pm"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-7:00"))
	form := "3 04 PM"
	now2 := time.Now().Format(time.RFC3339)
	t2, e := time.Parse(form, now2)
	p(t2, e)

	//我们也可以使用标准的格式化字符串来提取时间值
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
	t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	//Parse函数在输入的时间格式不正确的时候会返回一个错误
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e2 := time.Parse(ansic, "8:41PM")
	if e2 != nil {
		panic(e2) //引发异常
	}

	/*
		2020-09-24T00:04:26+08:00
		2020-09-24 00:04:26 +0800 CST
		12:04am
		Thu Sep 24 00:04:26 2020
		2020-09-24T00:04:26.020285-7:00
		0001-01-01 00:00:00 +0000 UTC parsing time "2020-09-24T00:04:26+08:00": hour out of range
		2020-09-24T00:04:26-00:00
		panic: parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"

		goroutine 1 [running]:
		main.main()
		        /Users/liuyang/IdeaProjects/GoByExample/时间格式化和解析.go:37 +0xb24

	*/
}
