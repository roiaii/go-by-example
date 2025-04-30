package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name} // 创建一个person实例，age使用默认值
	p.age = 42
	return &p
}

type rect struct {
	width, height int
}

// 结构体带有方法
func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println(person{"lijun", 27})

	fmt.Println(person{name: "test"})
	fmt.Println(person{name: "test", age: 27})

	fmt.Println(&person{"lisi", 25})

	fmt.Println(newPerson("join"))

	p := newPerson("hi")

	var pp *person
	pp = p
	fmt.Println(pp.age)

	pp.age = 51
	fmt.Println(*pp)

	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())

}
