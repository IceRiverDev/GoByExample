package main

import "fmt"

/*
除了可以实例化一个类型，我们也可以该类对象定义实例方法
 */

type Person2 struct {
	Name string
	Age int
}

//定义值类型接收者方法
func (p Person2) SayName() string {
	return p.Name
}

func (p Person2) SayAge() int {
	return p.Age
}

//定义指针类型接收者方法
func (p *Person2) SpeakName() string {
	return p.Name
}

func (p *Person2) SpeakAge() int {
	return p.Age
}

func main() {
	p := Person2{
		"larry wall",
		70,
	}

	fmt.Println(p.SayAge())
	fmt.Println(p.SayName())

	fmt.Println(p.SpeakAge())
	fmt.Println(p.SpeakName())

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，
	// 或者让方法能够改变接受的结构体。

	var p2 *Person2 = &Person2{
		"matz",
		60,
	}
	fmt.Println(p2.SayName())
	fmt.Println(p2.SayAge())

	var p3 Person2 = Person2{
		"Tom",
		60,
	}

	fmt.Println(p3.SpeakName())
	fmt.Println(p3.SpeakAge())
}
