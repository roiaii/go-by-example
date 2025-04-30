package main

import (
	"fmt"
	"time"
)

func f(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// 学习协程

	// 在主线程中同步调用一个函数
	f("hi")

	// 在新的协程中并发执行函数
	go f("goroutine")

	// 并发调用匿名函数
	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg, ":", i)
			time.Sleep(time.Millisecond * 100)
		}
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}
