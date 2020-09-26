package main

import "fmt"

func main()  {
	//要创建一个空map，需要使用内建的"make":
	//make(map[key-type]val-type).
	m := make(map[string]int)

	//通过键赋值
	m["hello"] = 1
	m["world"] = 2

	fmt.Println(m)

	//通过键取值
	a := m["hello"]
	b := m["world"]
	fmt.Println(a, b)

	//通过len获取map的长度, len返回键值对的数目
	lengs := len(m)
	fmt.Println(lengs)

	//使用delete来删除键值对
	delete(m, "hello")
	fmt.Println(m) //map[world:2]

	//v, ok := m[key]用来判断一个键是否存在

	v, ok := m["world"]
	fmt.Println(v, ok)//2 true

	v2, ok := m["hello"]
	fmt.Println(v2, ok) //0 false

	//字面量创建一个map
	n := map[string]int{"hello":11}
	fmt.Println(n) //map[hello:11]
}
