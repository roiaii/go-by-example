package main

import "fmt"

func main() {
	// map基本操作
	a := make(map[string]string) // 使用内置函数make创建map类型数据
	fmt.Println("a:", a)

	a["key1"] = "li"
	a["key2"] = "jun"

	fmt.Println("a:", a)

	fmt.Println("a size:", len(a)) // 输出map长度

	s := a["key2"] // 获取map的值通过key
	fmt.Println("key2:", s)

	delete(a, "key2") // 删除map中键值对
	fmt.Println("a:", a)

	_, prs := a["key2"]
	fmt.Println("prs:", prs) // 如果该key存在则prs为true，并且将取到的值放在前面的变量。用于区别要取的值出现歧义，比如key不存在和空串的情况。由于不需要该值，使用空白标识符接收该值

	aa, prs3 := a["key1"] // 需要该值，放入aa中，并且判断该值是否存在放在prs3
	fmt.Println("aa:", aa, ", prs3:", prs3)

	b := map[string]int{"key1": 1, "key2": 2} // 通过声明并初始化的方式创建map类型数据

	fmt.Println("b:", b)

}
