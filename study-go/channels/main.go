package main

import "fmt"

func main() {
	// 通道，两个协程之间进行通信，从一个协程建中将值发送到通道，另一个协程丛通道中接受

	// 使用make函数来创建通道，通道类型就是需要传递的值的类型
	message := make(chan string)

	// 在一个新的协程中使用 channel <- 语法将消息发送到通道中
	go func() {
		message <- "hi"
	}()

	// 使用 <- channel语法从通道中接受一个值
	msg := <-message
	fmt.Println(msg)
}
