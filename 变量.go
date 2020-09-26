package main

import "fmt"

func main()  {
	var a = "initial" // 短声明变量
	fmt.Println(a)

	var b, c int = 1, 2 // 一次声明多个变量
	fmt.Println(b, c)

	d, e := 1, 2 // 短声明变量
	fmt.Println(d, e)

	var f  = true // 类型推导
	fmt.Println(f)

	/*
	声明后没有赋值，变量会被初始化为该类型相应的零值
	int `0`
	string ""
	[]int nil
	 */

	var g string // 被初始化为""空字符串
	g = "hello world"
	fmt.Println(g)

	var _ []int // 被初始化为nil值
	/*
	h[0] = 1
	如果切片以这种形式声明，并且直接赋值，那么程序会报错。
	因为切片由三个元素组成，指向底层数组数据结构的地址，长度和容量。
	由于此时切片底层并没有被初始化的数组，因此赋值会产生异常，消除
	这个异常的办法是改为make([]int, 0), 或者使用append
	 */

}
