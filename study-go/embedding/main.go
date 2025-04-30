package main

import (
	"fmt"
)

// 学习组合类型，在struct中嵌套struct

type base struct {
	num int
}

func (b base) describle() string {
	return fmt.Sprintf("base.describle:%v", b.num)
}

type container struct {
	base // 嵌套struct
	str  string
}

func main() {
	co := container{base: base{num: 0}, str: "hi"}
	fmt.Println(co.num)      // 直接调用嵌套的struct的成员变量
	fmt.Println(co.base.num) // 完整的调用链路

	fmt.Println(co.describle()) // 调用嵌套的struct的方法

	type describler interface {
		describle() string
	}

	var d describler = co // 使用嵌套了结构体且有对应方法的结构体来赋值给接口，实现到接口上
	fmt.Println(d.describle())
	d = base{num: 1}
	fmt.Println(d.describle())
}
