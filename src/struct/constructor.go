package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func newRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func main() {
	rect := newRectangle(30, 20)
	fmt.Println(rect)

}
