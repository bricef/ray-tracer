package test

import (
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
			"Reflection helpers can help",
			Transform().ReflectX().ReflectY().ReflectZ(),
			Point(2, 3, 4).Quaternion,
			Point(-2, -3, -4).Quaternion,
		},
	}
	for _, c := range cases {
		result := c.transform.Apply(c.input)
		if !result.Equal(c.expected) {
			t.Errorf("FAILED [%s]: transform %v, input %v, expected %v, got %v", c.name, c.transform, c.input, c.expected, result)
		}
	}
}
