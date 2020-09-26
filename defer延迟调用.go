package main

import (
	"fmt"
	"io"
	"os"
)

//defer用来确保在一个函数退出前执行额外的操作。
//例如当我们想要创建一个文件，然后对它进行写操作，
//然后在结束时来关闭它

func main() {
	f := createFile("helloWorld")
	defer f.Close()

	writeFile(f)
	readFile(f)

}

func createFile(f string) *os.File {
	fmt.Println("creating")
	file, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	i, err := f.Write([]byte("hello worlds"))
	if err != nil {
		panic(err)
	}
	fmt.Println("写入：", i)
}

func readFile(f *os.File) {
	str := make([]byte, 30)
	f.Seek(0, io.SeekStart)
	i, err := f.Read(str)
	if err == io.EOF {
		fmt.Println("读取完毕")
	} else if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(str[:i]))
}
