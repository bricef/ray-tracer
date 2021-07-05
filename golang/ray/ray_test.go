package ray

import (
	"testing"

	"github.com/bricef/ray-tracer/entity"
	q "github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/shapes"
	"github.com/bricef/ray-tracer/transform"
	"github.com/bricef/ray-tracer/utils"
)

func IntersectTestHelper(t *testing.T, ray Ray, e *entity.Entity, expected []float64) {
	xs := ray.Intersect(e)
	for i, v := range expected {
		if !utils.AlmostEqual(xs.All[i].T, v) {
			t.Errorf("Ray %v sphere intersect failure. Expected %v, got %v", ray, expected, xs)
		}
	}
}

func TestRayIntersectSphere(t *testing.T) {
	IntersectTestHelper(
		t,
		NewRay(q.NewPoint(0, 0, -5), q.NewVector(0, 0, 1)),
		shapes.Sphere(),
		[]float64{4.0, 6.0},
	)
	IntersectTestHelper(
		t,
		NewRay(q.NewPoint(0, 1, -5), q.NewVector(0, 0, 1)),
		shapes.Sphere(),
		[]float64{5.0, 5.0},
	)
	IntersectTestHelper(
		t,
		NewRay(q.NewPoint(0, 2, -5), q.NewVector(0, 0, 1)),
		shapes.Sphere(),
		[]float64{},
	)
	IntersectTestHelper(
		t,
		NewRay(q.NewPoint(0, 0, 0), q.NewVector(0, 0, 1)),
		shapes.Sphere(),
		[]float64{-1, 1},
	)
	IntersectTestHelper(
		t,
		NewRay(q.NewPoint(0, 0, 5), q.NewVector(0, 0, 1)),
		shapes.Sphere(),
		[]float64{-6.0, -4.0},
	)
}

func TestRayIntersectsHaveEntities(t *testing.T) {
	a := shapes.Sphere()
	b := shapes.Sphere()
	r := NewRay(q.NewPoint(0, 0, -5), q.NewVector(0, 0, 1))
	as := r.Intersect(a)
	bs := r.Intersect(b)

	for _, x := range as.All {
		if x.Entity != a {
			t.Errorf("Got wrong entity for intersection")
		}
	}

	for _, x := range bs.All {
		if x.Entity != b {
			t.Errorf("Got wrong entity for intersection")
		}
	}
}

func TestIntersectHaveHits(t *testing.T) {
	a := shapes.Sphere()
	r := NewRay(q.NewPoint(0, 0, -5), q.NewVector(0, 0, 1))
	xs := r.Intersect(a)
	if !(xs.Hit.Entity == a) {
		t.Errorf("Intersection %v failed to provide correct entity. Expected %v, got %v", xs.Hit, a, xs.Hit.Entity)
	}
}

func TestRaysAreTransformable(t *testing.T) {
	r := NewRay(q.NewPoint(1, 2, 3), q.NewVector(0, 1, 0))
	tr := transform.NewTransform().Translate(3, 4, 5)
	nr := r.Transform(tr)

	if !nr.Origin.Equal(q.NewPoint(4, 6, 8)) || !nr.Direction.Equal(q.NewVector(0, 1, 0)) {
		t.Errorf("Failed to transform Ray %v with %v. Got %v", r, tr, nr)
	}
}

func TestRaysAreScalable(t *testing.T) {
	r := NewRay(q.NewPoint(1, 2, 3), q.NewVector(0, 1, 0))
	tr := transform.NewTransform().Scale(2, 3, 4)
	nr := r.Transform(tr)

	if !nr.Origin.Equal(q.NewPoint(2, 6, 12)) || !nr.Direction.Equal(q.NewVector(0, 3, 0)) {
		t.Errorf("Failed to transform Ray %v with %v. Got %v", r, tr, nr)
	}
}

func TestRayIsTransformedBeforeIntersect(t *testing.T) {
	r := NewRay(q.NewPoint(0, 0, -5), q.NewVector(0, 0, 1))
	s := shapes.Sphere()
	s.SetTransform(transform.NewTransform().Scale(2, 2, 2))
	xs := r.Intersect(s)
	if len(xs.All) != 2 {
		t.Errorf("Wrong number of intersects. Expected 2, got %v", len(xs.All))
	}
	if xs.All[0].T != 3 {
		t.Errorf("First intersection at wrong distance. Expected 3.0, got %v", xs.All[0].T)
	}
	if xs.All[1].T != 7 {
		t.Errorf("First intersection at wrong distance. Expected 7.0, got %v", xs.All[1].T)
	}

}
