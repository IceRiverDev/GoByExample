package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
	    // 通过errors.new来构造一个使用给定错误消息的error对象
		//这个error对象实现了内建error接口
		return -1, errors.New("can't work with 42")
	}
	// 返回错误值为nil代表没有错误
	return arg + 3, nil
}

//除了使用errors包来创建error对象，我们也可以自己定义一个error类型
//并且让它实现内建error接口类型
type argErrors struct {
	arg int
	problem string
}

func (arg *argErrors) Error() string {
	return fmt.Sprintf("%d - %s", arg.arg, arg.problem)
}

func f2(arg int) (int, error) {
	//使用我们自定义的错误类型
	if arg == 42 {
		return -1, &argErrors{
			arg:     arg,
			problem: "arguments wrong, can't work with arg",
		}
	}
	return arg + 3, nil
}

func main() {
	for _, v := range []int{7, 42} {
		if r, err := f1(v); err != nil {
			fmt.Println("f2 failed:", err)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	//通过类型断言来获取实际类型的值
	_, err := f2(42)
	if t, ok := err.(*argErrors); ok {
		fmt.Println("the argument is", t.arg)
		fmt.Println("the problem is", t.problem)
	}
}
