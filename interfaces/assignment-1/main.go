package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	var sideLength, base, height float64
	fmt.Print("Enter side of square : ")
	fmt.Scanln(&sideLength)
	fmt.Print("Enter base of triangle : ")
	fmt.Scanln(&base)
	fmt.Print("Enter height of triangle : ")
	fmt.Scanln(&height)
	sq := square{sideLength: sideLength}
	t := triangle{base: base, height: height}
	printArea(sq)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (sq square) getArea() float64 {
	return (sq.sideLength * sq.sideLength)
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}
