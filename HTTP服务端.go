package main

import (
	"fmt"
	"net/http"
)

//定义一个处理针对hello的处理器
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

//定义一个处理headers响应的处理器
func Headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	//当请求hello资源时，会调用hello处理器函数来执行响应操作
	http.HandleFunc("/hello", Hello)

	//当请求hello资源时，会调用headers处理器函数来执行响应操作
	http.HandleFunc("/headers", Headers)

	http.ListenAndServe(":8080", nil)
}


