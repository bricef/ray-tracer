package ray

import (
	"testing"

	"github.com/bricef/ray-tracer/entity"
	q "github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/utils"
)

func Sphere() entity.Entity {
	return entity.New()
}

func IntersectTestHelper(t *testing.T, ray Ray, e entity.Entity, expected []float64) {
	xs := intersects(ray, e)
	for i, v := range expected {
		if !utils.AlmostEqual(xs[i], v) {
			t.Errorf("Ray %v sphere intersect failure. Expected %v, got %v", ray, expected, xs)
		}
	}
}

func TestRayIntersectSphere(t *testing.T) {
	IntersectTestHelper(
		t,
		New(q.NewPoint(0, 0, -5), q.NewVector(0, 0, 1)),
		Sphere(),
		[]float64{4.0, 6.0},
	)
	IntersectTestHelper(
		t,
		New(q.NewPoint(0, 1, -5), q.NewVector(0, 0, 1)),
		Sphere(),
		[]float64{5.0, 5.0},
	)
	IntersectTestHelper(
		t,
		New(q.NewPoint(0, 2, -5), q.NewVector(0, 0, 1)),
		Sphere(),
		[]float64{},
	)
	IntersectTestHelper(
		t,
		New(q.NewPoint(0, 0, 0), q.NewVector(0, 0, 1)),
		Sphere(),
		[]float64{-1, 1},
	)
	IntersectTestHelper(
		t,
		New(q.NewPoint(0, 0, 5), q.NewVector(0, 0, 1)),
		Sphere(),
		[]float64{-6.0, -4.0},
	)
}
