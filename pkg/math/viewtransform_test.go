package math_test

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/math"
)

type TransformTestCase struct {
	transform math.Transform
	expected  math.Transform
}

func TestViewTransformCases(t *testing.T) {
	testCases := []TransformTestCase{
		{
			math.ViewTransform(
				math.NewPoint(0, 0, 0),
				math.NewPoint(0, 0, -1),
				math.NewVector(0, 1, 0),
			),
			math.NewTransform(),
		},
		{
			math.ViewTransform(
				math.NewPoint(0, 0, 0),
				math.NewPoint(0, 0, 1),
				math.NewVector(0, 1, 0),
			),
			math.NewTransform().Scale(-1, 1, -1),
		},
		{
			math.ViewTransform(
				math.NewPoint(0, 0, 8),
				math.NewPoint(0, 0, 0),
				math.NewVector(0, 1, 0),
			),
			math.NewTransform().Translate(0, 0, -8),
		},
		{
			math.ViewTransform(
				math.NewPoint(1, 3, 2),
				math.NewPoint(4, -2, 8),
				math.NewVector(1, 1, 0),
			),
			math.NewTransform().Raw([][]float64{
				{-0.50709, 0.50709, 0.67612, -2.36643},
				{0.76772, 0.60609, .12122, -2.82843},
				{-.35857, .59761, -.71714, 0.0},
				{0.0, 0.0, 0.0, 1.0},
			}),
		},
	}

	for _, c := range testCases {
		if !c.transform.Equal(c.expected) {
			t.Errorf("View Transform initialisation failed. Expected %v, got %v.", c.expected, c.transform)
		}
	}
}
