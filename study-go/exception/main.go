package main

import (
	"errors"
	"fmt"
)

// 使用原有的error接口
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42.")
	}
	return arg + 2, nil
}

// 实现Error方法，实现自定义异常
type argError struct {
	arg  int
	prob string
}

// 实现Error方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with 42."}
	}
	return arg + 2, nil
}

func main() {
	// go语言习惯使用一个独立明确的返回值来传递错误，这能清楚的知道是哪个函数返回了错误信息
	// 按照惯例，错误信息放在最后一个参数，且是error类型

	s := []int{42, 5, 6}
	for _, v := range s {
		if r, e := f1(v); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, v := range s {
		if r, e := f2(v); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

}
