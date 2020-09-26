package main

import (
	"flag"
	"fmt"
	"os"
)

/*
go和git这种命令行工具，都有很多的子命令。
并且每个工具都有一套自己的 flag，
比如：go build和go get 是go里面的两个不同的子命令。
flag包让我们可以轻松的为工具定义简单的子命令。
 */

func main() {
	//使用NewFlagSet函数声明一个子命令
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	//期望前面定义的子命令作为第一个参数传入
	if len(os.Args) < 2 {
		fmt.Println("expect 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand:", "foo")
		fmt.Println("\tenable:", *fooEnable)
		fmt.Println("\tname:", *fooName)
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand:", "bar")
		fmt.Println("\tenable:", *barLevel)
		fmt.Println("\ttail:", barCmd.Args())
	default:
		fmt.Println("expect 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
