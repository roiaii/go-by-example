package main

import "fmt"

/*
在学习的时候要注意对比学习，与java对比学习，加深印象
1. 基本数据类型
整型
浮点型
布尔型
字符串

以及基本数据类型的基本使用
*/
func main() {
	fmt.Println("1 + 2 = ", 1+2) // 1 + 2 = 3

	fmt.Println("sb" + "lijun")

	fmt.Println("1.0 / 2 = ", 1.0/2) // 0.5

	fmt.Println("false && true = ", false && true) // false

	// 声明变量
	// var可以声明一个或者多个变量。可以根据初始值的变量类型推断出变量的数据数据类型
	var a = 1
	fmt.Println(a)

	var str = "sb lijun"
	fmt.Println(str)

	var b, c, d = 1, 2, "sb xiaoming"
	fmt.Println(b, c, d)

	// 声明后没有给初始值，会给默认初始值，比如int为0
	var f int
	fmt.Println(f) // 0

	// var := 1 是声明变量类型并赋初始值的简写
	s := "sbsbsb"
	fmt.Println(s)

	m := 66
	fmt.Println(m)

}
