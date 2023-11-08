package interfaces

import (
	"fmt"
	"math"
)

type ShapeInterface interface {
	Area() float64
}

type ShapeStruct struct {
	s ShapeInterface
}

type ShapeType ShapeInterface

func (s ShapeStruct) Area() float64 {
	return s.s.Area()
}

type Circle struct {
	r float64
}

type Rectangle struct {
	h, w float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) Area() float64 {
	return r.h * r.w
}

func test() {
	c := Circle{5}
	csi := ShapeInterface(c) // note round ()
	css := ShapeStruct{c}    // note curly {}
	cst := ShapeType(c)      // note round ()

	fmt.Println("Area of circle is ", c.Area())
	fmt.Println("Area of shape type is ", cst.Area())
	fmt.Println("Area of shape struct is ", css.Area())
	fmt.Println("Area of shape interface is ", csi.Area())

	r := Rectangle{5, 4}
	rsi := ShapeInterface(&r) // note pointer - due to Area() defined on pointer receiver
	rss := ShapeStruct{&r}
	rst := ShapeType(&r)

	fmt.Println("Area of rectangle is ", r.Area())
	fmt.Println("Area of shape type is ", rst.Area())
	fmt.Println("Area of shape struct is ", rss.Area())
	fmt.Println("Area of shape interface is ", rsi.Area())
}
