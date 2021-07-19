package meshes_test

import (
	"testing"

	"math"

	"github.com/bricef/ray-tracer/pkg/core"
	m "github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/meshes"
)

type TestCase struct {
	Mesh     core.Mesh
	Point    m.Point
	Expected m.Vector
}

var k float64 = math.Sqrt(3) / 3.0

func TestSphereNormals(t *testing.T) {

	cases := []TestCase{
		{
			meshes.SphereMesh(),
			m.NewPoint(1, 0, 0),
			m.NewVector(1, 0, 0),
		},
		{
			meshes.SphereMesh(),
			m.NewPoint(0, 1, 0),
			m.NewVector(0, 1, 0),
		},
		{
			meshes.SphereMesh(),
			m.NewPoint(0, 0, 1),
			m.NewVector(0, 0, 1),
		},
		{
			meshes.SphereMesh(),
			m.NewPoint(k, k, k),
			m.NewVector(k, k, k),
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
	s := meshes.SphereMesh()
	normal := s.Normal(m.NewPoint(k, k, k))

	if normal.Magnitude() != 1.0 {
		t.Errorf("Sphere normal %v. Expected magnitude 1.0, got %v", normal, normal.Magnitude())
	}
}
