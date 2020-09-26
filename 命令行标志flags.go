package main

import (
	"flag"
	"fmt"
)

//go提供了一个flag包来支持基本的命令行标志解析

func main() {
	//第一个参数为标志名， 第二个参数为默认值，第三个参数为使用说明
	//需要注意这里返回的是指针值，要获取值需要通过*来获取
	//如果命令行省略了该标志，则这个标志的值会被设置为默认值
	wordPtr := flag.String("word", "foo", "a string")

	numPtr  := flag.Int("number", 42, "a integer")

	boolPtr := flag.Bool("fork", false, "a bool")

	//预先声明一个变量，然后将变量传入StringVar
	//这样命令行传入的值会被赋值给该变量
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//当所有标志声明以后，需要调用Parse来完成命令行参数解析
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("number:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args()) //返回非标志命令行参数
	/*
	./命令行标志flags  -word=opt -number=100  -fork=true 1 2 3
	需要注意的是
	1.尾随的位置参数可以出现在任何标志后面。
	2.注意，flag 包需要所有的标志出现位置参数之前（ 否则，这个标志将会被解析为位置参数）。
	3.使用 -h 或者 --help 标志来得到自动生成的这个命 令行程序的帮助文本。
	4.如果你提供一个没有使用 flag 包指定的标志，程序会输出一 个错误信息，并再次显示帮助文本。
	 */
}
