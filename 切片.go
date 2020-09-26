package main

import "fmt"

// slice切片是引用类型
// 与数组不同，slice 的类型仅由它所包含的元素决定（不需要元素的个数）。
// 要创建一个长度非零的空 slice，需要使用内建的方法 `make`。
// 这里我们创建了一个长度为3的 `string` 类型 slice（初始化为零值）。

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	//和数组一样的设值方式
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("sets:", s)   // [a b c]
	fmt.Println("get:", s[2]) // c

	//len返回切片长度
	fmt.Println(len(s)) // 3

	// cap返回切片容量
	fmt.Println(cap(s)) // 3

	// append用来给切片添加新的值，并且返回一个新的切片
	// 因此我们需要接受它的返回值
	//创建一个长度和容量为3，类型为int的切片
	i := make([]int, 3, 3) //效果与make([]int, 3)一样
	i = append(i, 1, 2, 3)

	// 注意这里由于一开始创建一个空切片，元素默认初始化为int类型的零值
	// 所以为0，然后我们在切片尾部追加了四个元素，所以0值元素依然存在
	// 注意这里由于是在原来基础上追加元素，而切片最大长度为3，因此切片会自动扩容，并且返回一个新的切片
	fmt.Println("len:",len(i), "cap:",cap(i)) // 此时长度为6, 容量为6
	fmt.Println(i) // [0 0 0 1 2 3]

	// 如果想要把前面的零值剔除掉，那么需要对切片以索引的方式来获取新的切片
	i2 := i[3:6]
	fmt.Println(i2) // [1, 2, 3]

	//Slice支持通过slice[low:high] 语法进行切片操作。
	i3 := []int{1, 2, 3, 4, 5, 6, 7}//切片也支持字面量创建法
	fmt.Println(i3) // 长度容量都为7

	l := i3[0:] //返回整个切片
	fmt.Println(l) //[1 2 3 4 5 6 7]

	l = i3[2:5] // 需要注意右半侧的值无法取到，因此实际取的是2 3 4索引位置上的值
	fmt.Println(l) // [3 4 5]

	l = i3[:] // 返回整个切片
	fmt.Println(l) //[1 2 3 4 5 6 7]

	// copy可以用来拷贝一个切片
	s1 := []int{1, 2, 3}
	s2 := make([]int, 2)
	copy(s2, s1)
	fmt.Println(s2) //[1 2] 由于这里s2的长度为2，所以多出来的元素会被丢弃

	s3 := []int{1, 2}
	s4 := make([]int, 4)
	copy(s4, s3)
	fmt.Println(s4) // [1 2 0 0] 由于这里s4的长度大于s3的元素个数，因此空出来的位置会用初始值0填充

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice可以组成多维数据结构。
	// 内部的slice长度可以不一致，这和多维数组不同。
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]

	/*
	注意点1：
	对于切片的数据类型来说，它本身是一个引用类型，因此当它赋值给一个变量时，相当于是把它的地址赋值给了这个变量
	 */
	a1 := []int{1, 2, 3}
	fmt.Printf("%p\n", a1) //0xc000018100

	a2 := a1
	fmt.Printf("%p\n", a2) // 0xc000018100

	/*
	注意点2：
	当把变量a1赋值给a2并不是值传递，而是地址传递或者引用传递，两个变量指向同一个值的内存地址
	因此，如果改变变量a1中切片的值，那么a2也会发生改变
	 */

	a1[0] = 11
	fmt.Println(a1) //[11 2 3]
	fmt.Println(a2) //[11 2 3]

	/*
	注意点3：
	切片的地址与切片的第一个元素的地址是一样的
	 */

}
