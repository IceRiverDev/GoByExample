package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//在前面的例子中，我们了解了生成外部进程的知识，
//当我们需要访问外部进程时需要这样做，但是有时候，我们只想用其他的（也许是非 Go 程序）来完全替代当前的 Go 进程。
//这时候，我们 可以使用经典的 exec 方法的Go实现。

func main() {

	//获取命令ls的绝对路径
	binary, lookerr := exec.LookPath("ls")
	if lookerr != nil {
		panic(lookerr)
	}

	//构建参数切片
	args := []string{"ls", "-a", "-l", "-h"}

	//获取系统的环境变量
	env := os.Environ()

	//传入三个参数，第一个参数是该命令的搜索路径
	//第二个参数是该命令以及参数的切片
	//第三个参数是环境变量切片
	//如果Exec调用成功，那么我们当前的go进程会被替换为ls进程
	//当ls进程执行完毕以后，那么程序就会直接退出，而不会把
	//控制权交还给go进程，因此最后的打印不会被执行
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	//不会被执行
	fmt.Println("hello world")
	/*
	....
	drwxr-xr-x  78 liuyang  staff   2.4K Sep 26 09:11 .
	drwxr-xr-x  16 liuyang  staff   512B Sep 25 08:31 ..
	drwxr-xr-x   9 liuyang  staff   288B Sep 26 09:06 .idea
	-rw-r--r--   1 liuyang  staff   972B Sep 24 07:45 Base64编码.go
	 */
}
