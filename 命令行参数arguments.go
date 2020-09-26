package main

import (
	"fmt"
	"os"
)

//命令行参数是指定程序运行的一个常见方式。
//例如：go run hello.go,go使用了run和hello.go两个参数

func main() {
	//参数里面保存了程序名和紧跟着参数
	argsWithProg := os.Args

	//去除程序名，只显示参数
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg) //[./命令行参数arguments 1 2 3]
	fmt.Println(argsWithoutProg) //[1 2 3]

	fmt.Println(os.Args[3])
}
