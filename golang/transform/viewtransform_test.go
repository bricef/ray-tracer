package transform

import (
	"testing"

	"github.com/bricef/ray-tracer/matrix"
	"github.com/bricef/ray-tracer/quaternion"
)

type Case struct {
	transform Transform
	expected  Transform
}

func TestViewTransformCases(t *testing.T) {
	testCases := []Case{
		{
			ViewTransform(
				quaternion.NewPoint(0, 0, 0),
				quaternion.NewPoint(0, 0, -1),
				quaternion.NewVector(0, 1, 0),
			),
			Transform{matrix.Identity(4)},
		},
		{
			ViewTransform(
				quaternion.NewPoint(0, 0, 0),
				quaternion.NewPoint(0, 0, 1),
				quaternion.NewVector(0, 1, 0),
			),
			NewTransform().Scale(-1, 1, -1),
		},
		{
			ViewTransform(
				quaternion.NewPoint(0, 0, 8),
				quaternion.NewPoint(0, 0, 0),
				quaternion.NewVector(0, 1, 0),
			),
			NewTransform().Translate(0, 0, -8),
		},
		{
			ViewTransform(
				quaternion.NewPoint(1, 3, 2),
				quaternion.NewPoint(4, -2, 8),
				quaternion.NewVector(1, 1, 0),
			),
			Transform{matrix.New([][]float64{
				{-0.50709, 0.50709, 0.67612, -2.36643},
				{0.76772, 0.60609, .12122, -2.82843},
				{-.35857, .59761, -.71714, 0.0},
				{0.0, 0.0, 0.0, 1.0},
			})},
		},
	}

	for _, c := range testCases {
		if !c.transform.Equal(c.expected) {
			t.Errorf("View Transform initialisation failed. Expected %v, got %v.", c.expected, c.transform)
		}
	}
}
