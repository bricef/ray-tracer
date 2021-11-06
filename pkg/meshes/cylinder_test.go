package meshes_test

import (
	"fmt"
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
		fmt.Printf("test.expected: %v ts: %v\n", test.expected, ts)
		if !utils.AlmostEqual(ts[0], test.expected[0]) || !utils.AlmostEqual(ts[1], test.expected[1]) {
			t.Errorf("Cylinder hit failed. Expected %v, got %v", test.expected, ts)
		}
	}
}
