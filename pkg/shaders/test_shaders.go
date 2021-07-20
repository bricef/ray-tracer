package shaders

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/math"
)

func TestSimpleShader(t *testing.T) {

	c1 := Pigment(color.Red)(math.NewPoint(0, 23, 33))
	c2 := Pigment(color.Green)(math.NewPoint(0, 23, 33))

	if !c1.Equal(color.Red) {
		t.Errorf("Failed to set pigment color. Expected %v, got %v", color.Red, c1)
	}

	if !c2.Equal(color.Green) {
		t.Errorf("Failed to set pigment color. Expected %v, got %v", color.Green, c2)
	}

}

func TestStriped(t *testing.T) {
	shader := Striped(
		Pigment(color.White),
		Pigment(color.Black),
	)

	type test struct {
		point    math.Point
		expected color.Color
	}

	tests := []test{
		// constant in y
		{math.NewPoint(0, 0, 0), color.White},
		{math.NewPoint(0, 1, 0), color.White},
		{math.NewPoint(0, 2, 0), color.White},
		// Constant in Z
		{math.NewPoint(0, 0, 0), color.White},
		{math.NewPoint(0, 0, 1), color.White},
		{math.NewPoint(0, 0, 2), color.White},
		// Changes in X
		{math.NewPoint(0, 0, 0), color.White},
		{math.NewPoint(0.9, 0, 0), color.White},
		{math.NewPoint(1, 0, 0), color.Black},
		{math.NewPoint(-0.1, 0, 0), color.Black},
		{math.NewPoint(-1, 0, 0), color.Black},
		{math.NewPoint(-1.1, 0, 0), color.White},
	}

	for _, test := range tests {
		result := shader(test.point)
		if !result.Equal(test.expected) {
			t.Errorf("Striped pattern failure at %v. Expected %v, got %v ", test.point, test.expected, result)
		}
	}
}
