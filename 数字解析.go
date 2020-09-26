package main

import (
	"fmt"
	"strconv"
)

//使用parseFloat解析浮点数，第二个参数代表位数
func main()  {
	f, _ := strconv.ParseFloat("1.2345", 64)
	fmt.Printf("%T %v\n", f, f) //float64 1.2345

	f2, _ := strconv.ParseFloat("1.234", 32)
	fmt.Printf("%T %v\n", f2, f2) //float64 1.2339999675750732

	//ParseInt用来解析整型数值
	//第二个参数0代表自动推断字符串所表示的数字的进制
	//64表示返回的整数值是以64位存储的
	i1, _ := strconv.ParseInt("2", 0, 64)
	fmt.Printf("%T %v\n", i1, i1)//int64 2

	i2, _ := strconv.ParseInt("2", 0, 32)
	fmt.Printf("%T %v\n", i2, i2)//int64 2

	//自动识别十六进制数
	i3, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Printf("%T %v\n", i3, i3)//int64 456

	// `ParseUint` 也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Printf("%T %v\n", u, u)//uint64 789

	// `Atoi` 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k) //135

	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat") //无法正常解析。返回一个错误
	fmt.Println(e) //strconv.Atoi: parsing "wat": invalid syntax

}
