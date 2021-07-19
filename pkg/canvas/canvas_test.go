package canvas

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/color"
)

func TestCanvasCreation(t *testing.T) {
	width, height := 10, 20
	c := NewImageCanvas(width, height)
	if !(c.width == 10 && c.height == 20) {
		t.Errorf("Failed to create canvas with correct dimension. Got %vx%v expected %vx%v.", c.width, c.height, width, height)
	}
	for _, column := range c.pixels {
		for _, pixel := range column {
			pixel.Equal(color.New(0, 0, 0))
		}
	}
}

func TestWriteToCanvas(t *testing.T) {
	c := NewImageCanvas(10, 10)
	red := color.New(1, 0, 0)
	c.Set(1, 1, red)
	pixel, err := c.Get(1, 1)

	if !pixel.Equal(red) || err != nil {
		t.Errorf("Failed to set pixel on canvas. Expected %v, got %v. (%v)", red, pixel, err)
	}
}

func TestCanvasPixelIteratorIteratesOverAllPixels(t *testing.T) {
	c := NewImageCanvas(10, 10)
	pi := c.Pixels()
	count := 0
	for pi.More() {
		pi.Get()
		count += 1
	}
	expected := 100
	if count != expected {
		t.Errorf("Pixel iterator returned wrong number of pixels. Expected %v got %v.", expected, count)
	}
}
