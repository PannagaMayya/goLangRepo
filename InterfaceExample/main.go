package main

import (
	"fmt"
	"io"
	"os"
)

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
type rectangle struct {
	length float64
	height float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (r rectangle) getArea() float64 {
	return r.length * r.height
}
func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}
func main() {
	//Part 1 Area calculation
	sq := square{2.0}
	re := rectangle{2.54, 8.23}
	tr := triangle{3.0, 4.0}
	printArea(sq)
	printArea(re)
	printArea(tr)
	//Part 2 - Reading file qwerty.txt which is passed in the args as go run main.go qwerty.txt
	data, _ := os.Open(os.Args[1])
	io.Copy(os.Stdout, data)

}
