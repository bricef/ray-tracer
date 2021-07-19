package math_test

import (
	m "math"
	"testing"

	"github.com/bricef/ray-tracer/pkg/math"
)

func TestTranslateDoesntChangOriginal(t *testing.T) {
	original := math.NewTransform()
	i := math.Identity(4)
	original.Translate(1, 2, 3)
	if !original.Matrix.Equal(i) {
		t.Errorf("Invalid mutation of matrix %v. Shuld be %v", original, i)
	}

}

type Case struct {
	name      string
	transform math.Transform
	input     math.Quaternion
	expected  math.Quaternion
}

func TestCases(t *testing.T) {
	cases := []Case{
		{
			"Translate Point",
			math.NewTransform().Translate(5, -3, 2),
			math.NewPoint(-3, 4, 5),
			math.NewPoint(2, 1, 7),
		},
		{
			"Translate from Inverse",
			math.NewTransform().Translate(5, -3, 2).Inverse(),
			math.NewPoint(-3, 4, 5),
			math.NewPoint(-8, 7, 3),
		},
		{
			"Translate Vector",
			math.NewTransform().Translate(5, -3, 2),
			math.NewVector(-3, 4, 5),
			math.NewVector(-3, 4, 5),
		},
		{
			"Scaling Point",
			math.NewTransform().Scale(2, 3, 4),
			math.NewPoint(-4, 6, 8),
			math.NewPoint(-8, 18, 32),
		},
		{
			"Scaling Vector",
			math.NewTransform().Scale(2, 3, 4),
			math.NewVector(-4, 6, 8),
			math.NewVector(-8, 18, 32),
		},
		{
			"Scaling a vector by the inverse",
			math.NewTransform().Scale(2, 3, 4).Inverse(),
			math.NewVector(-4, 6, 8),
			math.NewVector(-2, 2, 2),
		},
		{
			"Reflection is scaling negatively",
			math.NewTransform().Scale(-1, 1, 1),
			math.NewPoint(2, 3, 4),
			math.NewPoint(-2, 3, 4),
		},
		{
			"Reflection helpers",
			math.NewTransform().ReflectX().ReflectY().ReflectZ(),
			math.NewPoint(2, 3, 4),
			math.NewPoint(-2, -3, -4),
		},
		{
			"Rotation X",
			math.NewTransform().RotateX(m.Pi / 2),
			math.NewPoint(0, 1, 0),
			math.NewPoint(0, 0, 1),
		},
		{
			"Rotation X, fractional",
			math.NewTransform().RotateX(m.Pi / 4),
			math.NewPoint(0, 1, 0),
			math.NewPoint(0, m.Sqrt2/2, m.Sqrt2/2),
		},
		{
			"Rotation Y",
			math.NewTransform().RotateY(m.Pi / 4),
			math.NewPoint(0, 0, 1),
			math.NewPoint(m.Sqrt2/2, 0, m.Sqrt2/2),
		},
		{
			"Rotation Z",
			math.NewTransform().RotateZ(m.Pi / 4),
			math.NewPoint(0, 1, 0),
			math.NewPoint(-m.Sqrt2/2, m.Sqrt2/2, 0),
		},
		{
			"Shearing X wrt Y",
			math.NewTransform().Shear(1, 0, 0, 0, 0, 0),
			math.NewPoint(2, 3, 4),
			math.NewPoint(5, 3, 4),
		},
		{
			"Shearing X wrt Z",
			math.NewTransform().Shear(0, 1, 0, 0, 0, 0),
			math.NewPoint(2, 3, 4),
			math.NewPoint(6, 3, 4),
		},
		{
			"Shearing Y wrt X",
			math.NewTransform().Shear(0, 0, 1, 0, 0, 0),
			math.NewPoint(2, 3, 4),
			math.NewPoint(2, 5, 4),
		},
		{
			"Shearing Y wrt Z",
			math.NewTransform().Shear(0, 0, 0, 1, 0, 0),
			math.NewPoint(2, 3, 4),
			math.NewPoint(2, 7, 4),
		},
		{
			"Shearing Z wrt X",
			math.NewTransform().Shear(0, 0, 0, 0, 1, 0),
			math.NewPoint(2, 3, 4),
			math.NewPoint(2, 3, 6),
		},
		{
			"Shearing Z wrt Y",
			math.NewTransform().Shear(0, 0, 0, 0, 0, 1),
			math.NewPoint(2, 3, 4),
			math.NewPoint(2, 3, 7),
		},
		{
			"MoveTo",
			math.NewTransform().Translate(3, 5, 6).MoveTo(math.NewPoint(1, 2, 3)),
			math.NewPoint(0, 0, 0),
			math.NewPoint(1, 2, 3),
		},
	}
	for _, c := range cases {
		result := c.transform.Apply(c.input)
		if !result.Equal(c.expected) {
			t.Errorf("FAILED [%s]: transform %v, input %v, expected %v, got %v", c.name, c.transform, c.input, c.expected, result)
		}
	}
}
