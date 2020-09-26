package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//大部分的文件读取任务是将文件内容读取到内存中
	data, err := ioutil.ReadFile("helloWorld")
	checkError(err)
	fmt.Println(string(data))// {"apple":5,"lettuce":7}

	// 你经常会想对于一个文件是怎么读并且读取到哪一部分进行更多的控制。
	// 对于这个任务，从使用 `os.Open` 打开一个文件获取一个 `os.File` 值开始。
	f, err := os.Open("helloWorld")

	//一般我们会把切片长度和容量设置为1024或者1024的整数倍
	data2 := make([]byte, 5)
	num, err := f.Read(data2)
	checkError(err)
	fmt.Printf("读取的字节数量%d, 字符串内容%s\n", num, string(data2)) //读取的字节数量5, 字符串内容{"app

	o2, err := f.Seek(6, 0)
	checkError(err)
	data3 := make([]byte, 2)
	num2, err := f.Read(data3)
	fmt.Printf("%d bytes @ %d: %s\n", num2, o2, string(data3))

	o3, err := f.Seek(6, 0)
	checkError(err)
	data4 := make([]byte, 2)
	//这里读取的时候使用了缓冲，因此读取性能更好
	num3, err := io.ReadAtLeast(f, data4, 2)
	checkError(err)
	fmt.Printf("%d bytes @ %d: %s\n", num3, o3, string(data4))

	// 没有内置的回转支持，但是使用 `Seek(0, 0)` 实现。
	_, err = f.Seek(0, 0)
	checkError(err)

	// bufio包实现了带缓冲的读取，这不仅对于很多小的读
	// 取操作能够提升性能，也提供了很多附加的读取函数。
	r := bufio.NewReader(f)
	b4, err := r.Peek(5)
	checkError(err)
	fmt.Println("5 bytes: %s\n", string(b4))

	f.Close() //释放文件资源
	/*
	{"apple":5,"lettuce":7}
	读取的字节数量5, 字符串内容{"app
	2 bytes @ 6: e"
	2 bytes @ 6: e"
	5 bytes: %s
	 {"app
	 */
}
