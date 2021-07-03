package canvas

import (
	"testing"

	"github.com/bricef/ray-tracer/color"
)

func TestCanvasCreation(t *testing.T) {
	width, height := 10, 20
	c := New(width, height)
	if !(c.Width == 10 && c.Height == 20) {
		t.Errorf("Failed to create canvas with correct dimension. Got %vx%v expected %vx%v.", c.Width, c.Height, width, height)
	}
	for _, column := range c.pixels {
		for _, pixel := range column {
			pixel.Equal(color.New(0, 0, 0))
		}
	}
}

func TestWriteToCanvas(t *testing.T) {
	c := New(10, 10)
	red := color.New(1, 0, 0)
	c.Set(1, 1, red)
	pixel, err := c.Get(1, 1)

	if !pixel.Equal(red) || err != nil {
		t.Errorf("Failed to set pixel on canvas. Expected %v, got %v. (%v)", red, pixel, err)
	}
}
