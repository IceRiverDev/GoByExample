package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
单引号:
要声明字节或符文，我们使用单引号。在声明字节时，我们必须指定类型。如果我们未指定类型，则默认类型为符文。
单引号只能包含一个字符。在用单引号声明一个带有两个字符的字节或符文时，编译器将引发错误。
invalid character literal (more than one character)

双引号:
它用于定义字符串。在双引号中定义的字符串将使用转义字符。
例如，当打印具有\n的字符串时，将打印新行。同样，\t将打印制表符。

反引号:
它也用于定义字符串。用反引号编码的字符串是原始文字字符串，不支持任何转义。
 */

//让我们根据上面的概念来写一些示例
//需要注意的是反引号中的特殊转义字符不会按照特殊含义被解释

//示例1

func main()  {
	//r := 'ab' //invalid character literal (more than one character)

	//字符串在双引号中
	x := "tit\nfor\ttat"
	fmt.Println("Printing String in Double Quotes:")
	fmt.Printf("x is: %s\n", x)

	//字符串在反引号中
	y := `tit\nfor\ttat`
	fmt.Println("\nPriting String in Back Quotes:")
	fmt.Printf("y is: %s\n", y) //按照字面量输出，不会被解释

	//用单引号声明一个字节
	var b byte = 'a'
	fmt.Println("\nPrinting Byte:")
	fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(b), reflect.TypeOf(b), b)

	//用单引号声明一个rune
	r := '£'
	fmt.Printf("Size: %d\nType: %s\nUnicode CodePoint: %U\nCharacter: %c\n", unsafe.Sizeof(r), reflect.TypeOf(r), r, r)

	//int32和rune是别名的关系
	var a11 int32 = 'a'
	var b11 rune = 'a'
	fmt.Println(a11 == b11) //true

	//r = 'ab'这个表达式会引发异常，因为单引号只能声明单个字符

	//反引号还有一个作用是类似于heredoc或者js里面的模板字符串那样，可以换行写入多行字符串
	//字符串前面的空格会被正常显示
	var str = `
hello world
hello god
hello girl
hello boy`

	fmt.Println(str)
	/*
	输出：
	hello world
	hello god
	hello girl
	hello boy
	 */

	var str2 = `
    hello world
    hello god
    hello girl
    hello boy`
	fmt.Println(str2)
	/*
	输出：
	   hello world
	   hello god
	   hello girl
	   hello boy
	 */
}
