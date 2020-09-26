package main

import (
	"fmt"
	"sort"
)

/*
Go的sort包实现了内置和用户自定义数据类型的排序 功能。我们首先关注内置数据类型的排序。
 */
func main()  {
	strs := []string{
		"c",
		"a",
		"b",
	}

	//Note:注意排序是原地更新的，所以它会改变给定的序列并且不返回一个新值。

	//字符串排序
	sort.Strings(strs)
	fmt.Println(strs) //a b c

	//数值排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println(ints) // 2 4 7

	//我们可以检查一个序列是否已经被排过序
	s1 := sort.IntsAreSorted(ints)
	s2 := sort.StringsAreSorted(strs)
	fmt.Println("s1 sorted:", s1) // true
	fmt.Println("s2 sorted:", s2) // true
}
