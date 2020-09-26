package main

import "os"

/*
通常程序中出现了不该出现的错误，那么就会引发一个panic(恐慌)。
一旦引发恐慌，那么它会沿着调用栈向上蔓延，直到整个程序都被终止
*/


func main() {
	//panic("hello world")
    //运行程序将会引起 panic，输出一个错误消息和 Go 运行时 栈信息，并且返回一个非零的状态码。
	//注意，不像在有些语言中使用异常处理错误，在 Go 中则习惯通过 返回值来标示错误。
	_, err := os.Create("file5")
	if err != nil {
		panic(err)
	}
}
