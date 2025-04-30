package main

import "fmt"

// 传入comparable类型的K，任意类型V的map，并返回map中K的切片
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// 使用泛型实现任意类型的链表
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	val  T
	next *element[T]
}

func (l *List[T]) putValue(val T) {
	if l.tail == nil {
		l.head = &element[T]{val: val}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{val: val}
		l.tail = l.tail.next
	}
}

func (l *List[T]) getAllValues() []T {
	var r []T
	for e := l.head; e != nil; e = e.next {
		r = append(r, e.val)
	}
	return r
}

func main() {
	// 学习泛型
	// 类型参数
	m := map[string]int{"lijun": 1, "bilbo": 2}
	MapKeys(m) // 当调用泛型函数时，不需要为K和V指定类型，编译器会自动进行类型推断

	_ = MapKeys[string, int](m) // 显式的为泛型函数指定类型

	// 测试链表
	l := List[int]{}
	l.putValue(1)
	l.putValue(2)
	l.putValue(5)

	fmt.Println(l.getAllValues())

}
