package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	//提供一个url，它包含了schema，认证用户名和密码，主机名，端口号和查询参数
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	//解析这个url并且确保没有错误返回
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme) // postgres

	fmt.Println(u.Hostname()) //host.com

	fmt.Println(u.Host) //host.com:5432

	fmt.Println(u.User) //user:pass

	fmt.Println(u.User.Username()) //user

	fmt.Println(u.User.Password()) // pass

	//提取端口
	h := strings.Split(u.Host,  ":")
	fmt.Println(h[1])//5432

	//提取路径和片段信息
	fmt.Println(u.Path)//path
	fmt.Println(u.Fragment) //f

	//提取查询字符串
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	//已解析的查询参数 map 以查询字符串为键，对应
	// 值字符串切片为值，所以如果只想得到一个键对应的第
	// 一个值
	fmt.Println(m) //map[k:[v]]
	fmt.Println(m["k"][0]) // v

}
