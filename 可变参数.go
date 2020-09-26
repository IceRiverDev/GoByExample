package main

import "fmt"

//这个函数接受任意数目的int类型的值作为参数
//输入的参数会放入变量a中，此时a的类型为切片
func sum(a ...int) (total int){
	for _, v := range a {
		total += v
	}

	return
}

func sum2(a ...int) (total int) {
	for _, v := range a {
		total += v
	}
	return
}

func main()  {
	fmt.Println(sum(1, 2, 3, 4, 5, 6)) //21
	fmt.Println(sum(1, 2, 3)) //6

	//可变参数除了可以自动构造为切片，输入的切片参数也可以解构为单独的一个个元素

	s := []int{1, 2, 3}
	fmt.Println(sum(s...)) // 6

	s2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum2(s2...)) // 21
}
