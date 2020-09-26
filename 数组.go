package main

import "fmt"

func main()  {
	//数组是固定长度的，且长度不可更改
	//数组是值类型数据，不是引用类型

	var a [5]int // 会被初始化，且初始值为元素对应类型的零值，这里为0
	fmt.Println(a) // [0,0,0,0,0]

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	//a[5] = 6 // error，这里超过了数组长度，所以会产生异常

	fmt.Println(a) //[1 2 3 4 5]

	//获取数组长度 len()
	fmt.Println(len(a)) // 5

	// 二维数组
	var aa [2][3]int // 定义外部数组长度为2，内部数组长度为3的二维数组
	aa = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(aa) // [[1 2 3] [4 5 6]]

	// 遍历内部数组元素
	for i := 0; i < len(aa); i++ {
		for j := 0; j < len(aa[0]); j++ {
			fmt.Println(aa[i][j])
		}
	}

	// 自动推测长度
	var bb = [...]string{"hello", "world"}
	fmt.Println(bb)
}
