package test

import (
	"math"
	"testing"

	m "github.com/bricef/ray-tracer/matrix"
	"github.com/bricef/ray-tracer/quaternion"
	. "github.com/bricef/ray-tracer/raytracer"
	"github.com/bricef/ray-tracer/transform"
)

func TestTranslateDoesntChangOriginal(t *testing.T) {
	original := Transform()
	i := m.Identity(4)
	original.Translate(1, 2, 3)
	if !original.Equal(i) {
		t.Errorf("Invalid mutation of matrix %v. Shuld be %v", original, i)
	}

}

type Case struct {
	name      string
	transform transform.Transform
	input     quaternion.Quaternion
	expected  quaternion.Quaternion
}

func TestCases(t *testing.T) {
	cases := []Case{
		{
			"Translate Point",
			Transform().Translate(5, -3, 2),
			Point(-3, 4, 5).Quaternion,
			Point(2, 1, 7).Quaternion,
		},
		{
			"Translate from Inverse",
			Transform().Translate(5, -3, 2).Inverse(),
			Point(-3, 4, 5).Quaternion,
			Point(-8, 7, 3).Quaternion,
		},
		{
			"Translate Vector",
			Transform().Translate(5, -3, 2),
			Vector(-3, 4, 5).Quaternion,
			Vector(-3, 4, 5).Quaternion,
		},
		{
			"Scaling Point",
			Transform().Scale(2, 3, 4),
			Point(-4, 6, 8).Quaternion,
			Point(-8, 18, 32).Quaternion,
		},
		{
			"Scaling Vector",
			Transform().Scale(2, 3, 4),
			Vector(-4, 6, 8).Quaternion,
			Vector(-8, 18, 32).Quaternion,
		},
		{
			"Scaling a vector by the inverse",
			Transform().Scale(2, 3, 4).Inverse(),
			Vector(-4, 6, 8).Quaternion,
			Vector(-2, 2, 2).Quaternion,
		},
		{
			"Reflection is scaling negatively",
			Transform().Scale(-1, 1, 1),
			Point(2, 3, 4).Quaternion,
			Point(-2, 3, 4).Quaternion,
		},
		{
			"Reflection helpers",
			Transform().ReflectX().ReflectY().ReflectZ(),
			Point(2, 3, 4).Quaternion,
			Point(-2, -3, -4).Quaternion,
		},
		{
			"Rotation X",
			Transform().RotateX(math.Pi / 2),
			Point(0, 1, 0).Quaternion,
			Point(0, 0, 1).Quaternion,
		},
		{
			"Rotation X, fractional",
			Transform().RotateX(math.Pi / 4),
			Point(0, 1, 0).Quaternion,
			Point(0, math.Sqrt2/2, math.Sqrt2/2).Quaternion,
		},
		{
			"Rotation Y",
			Transform().RotateY(math.Pi / 4),
			Point(0, 0, 1).Quaternion,
			Point(math.Sqrt2/2, 0, math.Sqrt2/2).Quaternion,
		},
		{
			"Rotation Z",
			Transform().RotateZ(math.Pi / 4),
			Point(0, 1, 0).Quaternion,
			Point(-math.Sqrt2/2, math.Sqrt2/2, 0).Quaternion,
		},
		{
			"Shearing X wrt Y",
			Transform().Shear(1, 0, 0, 0, 0, 0),
			Point(2, 3, 4).Quaternion,
			Point(5, 3, 4).Quaternion,
		},
		{
			"Shearing X wrt Z",
			Transform().Shear(0, 1, 0, 0, 0, 0),
			Point(2, 3, 4).Quaternion,
			Point(6, 3, 4).Quaternion,
		},
		{
			"Shearing Y wrt X",
			Transform().Shear(0, 0, 1, 0, 0, 0),
			Point(2, 3, 4).Quaternion,
			Point(2, 5, 4).Quaternion,
		},
		{
			"Shearing Y wrt Z",
			Transform().Shear(0, 0, 0, 1, 0, 0),
			Point(2, 3, 4).Quaternion,
			Point(2, 7, 4).Quaternion,
		},
		{
			"Shearing Z wrt X",
			Transform().Shear(0, 0, 0, 0, 1, 0),
			Point(2, 3, 4).Quaternion,
			Point(2, 3, 6).Quaternion,
		},
		{
			"Shearing Z wrt Y",
			Transform().Shear(0, 0, 0, 0, 0, 1),
			Point(2, 3, 4).Quaternion,
			Point(2, 3, 7).Quaternion,
		},
	}
	for _, c := range cases {
		result := c.transform.Apply(c.input)
		if !result.Equal(c.expected) {
			t.Errorf("FAILED [%s]: transform %v, input %v, expected %v, got %v", c.name, c.transform, c.input, c.expected, result)
		}
	}
}
