package main

import "fmt"

const perl = "larry wall"

const ruby string = "matz"

func printNumber(n float64)  {
	fmt.Println(n)
}

/*
数值型常量没有确定的类型，直到被给定
当上下文需要时，比如变量赋值或者函数调用
 */
const num = 5000 // 整数类型

func main() {

	fmt.Println(perl)

	fmt.Println(ruby)

	printNumber(num) // 由于数值型常量的类型在给定的上下文才会被确认
	                 // 因此这里不会产生编译错误

	const num2 = 1
	fmt.Println(int64(num2)) // 手动进行类型转换

}
