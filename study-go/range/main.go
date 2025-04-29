package main

import "fmt"

func main() {
	// range基本使用，在各种数据结构中的使用

	// 对切片slice的使用
	a := []int{1, 2, 3}
	sum := 0

	for _, v := range a {
		sum += v
	}
	fmt.Println("sum:", sum)

	for i, v := range a {
		fmt.Println("index:", i, "value:", v)
	}

	// 遍历map类型
	b := map[string]int{"key1": 1, "key2": 2}
	for k, v := range b {
		fmt.Println("%s -> %d", k, v)
	}

	for k := range b {
		fmt.Println("key:", k)
	}

}
