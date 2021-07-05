package main

import (
	// "github.com/bricef/ray-tracer/canvas"
	// "github.com/bricef/ray-tracer/color"
	"fmt"

	matrix "github.com/bricef/ray-tracer/matrix"
)

func main() {
	// c := canvas.New(100, 100)
	// for x := 0; x < c.Width; x++ {
	// 	for y := 0; y < c.Height; y++ {
	// 		c.Set(x, y, color.New(.1, .1, .1))
	// 	}
	// }
	// c.Set(0, 0, color.New(1, 1, 1))

	// c.WritePNG("test.png")

	m := matrix.Identity(4)
	i, _ := m.Inverse()
	fmt.Printf("Inverse of the identity %v is %v", m, i)

}
