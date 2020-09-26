package main

import "fmt"

func f(from string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//一般只有函数f执行完毕以后才会执行后面的代码
	//f("direct")

	// 等函数f执行完了以后开始执行
	fmt.Println("start run...")

	//我们可以使用goroutine来进行异步执行
	//f函数可以独立于主线程单独运行，而不会影响后续代码的执行
	//golang中使用go关键字来创建一个新的goroutine
	go f("goroutine")

	//可以直接为匿名函数指定一个go协程
	go func(msg string) {
		fmt.Println(msg)
	}("Let's go!")

	/*
	由于程序开始运行时，main函数是一个主goroutine，如果main函数先返回，
	那么及时子goroutine没有执行完，程序也会退出，因此为了等子goroutine
	执行完毕，我们需要让主goroutine阻塞等待，直到子goroutine执行完毕
	 */
	fmt.Scanln() //调用标准输入io，这里主要是起到一个阻塞作用,在终端随便输入一个值使其退出

	fmt.Println("main func done...")
}
