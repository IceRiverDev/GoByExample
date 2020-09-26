package main

import "fmt"

func main() {
	count := 0
	// 相当于while count < 10 ...
	for ; count < 10; {
		count++
		fmt.Println(count)
	}

	fmt.Println()
	count = 0
	for i := 0; i < 10; i++ {
		count++
		fmt.Println(count)
	}

	fmt.Println()
	count = 0
	// 无限循环，相当于while true ...
	for {
		count++
		if count > 10 {
			break
		}
		fmt.Println(count)
	}

	/*
		循环控制关键字：
		continue 跳过该次循环后面的代码，执行下一次新的循环
		break 终止循环

		循环控制标签
		go语言支持循环控制标签
	*/

	count = 0

Loop:
	for count < 10 {
		count++

		if count == 5 {
			continue Loop
		} else if count == 8 {
			break Loop
		}
		fmt.Println(count)
	}
}
