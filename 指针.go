package main

import "fmt"

/*
指针是存储另一个变量的内存地址的变量。
我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。
 */

func main() {
	var b int = 10
	var a *int = &b
	/*由于a变量保存的是变量b的内存地址，因此访问a变量的内存地址中的值就
	相当于访问b变量内存地址当中的值
	*/

	//地址取值
	c := *a //将取出的值赋值给变量c
	fmt.Println(c) //10

	//地址赋值
	//除了从地址当中取值，我们也可以给地址赋值
	//一旦地址中的值变为11，那么变量b的值也会变为11
	//因为变量a和变量b指向同一个内存地址
	*a = 11
	fmt.Printf("a的值:%d b的值:%d\n", *a, b)

	fmt.Printf("%p %p\n", a, &b) //0xc000016088 0xc000016088

	/*
	Go空指针当一个指针被定义后没有分配到任何变量时，它的值为nil。 nil指针也称为空指针。 nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。 一个指针变量通常缩写为 ptr。
	空指针判断：
	 */
	var ptr map[string]int
	var ptr2 = make(map[string]int)
	fmt.Println(ptr2 != nil) //true ptr2 不是空指针
	fmt.Println(ptr == nil) //true ptr 是空指针
}
