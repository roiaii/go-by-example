package main

import "fmt"

func main() {
	// 使用数据
	var a [5]int
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	b[4] = 100
	fmt.Println(b)

	var s [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			s[i][j] = i + j
		}
	}
	fmt.Println(s)
}
