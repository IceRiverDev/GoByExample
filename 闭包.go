package main

import "fmt"

/*
通常一个函数的局部变量在函数返回以后，会被销毁
外部无法访问该局部变量，通常一个变量都会有生命周期。
闭包的出现可以使局部变量的生命周期延长，即使函数返回了，
我们依然有办法访问该变量
 */

func incrementNum(num int) func() int {
	/*
	num变量为该函数的局部变量,通常在该函数返回以后，这个变量就被销毁了。
	但是由于我们定义了另外一个匿名函数来访问变量，因此即使incrementNum函数返回了，
	但是在匿名函数中依然存在对该变量的引用，因此它不会被销毁。
	 */
	return func() int {
		num++
		return num
	}
}

//我们定义个函数，返回两个函数，一个函数用来设置值，一个函数用来获取值
func SetGetNum(num int) (func(), func() int) {
	return func() { num++ }, func() int { return num }
}

func main()  {
	f1 := incrementNum(0)
	fmt.Println(f1()) //1
	fmt.Println(f1()) //2
	fmt.Println(f1()) //3

	fmt.Println()
	set, get := SetGetNum(0)
	set()
	fmt.Println(get()) //1

	set()
	fmt.Println(get()) //2

	set()
	fmt.Println(get()) //3
}
