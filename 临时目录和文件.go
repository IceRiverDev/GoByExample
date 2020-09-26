package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//在程序运行时，我们经常会创建一些运行时用到的，程序结束以后就不再使用的数据。
//临时目录和文件对于上面的问题很有用，因为它不会随着时间推移而污染文件系统

func chekSomeErrors(e error)  {
	if e != nil {
		panic(e)
	}
}

func main() {
	//如果传入的第一个参数是""，那么会在系统默认位置创建临时文件
	//第二个参数作为前缀
	//剩余的部分自动生成，这样在并发调用时，生成不重复的文件名
	//在类unix操作系统下，临时目录一般是/tmp.
	f, err := ioutil.TempFile("", "sample")
	chekSomeErrors(err)
	fmt.Println("Temp file name : ", f.Name())

	//程序退出时清理临时文件
	//虽然操作系统会在某个时间清理临时文件
	//但是手动清理是好习惯
	defer os.Remove(f.Name())

	//写入一些数据到临时文件
	_, err = f.Write([]byte("hello world"))
	chekSomeErrors(err)

	//如果需要写入多个临时文件，最好是为其创建一个临时目录
	//TempDir返回一个目录名而不是打开的文件
	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("Temp dir name:", dname)
	defer os.RemoveAll(dname)

	//将临时目录与临时文件合成一个完整的临时文件路径，并写入数据
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte("hello world"), 0664)
	chekSomeErrors(err)

	b, err := ioutil.ReadFile(fname)
	fmt.Println(string(b))
}
