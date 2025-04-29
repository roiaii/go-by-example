package main

import "fmt"

// 值传递
func zeroVal(val int) {
	val = 0
}

// 址传递
func zeroPtr(val *int) {
	*val = 0
}

func main() {
	i := 1
	fmt.Println(i)

	zeroVal(i)
	fmt.Println(i)

	zeroPtr(&i)
	fmt.Println(i)

	fmt.Println(&i)
}
