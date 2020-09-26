package main

import "fmt"
import s "strings" //给导入的包去一个别名

//标准库的strings包提供了很多有用的字符串相关的函数。
//这里是一些用来让你对这个包有个初步了解的例子。
//你可以在strings包文档中找到更多的函数

var p = fmt.Println

func main() {

	p("Contains: ", s.Contains("test", "e"))

	p("Count: ", s.Count("test", "t"))

	p("HasPrefix: ", s.HasPrefix("test", "t"))

	p("Index: ", s.Index("test", "e"))

	p("Join: ", s.Join([]string{"a", "b", "c"}, "-"))

	//字符串也可以像切片一样通过索引来获取子串
	//字符串本质上来说就是字符切片，因此切片的操作
	//对字符串来说也适用

	p("len: ", len("test"))
	p("Char: ", "hello"[0])

	/*
	输出结果：
	Contains:  true
	Count:  2
	HasPrefix:  true
	Index:  1
	Join:  a-b-c
	len:  4
	Char:  104
	 */

}
