package main

import (
	"fmt"
	"sort"
)

/*
有时候我们想使用和集合的自然排序不同的方法对集合进行排序。
例如，我们想按照字母的长度而不是首字母顺序对字符串排序。
这里是一个 Go 自定义排序的例子。
 */

type ByLength []string

func (b ByLength) Len() int {
	return len(b)
}

func (b ByLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func (b ByLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	bs := ByLength{"hellos", "world", "larry", "walls"}
	sort.Sort(bs)
	fmt.Println(bs)
}
