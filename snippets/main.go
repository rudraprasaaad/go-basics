package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	Size float64
}

type Rectangle struct {
	Width, Height float64
}

func (s *Square) Area() float64 {
	return s.Size * s.Size
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	square := Square{Size: 2}
	fmt.Println("Area of square is:", square.Area())

	rect := Rectangle{Width: 5, Height: 4}
	fmt.Println("Area of Rectangle is:", rect.Area())
}
