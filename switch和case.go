package main

import (
	"fmt"
	"time"
)

func main()  {
	num := 2

	switch num {
	case 2:
		fmt.Println("number is 2")
	case 3:
		fmt.Println("number is 3")
	case 4:
		fmt.Println("number is 4")
	default:
		fmt.Println("number unknown..")
	}

	// 在case子句中，可以使用逗号分隔多个匹配条件
	switch t := time.Now().Weekday(); t {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch true {
	case t.Hour() < 12:
		fmt.Println("morning...")
	default:
		fmt.Println("afternoon...")
	}

	// 类型开关用法

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("I am string")
		case int:
			fmt.Println("I am int")
		case bool:
			fmt.Println("I am bool")
		default:
			fmt.Println("unknow type")
		}
	}

	whoAmI("hello world")
	whoAmI(100)
}
