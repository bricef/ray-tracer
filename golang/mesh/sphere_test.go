package mesh

import (
	"testing"

	"github.com/bricef/ray-tracer/quaternion"
)

type TestCase struct {
	Mesh     Mesh
	Point    quaternion.Quaternion
	Expected quaternion.Quaternion
}

func TestSphereNormals(t *testing.T) {

	cases := []TestCase{
		{
			NewSphere(),
			quaternion.NewPoint(1, 0, 0),
			quaternion.NewVector(1, 0, 0),
		},
	}

	for _, c := range cases {
		got := c.Mesh.Normal(c.Point)
		if !got.Equal(c.Expected) {
			t.Errorf("Normal at %v failed for %v. Expected %v, got %v", c.Point, c.Mesh, c.Expected, got)
		}
	}

}
