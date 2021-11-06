package meshes_test

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/entity"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/meshes"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/utils"
)

func TestCylinderRayMiss(t *testing.T) {
	e := entity.NewEntity()
	cm := meshes.CylinderMesh()
	e.AddComponent(cm)

	rays := []ray.Ray{
		ray.NewRay(math.NewPoint(1, 0, 0), math.NewVector(0, 1, 0)),
		ray.NewRay(math.NewPoint(0, 0, 0), math.NewVector(0, 1, 0)),
		ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(1, 1, 1)),
	}

	for _, r := range rays {
		xs := r.Intersect(e)
		if xs.Hit != nil {
			t.Errorf("Ray should miss cylinder, but got hit. %v intersect at  %v", r, xs.Hit.Point)
		}
	}
}

func TestCylinderHit(t *testing.T) {
	cm := meshes.CylinderMesh()

	type Test struct {
		ray      ray.Ray
		expected []float64
	}

	tests := []Test{
		{
			ray.NewRay(math.NewPoint(1, 0, -5), math.NewVector(0, 0, 1)),
			[]float64{5, 5}},
		{
			ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1)),
			[]float64{4, 6}},
		{
			ray.NewRay(math.NewPoint(0.5, 0, -5), math.NewVector(0.1, 1, 1)),
			[]float64{6.80798, 7.08872}},
	}

	for _, test := range tests {
		ts := cm.Intersect(test.ray)
		if !utils.AlmostEqual(ts[0], test.expected[0]) || !utils.AlmostEqual(ts[1], test.expected[1]) {
			t.Errorf("Cylinder hit failed. Expected %v, got %v", test.expected, ts)
		}
	}
}

func TestCylinderNormal(t *testing.T) {
	cm := meshes.CylinderMesh()

	type Test struct {
		p math.Point
		n math.Vector
	}

	tests := []Test{
		{math.NewPoint(1, 0, 0), math.NewVector(1, 0, 0)},
		{math.NewPoint(0, 5, -1), math.NewVector(0, 0, -1)},
		{math.NewPoint(0, -2, 1), math.NewVector(0, 0, 1)},
		{math.NewPoint(-1, 1, 0), math.NewVector(-1, 0, 0)},
	}

	for _, test := range tests {
		n := cm.Normal(test.p)
		if !n.Equal(test.n) {
			t.Errorf("Failed to compute normal on cylinder. Expected %v, got %v", test.n, n)
		}
	}
}

func TestCylindersHaveLimits(t *testing.T) {
	cm := meshes.CylinderMeshLimited(1, 2)

	type Test struct {
		r  ray.Ray
		ts int
	}

	tests := []Test{
		{ray.NewRay(math.NewPoint(0, 1.5, 0), math.NewVector(0.1, 1, 0)), 0},
		{ray.NewRay(math.NewPoint(0, 3, -5), math.NewVector(0, 0, 1)), 0},
		{ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1)), 0},
		{ray.NewRay(math.NewPoint(0, 2, -5), math.NewVector(0, 0, 1)), 0},
		{ray.NewRay(math.NewPoint(0, 1, -5), math.NewVector(0, 0, 1)), 0},
		{ray.NewRay(math.NewPoint(0, 1.5, -5), math.NewVector(0, 0, 1)), 2},
	}

	for _, test := range tests {
		ts := cm.Intersect(test.r)
		if len(ts) != test.ts {
			t.Errorf("Did not get correct number of intersection for %v with finite cylinder. Expected %v, got %v", test.r, test.ts, ts)
		}
	}
}

func TestClosedCylinderIntersect(t *testing.T) {
	cm := meshes.CylinderClosedMesh(1, 2)

	type Test struct {
		r  ray.Ray
		ts int
	}

	tests := []Test{
		{ray.NewRay(math.NewPoint(0, 3, 0), math.NewVector(0, -1, 0)), 2},
		{ray.NewRay(math.NewPoint(0, 3, -2), math.NewVector(0, -1, 2)), 2},
		{ray.NewRay(math.NewPoint(0, 4, -2), math.NewVector(0, -1, 1)), 2},
		{ray.NewRay(math.NewPoint(0, 0, -2), math.NewVector(0, 1, 2)), 2},
		{ray.NewRay(math.NewPoint(0, -1, -2), math.NewVector(0, 1, 1)), 2},
	}

	for _, test := range tests {
		ts := cm.Intersect(test.r)
		if len(ts) != test.ts {
			t.Errorf("Closed cylinder intersect failure with %v. Expected %v intersect. Got %v", test.r, test.ts, ts)
		}
	}

}
