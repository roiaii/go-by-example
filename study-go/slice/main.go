package main

import "fmt"

func main() {
	// 学习切片基本使用
	a := make([]string, 3)
	fmt.Println("tmp:", a)

	a[0] = "li"
	a[1] = "jun"
	a[2] = "hi"

	fmt.Println("set:", a)

	fmt.Println("tmp size:", len(a))

	fmt.Println("get:", a[2])

	s := append(a, "good")
	fmt.Println("append:", s)
	fmt.Println("a:", a)

	ss := make([]string, len(s))
	copy(ss, s)
	fmt.Println("ss:", ss)

	// slice
	l := s[0:2]
	fmt.Println("l:", l)

	l = s[:1]
	fmt.Println("l:", l)

	l = s[2:]
	fmt.Println("l:", l)

	t := []string{"a", "b", "c"}
	fmt.Println("t:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

	fmt.Println("sss")

}
