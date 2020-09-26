package main

import (
	"fmt"
	"os"
)

//使用os.Exit来立即进行给定状态的退出
//注意，不像例如C语言，Go不使用在main中返回一个整数来指明退出状态。
//如果你想以非零状态退出，那么你就要使用os.Exit。
//如果你使用go run来运行 exit.go，
//那么退出状态将会被go捕获并打印。

func main() {
	//当使用os.Exit时，defer将不会执行，所以这里的fmt.println不会被执行
	defer fmt.Println("!")

	os.Exit(3) //Process finished with exit code 3
	/*

	go run exit.go
	exit status 3

	编译成二进制文件再执行
	$ go build exit.go
	$ ./exit
	$ echo $?
	3
	 */
}


