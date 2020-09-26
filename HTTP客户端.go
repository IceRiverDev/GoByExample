package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://baidu.com")
	if err != nil {
		panic(err)
	}

	//Header字段是一个map[string][]string类型的数据
	//我们对它进行遍历
	for k, v := range resp.Header {
		for _, elem := range v {
			fmt.Println(k, elem)
		}
	}

	defer resp.Body.Close()

	//打印响应状态
	fmt.Println(resp.Status)

	//打印响应状态码
	fmt.Println(resp.StatusCode)

	//打印前五行内容
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	/*
	Connection Keep-Alive
	Content-Type text/html
	Last-Modified Tue, 12 Jan 2010 13:48:00 GMT
	Etag "51-47cf7e6ee8400"
	Content-Length 81
	Cache-Control max-age=86400
	Expires Sat, 26 Sep 2020 13:33:34 GMT
	Date Fri, 25 Sep 2020 13:33:34 GMT
	Server Apache
	Accept-Ranges bytes
	200 OK
	200
	<html>
	<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
	</html>
	*/
}
