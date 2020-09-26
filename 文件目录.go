package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//defer os.RemoveAll("kobe") //最终清理掉临时目录

	//在当前目录下创建一个子目录
	//err := os.Mkdir("Kobe", 0755)
	//checkError(err)

	createEmptyFile := func(name string) {
		d := []byte("")
		checkError(ioutil.WriteFile(name, d, 0644))
	}
	//createEmptyFile("kobe/file1")

	//创建一个有层级的目录
	//类似于mkdir -p
	err := os.MkdirAll("subdir/parent/child", 0755)
	checkError(err)

	createEmptyFile("subdir/parent/child/file1")
	createEmptyFile("subdir/parent/child/file2")
	createEmptyFile("subdir/parent/child/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := ioutil.ReadDir("subdir/parent")
	checkError(err)

	fmt.Println("Listing subdir/parent")
	for _, file := range c {
		fmt.Println(" ", file.Name(), file.IsDir())
	}

	//更改当前目录可以使用Chdir
	err = os.Chdir("subdir/parent/child")
	checkError(err)

	c , err = ioutil.ReadDir(".")
	checkError(err)
	for _, file := range c {
		fmt.Println(" ", file.Name(), file.IsDir())
	}

	//回到之前的目录
	err = os.Chdir("../../../")
	checkError(err)

	fmt.Println("virsting subdir")
	err = filepath.Walk("subdir", visit)
	checkError(err)
	/*
	Listing subdir/parent
	  child true
	  file1 false
	  file2 false
	  file3 false
	  file4 false
	virsting subdir
	  subdir true
	  subdir/parent true
	  subdir/parent/child true
	  subdir/parent/child/file1 false
	  subdir/parent/child/file2 false
	  subdir/parent/child/file3 false
	  subdir/parent/child/file4 false
	 */
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
