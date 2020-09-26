package main

import "fmt"

func main()  {
	if 10 < 100 {
		fmt.Println(true)
	}

	// 这里定义的num是局部变量，只属于该if语句，外部无法访问
	if num := 10; num > 8 {
		fmt.Println(num)
	}

}
