package main

import "fmt"

/*
一个函数如果不断调用自身，那么就说这个函数是递归函数
递归函数必须有一个基线条件来终止递归，否则函数会不断调用自身
最终导致堆栈溢出。
 */

func totalValue(num int) int {
	if num == 1 {
		return 1 //函数出口（基线条件）如果没有这个条件，那么函数会无限调用自己而无法退出。
	} else {
		return num + totalValue(num - 1) // 调用自身
	}
}

//通常来说，递归的函数可以使用循环方案来替换
func totalValue2(num int) (sum int) {
	for i := 1;  i <= num; i++ {
		sum += i
	}
	return
}

func main()  {
	total := totalValue(10)
	fmt.Println(total) // 55

	total2 := totalValue2(10)
	fmt.Println(total2) // 55
}
