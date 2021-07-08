package mesh

import (
	"math"
	"testing"

	"github.com/bricef/ray-tracer/quaternion"
)

type TestCase struct {
	Mesh     Mesh
	Point    quaternion.Quaternion
	Expected quaternion.Quaternion
}

var k float64 = math.Sqrt(3) / 3.0

func TestSphereNormals(t *testing.T) {

	cases := []TestCase{
		{
			NewSphere(),
			quaternion.NewPoint(1, 0, 0),
			quaternion.NewVector(1, 0, 0),
		},
		{
			NewSphere(),
			quaternion.NewPoint(0, 1, 0),
			quaternion.NewVector(0, 1, 0),
		},
		{
			NewSphere(),
			quaternion.NewPoint(0, 0, 1),
			quaternion.NewVector(0, 0, 1),
		},
		{
			NewSphere(),
			quaternion.NewPoint(k, k, k),
			quaternion.NewVector(k, k, k),
		},
	}

	for _, c := range cases {
		got := c.Mesh.Normal(c.Point)
		if !got.Equal(c.Expected) {
			t.Errorf("Normal at %v failed for %v. Expected %v, got %v", c.Point, c.Mesh, c.Expected, got)
		}
	}

}

func TestSphereNormalsAreNormalised(t *testing.T) {
	s := NewSphere()
	normal := s.Normal(quaternion.NewPoint(k, k, k))

	if normal.Magnitude() != 1.0 {
		t.Errorf("Sphere normal %v. Expected magnitude 1.0, got %v", normal, normal.Magnitude())
	}
}
