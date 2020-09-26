package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//可以使用go编写行过滤器示例，将所有的小写文字转为大写文字

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
