package main

import (
	"fmt"
	"math"
)

// 定义一个几何体接口，有求面积方法和求周长接口方法
type geometry interface {
	area() float64
	perimeter() float64
}

// 定义长方形类，并实现接口
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// 定义圆类，并实现接口
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
