package main

import (
	"fmt"
	"math"
)

/*
接口类型是对动态类型的进一步抽象，
它定义了一组方法签名，只有方法的定义，
没有方法的具体实现
 */

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//要让一个具体类型实现一个接口，就必须实现接口中定义的所有方法
func (r *rect) area() float64 {
	return r.width * r.width
}

func (r *rect) perim() float64 {
	return r.width + r.height
}

func (c *circle) area() float64 {
	return c.radius * 2 * math.Pi
}

func (c *circle) perim() float64 {
	return c.radius * c.radius * math.Pi
}

func printCircleInfo(c *circle) {
	fmt.Printf("圆的面积是:%.2f 周长是:%.2f\n", c.area(), c.perim())
}

func printRectInfo(r *rect) {
	fmt.Printf("长方形的面积是:%.2f 长方形周长是:%.2f\n", r.area(), r.perim())
}

func printInfo(a geometry) {
	switch a.(type) {
	case *circle:
		fmt.Printf("圆的面积是:%.2f 圆周长是:%.2f\n", a.area(), a.perim())
	case *rect:
		fmt.Printf("长方形的面积是:%.2f 长方形周长是:%.2f\n", a.area(), a.perim())
	}
}

func main() {
	r := &rect{
		3.2,
		4.5,
	}

	c := &circle{
		2.2,
	}

	//如果需要打印不同几何体的信息，我们需要为不同的几何体定义相同的函数
	//有一千个几何体就需要定义一千个方法，这样就造成了代码的重复
	printCircleInfo(c)
	printRectInfo(r)

	fmt.Println()

	//当有了接口以后，我们可以将参数类型定义为接口类型，
	//这样只要是实现了这个接口的具体类型都可以调用这个函数
	//减少了不必要的重复代码，提供了一个统一的接口来打印各种几何体的信息，这就是接口强大的地方
	printInfo(c)
	printInfo(r)
}
