package main

import "fmt"

func main()  {

	nums := []int{1, 2, 3, 4, 5}
	sum := 0

	/*
	range遍历切片
	用range遍历切片返回索引和与其对应的值
	可以使用空白标识符忽略索引
	 */
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum) //15

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:",i)
		}
	}

	/*
	range遍历map
	 */
	kvs := map[string]string{
		"a":"apple",
		"b":"banana",
	}

	//遍历键值
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) //a -> apple
		                                      //b -> banana
	}

	//也可以只遍历键
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` 在字符串中迭代unicode码点(code point)。
	// 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
	/*
	输出：
	0 103
	1 111
	2 108
	3 97
	4 110
	5 103
	*/

}
