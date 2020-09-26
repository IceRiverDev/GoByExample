package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLNAME xml.Name `xml:"plant"`
	Id int `xml:"id,attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

type Nesting struct {
	XMLName xml.Name `xml:"nesting"`
	Plants []*Plant `xml:"parent>child>plant"`
}


func (p Plant) String() string {
	return fmt.Sprintf("plant id=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{
		Id: 27,
		Name: "Coffee",
	}
	coffee.Origin = []string{
		"Ethiopia", "Brazil",
	}

	//使用marshalIndent生成可读性更好的输出结果
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))

	//给输出结果添加一个通用的XML头部信息
	fmt.Println(xml.Header + string(out))

	//将xml数据重新解析成结构体形式的数据
	//如果xml格式不正确，或者无法映射到Struct，将会返回一个描述性错误
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Printf("%#v", p)

	lw := &Plant{
		Id: 81,
		Name: "larry",
	}

	lw.Origin = []string{"USA", "UK"}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, lw}

	output, _ := xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(output))

}
