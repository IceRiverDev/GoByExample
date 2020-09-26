package main

import (
	"fmt"
	"strings"
)

//若切片中存在某元素，则返回该元素的索引值
func Index(strs []string, s string) int {
	for i, v := range strs {
		if v == s {
			return i
		}
	}
	return -1
}

//判断切片是否包含某元素
func Include(strs []string, s string) bool {
	return Index(strs, s) >= 0
}

//只要存在满足条件的切片元素就返回true，否则返回false
func Any(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if f(v) {
			return true
		}
	}
	return false
}

//判断切片中所有元素是否都满足给定的条件
func All(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if !f(v) {
			return false
		}
	}
	return true
}

//将满足条件的切片元素过滤出来，返回一个新的切片
func Filter(strs []string, f func(string) bool) []string {
	filteredResults := make([]string, 0)

	for _, v := range strs {
		if f(v) {
			filteredResults = append(filteredResults, v)
		}
	}
	return filteredResults
}

//对原始的切片元素根据给定条件进行转换，并返回原始切片
func Map(strs []string, f func(string) string) []string {
	for i, v := range strs {
		strs[i] = f(v)
	}
	return strs
}

func main()  {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "apple"))

	fmt.Println(Include(strs, "apple"))

	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "a")
	}))

	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(Map(strs, func(s string) string {
		return s + "s"
	}))

	fmt.Println(Filter(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(Map(strs, func(s string) string {
		return strings.ToUpper(s)
	}))
}
