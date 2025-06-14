package main

import (
	"fmt"
)

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() float64 {
	a := float64(r.Width * r.Height)
	return a
}

func main() {
	rect := Rectangle{Width: 10, Height: 20}
	// First itteration
	fmt.Println(rect.Width)
	fmt.Println(rect.Height)
	fmt.Println(rect.Area())

	// Second itteration
	fmt.Println("Enter Witdth: ")
	fmt.Scan(&rect.Width)
	fmt.Println("Enter Height: ")
	fmt.Scan(&rect.Height)

	fmt.Println(rect.Width)
	fmt.Println(rect.Height)
	fmt.Println(rect.Area())
}
