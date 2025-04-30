package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// 函数命名，多个入参类型一样，可以只声明最后一个参数的类型，前面类型默认一样
func plusPlus(a, b, c int) int {
	return a + b + c
}

// 多值返回函数，返回多个值，通常包括计算结果和异常
func vals() (int, int) {
	return 1, 2
}

// 可变参数函数
// 表示可以传入多个参数，且参数个数是变化的
func sum(nums ...int) int {
	fmt.Println(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 匿名函数 闭包 返回一个函数，隐藏变量
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 递归函数
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/*
*
函数类型的学习：
函数类型：是指定输入参数和返回值的一组规则
*/
// 定义一个函数类型
type MyFunctionType func(int, int) (int, error)

// 使用函数类型声明一个变量
var fn MyFunctionType = func(a, b int) (int, error) {
	return a + b, nil
}

func main() {
	res1 := plus(1, 2)
	fmt.Println("1 + 2 =", res1)

	res2 := plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res2)

	a, b := vals()
	fmt.Println("a:", a, "b:", b)

	_, c := vals()
	fmt.Println("c:", c)

	sum1 := sum(1, 2, 3)
	fmt.Println("sum:", sum1)

	// 通过slice传入可变参数函数
	nums := []int{1, 6, 9, 10}
	sum1 = sum(nums...)
	fmt.Println("sum1:", sum1)

	nextInt := intSeq() // 获取匿名函数
	fmt.Println(nextInt())
	fmt.Println(nextInt()) // 多次调用获取的同一个匿名函数，使用的均是同一个变量i

	newNextInt := intSeq()
	fmt.Println(newNextInt()) // 获取新的匿名函数，并进行调用，以区别上述匿名函数

	res3 := fibonacci(5)
	fmt.Println(res3)

	fmt.Println("========函数类型===")

	fmt.Println(fn(1, 1))

}
