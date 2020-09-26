package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//有时，我们的Go程序需要生成其他的，非Go进程
//例如调用外部shell命令

func main() {

	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(">date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello world!\nwow cool!"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)
	if err != nil {
		os.Exit(2)
	}
	fmt.Println(string(grepBytes))
	grepCmd.Wait()

	//在生成命令时，我们需要提供显示描述的命令和参数
	//数组，而不能只传递一个命令行字符串，如果你想使用一个字符串生成一个完整的命令，那么你可以使用
	//bash命令的-c选项
	lsCmd := exec.Command("/bin/zsh", "-c", "ls -l")
	lsOut, err := lsCmd.Output()
	fmt.Println("ls output:")
	fmt.Println(string(lsOut))


	//sleepCmd := exec.Command("sleep", "3")
	//if err := sleepCmd.Start(); err != nil {
	//	panic(err)
	//}
	//fmt.Println("waiting....")
	//sleepCmd.Wait()

	sleepCmd2 := exec.Command("sleep", "10")
	sleepCmd2.Output()
	fmt.Println("waiting....")
}
