package meshes_test

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/meshes"
	"github.com/bricef/ray-tracer/pkg/ray"
)

type HitCase struct {
	Description string
	Origin      math.Point
	Direction   math.Vector
	T1          float64
	T2          float64
}

func TestCubeMeshIntersectsPerpendicularly(t *testing.T) {
	tests := []HitCase{
		{"+x", math.NewPoint(5, 0.5, 0), math.NewVector(-1, 0, 0), 4, 6},
		{"-x", math.NewPoint(-5, 0.5, 0), math.NewVector(1, 0, 0), 4, 6},
		{"+y", math.NewPoint(0.5, 5, 0), math.NewVector(0, -1, 0), 4, 6},
		{"-y", math.NewPoint(0.5, -5, 0), math.NewVector(0, 1, 0), 4, 6},
		{"+z", math.NewPoint(0.5, 0, 5), math.NewVector(0, 0, -1), 4, 6},
		{"-z", math.NewPoint(0.5, 0, -5), math.NewVector(0, 0, 1), 4, 6},
		{"inside", math.NewPoint(0, 0.5, 0), math.NewVector(0, 0, 1), -1, 1},
	}
	c := meshes.CubeMesh()
	for _, test := range tests {
		r := ray.NewRay(test.Origin, test.Direction)
		ts := c.Intersect(r)
		if ts[0] != test.T1 || ts[1] != test.T2 {
			t.Errorf("Cube intersection failure %v. Expected t1=%v, t2=%v. Got t1=%v, t2=%v", test.Description, test.T1, test.T2, ts[0], ts[1])
		}
	}
}

type MissCase struct {
	Origin    math.Point
	Direction math.Vector
}

func TestCubeRayIntersectionMiss(t *testing.T) {
	tests := []MissCase{
		{math.NewPoint(-2, 0, 0), math.NewVector(0.27, 0.53, 0.80)},
		{math.NewPoint(0, -2, 0), math.NewVector(0.80, 0.26, 0.53)},
		{math.NewPoint(0, 0, -2), math.NewVector(0.53, 0.80, 0.27)},
		{math.NewPoint(2, 0, 2), math.NewVector(0, 0, -1)},
		{math.NewPoint(0, 2, 2), math.NewVector(0, -1, 0)},
		{math.NewPoint(2, 2, 0), math.NewVector(-1, 0, 0)},
	}
	c := meshes.CubeMesh()
	for _, test := range tests {
		r := ray.NewRay(test.Origin, test.Direction)
		ts := c.Intersect(r)
		if len(ts) != 0 {
			t.Errorf("Expected cube ray intersection to miss, but got %v.", ts)
		}
	}
}

func TestCubeNormal(t *testing.T) {
	c := meshes.CubeMesh()
	helper := func(p math.Point, n math.Vector) {
		got := c.Normal(p)
		if !got.Equal(n) {
			t.Errorf("Failed to compute normal for a cube mesh at %v. Expected %v, got %v.", p, n, got)
		}
	}
	helper(math.NewPoint(1, 0.5, -0.8), math.NewVector(1, 0, 0))
	helper(math.NewPoint(-1, -0.2, 0.9), math.NewVector(-1, 0, 0))
	helper(math.NewPoint(-0.4, 1, -0.1), math.NewVector(0, 1, 0))
	helper(math.NewPoint(0.3, -1, -0.7), math.NewVector(0, -1, 0))
	helper(math.NewPoint(-0.6, 0.3, 1), math.NewVector(0, 0, 1))
	helper(math.NewPoint(0.4, 0.4, -1), math.NewVector(0, 0, -1))
	helper(math.NewPoint(1, 1, 1), math.NewVector(1, 0, 0))
	helper(math.NewPoint(-1, -1, -1), math.NewVector(-1, 0, 0))
}
