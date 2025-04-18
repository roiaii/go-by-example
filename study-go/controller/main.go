package main

import "fmt"

func main() {

	/*
		for 循环与中循环类似，初始值，循环条件，循环后执行逻辑
	*/
	for {
		fmt.Println("sb")
		break
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for n := 1; n < 10; n++ {
		fmt.Println(n / 2)
	}

	fmt.Println("======================")

	// if else语句

	if ii := 10; ii < 0 {
		fmt.Println("ii < 0")
	} else if ii < 10 {
		fmt.Println("ii < 10")
	} else {
		fmt.Println("ii >= 10")
	}

	if i := 100; i < 1000 {
		fmt.Println("我是你爸爸")
	}
}
