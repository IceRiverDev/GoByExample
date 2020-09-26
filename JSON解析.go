package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//将布尔值编码成json的布尔值
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) //true

	//将整数编码成json的整数
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 1

	//将浮点数编码成json的浮点数
	fltB, _ := json.Marshal(1.2)
	fmt.Println(string(fltB))//1.2

	//将字符串编码成json的字符串
	strB, _ := json.Marshal("hello")
	fmt.Println(string(strB)) // "hello"

	slcD := []string{"apple", "peach", "pear"}
	slcB, _:= json.Marshal(slcD)
	fmt.Println(string(slcB)) //["apple","peach","pear"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) //{"apple":5,"lettuce":7}

	res1D := &Response1{
		1,
		[]string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		1,
		[]string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//也可以将json格式的数据解码为go值
	byt := []byte(`{"num": 2.1, "strs":["a", "b"]}`)
	var dat map[string]interface{}

	//将JSON数据转为map格式的数据
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(2.1)
	n := dat["num"].(float64)
	s := dat["strs"].([]interface{})
	fmt.Printf("num: %T %v, strs: %T %v\n", n, n, s, s)

	//我们也可以解码 JSON 值到自定义类型。这个功能的好处就
	//是可以为我们的程序带来额外的类型安全加强，并且消除在
	//访问数据时的类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), res)
	fmt.Printf("%+v\n",res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	f, err := os.OpenFile("helloWorld.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f.Seek(0, io.SeekStart)
	enc2 := json.NewEncoder(f)
	d := map[string]int{"apple":5, "lettuce": 7}
	enc.Encode(d)
	enc2.Encode(d)
}
