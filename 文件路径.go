package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	//拼接目录
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)

	//使用join拼接目录而不是手动去拼接
	//join会自动将多余的分隔符和目录去除，使得路径更加规范
	fmt.Println(filepath.Join("dir1//", "file2name"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	//Dir和Base用来分隔路径中的目录和文件

	//提取目录
	fmt.Println(filepath.Dir("dir/file"))

	//提取文件
	fmt.Println(filepath.Base("dir/file"))

	//判断路径是否为绝对路径
	fmt.Println(filepath.IsAbs("dir/file"))

	//将扩展名提取出来
	ext := filepath.Ext("file.txt")
	fmt.Println("the extension is ", ext)

	//如果想要获取剔除后缀名的文件名，可以使用TrimSuffix
	fmt.Println(strings.TrimSuffix("file1.txt", ext))

	//使用Rel寻找basepath与targetpath之间的相对路径。
	//如果相对路径不存在，则返回错误
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	/*
	dir1/dir2/filename
	dir1/file2name
	dir1/filename
	dir
	file
	false
	the extension is  .txt
	file1
	t/file
	../c/t/file
	 */

}
