package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
}

// Struct yang memenuhi kontrak Shape
type Rectangle struct {
	Width, Height float64
}

// Implementasi metode Area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape = Rectangle{Width: 10, Height: 5}

	fmt.Println("Luas Rectangle:", s.Area())
}
