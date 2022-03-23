package main

import "fmt"

type figure2D interface {
	area() float64
}
type square struct {
	base float64
}
type rectangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.base * s.base
}
func (r rectangle) area() float64 {
	return r.base * r.height
}
func calculate(f figure2D) {
	fmt.Println("Area: ", f.area())
}
func InterfaceStructsFunction() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}
	calculate(mySquare)
	calculate(myRectangle)
	//Interface lists
	myInterface := []interface{}{"Hola", 12, 4.21}
	fmt.Println(myInterface)
}
