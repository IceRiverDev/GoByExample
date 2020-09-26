package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//go写入文件与读取文件内容差不多

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func main() {
	d1 := []byte("hello\nworld\n")
	err := ioutil.WriteFile("helloWorld", d1, 0664)
	check(err)

	f, err := os.Create("tempfile")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("写入%d个字节\n", n2)

	n3, err := f.WriteString("hello god\n")
	fmt.Printf("写入%d个字节\n", n3)
	f.Sync()

	//bufio提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("写入%d个字节\n", n4)
	//使用Flush将缓冲区内容写入文件中
	w.Flush()
	/*
	output:
	写入5个字节
	写入10个字节
	写入9个字节
	 */
}
