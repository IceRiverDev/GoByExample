package main

import (
	bs64 "encoding/base64"
	"fmt"
)
//这个语法引入了 `encoding/base64` 包并使用名称 `b64` 代替默认的 `base64`。
// 这样可以节省点空间。

func main() {
	data := "abcdef123!$&"
	sEnc := bs64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)//YWJjZGVmMTIzISQm

	//将字符串进行解码，如果不确定输入格式是否正确，那么需要进行错误检查
	sDec, err := bs64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sDec)) //abcdef123!$&
	fmt.Println()

	//获取与url兼容的base64格式进行编解码
	uEnc := bs64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)//YWJjZGVmMTIzISQm

	uDec, err := bs64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec)) //abcdef123!$&

	/*
	标准base64编码和URL兼容base64编码的编码字符串存在 稍许不同（后缀为 + 和 -），但是两者都可以正确解码为 原始字符串。
	 */
}
