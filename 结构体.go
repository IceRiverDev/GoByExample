package main

import "fmt"

/*
由于go不是一个真正的面向对象的语言，严格来说，
它是面向接口的语言。我们可以用结构体来模拟面向对象的
编程范式。
 */

//定义一个类
type Person struct {
	Name string
	Age int
}

func main() {
	//实例化一个类
	p := Person{
		"larry wall",
		70,
	}

	//在初始化一个结构体时，我们可以指定字段，也可以不指定
	p2 := Person{
		Name: "Matz",
		Age:  60,
	}
	fmt.Println(p2)

	//通过.来访问结构体字段
	name := p2.Name
	age := p2.Age
	fmt.Printf("我叫%s, 今年%d岁了。。\n", name, age) //我叫Matz, 今年60岁了。。
	fmt.Println(p) //只打印字段值，不打印字段。{larry wall 70}
	fmt.Printf("%+v", p) //连带字段一起打印，{Name:larry wall Age:70}

	//我们可以获取一个结构体的地址，通过地址来访问结构体的字段值
	var p3 *Person = &Person{
		Name: "Tom",
		Age:  20,
	}

	//取值有两种方式，可以通过常规方式获取，也可以通过简写方式，效果一样
	name2 := p3.Name
	name3 := (*p3).Name

	age2 := p3.Age
	age3 := (*p3).Age

	fmt.Println(name2, name3)
	fmt.Println(age2, age3)

}
