package main

import (
	"fmt"
	"os"
	"strings"
)

/*
环境变量是一个为Unix程序传递配置信息的普遍方式。 让我们来看看如何设置，获取并列举环境变量。
 */

func main() {
	//设置一个自定义的环境变量FOO,值为1
	os.Setenv("FOO", "1")

	//获取自定义的环境变量的值
	fmt.Println("FOO:", os.Getenv("FOO"))

	//需要注意如果该环境变量不存在，则返回空字符串
	fmt.Println("EOO:", os.Getenv("EOO"))

	//我们可以列出当前所有环境变量的键值对
	//通过Envrion函数来获取，它会返回一个键值对形式的字符串切片
	//通过split来获取独立的值
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		fmt.Println(pair[0], "=>", pair[1])
	}
	/*
	....
	NVM_BIN => /Users/liuyang/.nvm/versions/node/v10.14.2/bin
	HOME => /Users/liuyang
	FOO => 1
	*/
}
