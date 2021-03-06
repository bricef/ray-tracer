package ray_test

import (
	"testing"

	m "math"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/utils"
)

func PositionTestHelper(t *testing.T, r ray.Ray, d float64, expected math.Quaternion) {
	result := r.Position(d)
	if !result.Equal(expected) {
		t.Errorf("Position failed. Ray %v with t=%v. Expected %v, got %v", r, d, expected, result)
	}
}

func TestRayPosition(t *testing.T) {
	r := ray.NewRay(math.NewPoint(2, 3, 4), math.NewVector(1, 0, 0))

	PositionTestHelper(t, r, 0.0, math.NewPoint(2, 3, 4))
	PositionTestHelper(t, r, 1.0, math.NewPoint(3, 3, 4))
	PositionTestHelper(t, r, -1.0, math.NewPoint(1, 3, 4))
	PositionTestHelper(t, r, 2.5, math.NewPoint(4.5, 3, 4))
}

func IntersectTestHelper(t *testing.T, r ray.Ray, e core.Entity, expected []float64) {
	xs := r.Intersect(e)
	for i, v := range expected {
		if !utils.AlmostEqual(xs.All[i].T, v) {
			t.Errorf("Ray %v sphere intersect failure. Expected %v, got %v", r, expected, xs)
		}
	}
}

func TestRayIntersectSphere(t *testing.T) {
	sphere := entities.NewSphere()

	IntersectTestHelper(
		t,
		ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1)),
		sphere,
		[]float64{4.0, 6.0},
	)
	IntersectTestHelper(
		t,
		ray.NewRay(math.NewPoint(0, 1, -5), math.NewVector(0, 0, 1)),
		sphere,
		[]float64{5.0, 5.0},
	)
	IntersectTestHelper(
		t,
		ray.NewRay(math.NewPoint(0, 2, -5), math.NewVector(0, 0, 1)),
		sphere,
		[]float64{},
	)
	IntersectTestHelper(
		t,
		ray.NewRay(math.NewPoint(0, 0, 0), math.NewVector(0, 0, 1)),
		sphere,
		[]float64{-1, 1},
	)
	IntersectTestHelper(
		t,
		ray.NewRay(math.NewPoint(0, 0, 5), math.NewVector(0, 0, 1)),
		sphere,
		[]float64{-6.0, -4.0},
	)
}

func TestRayIntersectsHaveEntities(t *testing.T) {
	a := entities.NewSphere()
	b := entities.NewSphere()
	r := ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1))
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
	a := entities.NewSphere()
	r := ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1))
	xs := r.Intersect(a)
	if !(xs.Hit.Entity == a) {
		t.Errorf("Intersection %v failed to provide correct entity. Expected %v, got %v", xs.Hit, a, xs.Hit.Entity)
	}
}

func TestRaysAreTransformable(t *testing.T) {
	r := ray.NewRay(math.NewPoint(1, 2, 3), math.NewVector(0, 1, 0))
	tr := math.NewTransform().Translate(3, 4, 5)
	nr := r.Transform(tr)

	if !nr.Origin().Equal(math.NewPoint(4, 6, 8)) || !nr.Direction().Equal(math.NewVector(0, 1, 0)) {
		t.Errorf("Failed to transform Ray %v with %v. Got %v", r, tr, nr)
	}
}

func TestRaysAreScalable(t *testing.T) {
	r := ray.NewRay(math.NewPoint(1, 2, 3), math.NewVector(0, 1, 0))
	tr := math.NewTransform().Scale(2, 3, 4)
	nr := r.Transform(tr)

	if !nr.Origin().Equal(math.NewPoint(2, 6, 12)) || !nr.Direction().Equal(math.NewVector(0, 3, 0)) {
		t.Errorf("Failed to transform Ray %v with %v. Got %v", r, tr, nr)
	}
}

func TestRayIsTransformedBeforeIntersect(t *testing.T) {
	r := ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1))
	s := entities.NewSphere().Scale(2, 2, 2)
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

func TestIntersectionsHaveComputedPoint(t *testing.T) {
	r := ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1))
	s := entities.NewSphere()
	xs := r.Intersect(s)
	x := xs.All[0]
	if !(x.Entity == s &&
		x.Point.Equal(math.NewPoint(0, 0, -1)) &&
		x.EyeVector.Equal(math.NewVector(0, 0, -1)) &&
		x.Normal.Equal(math.NewVector(0, 0, -1))) {
		t.Errorf("Failed to compute intersection helpers")
	}
}

func TestIntersectionsHaveInsideFalse(t *testing.T) {
	r := ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1))
	s := entities.NewSphere()
	xs := r.Intersect(s)
	if xs.Hit.Inside != false {
		t.Errorf("Intersection thinks it's inside object when it is not.")
	}
}

func TestIntersectionsHaveInsideTrue(t *testing.T) {
	r := ray.NewRay(math.NewPoint(0, 0, 0), math.NewVector(0, 0, 1))
	s := entities.NewSphere()
	xs := r.Intersect(s)
	if xs.Hit.Inside != true {
		t.Errorf("Intersection thinks it's outside object when it is inside.")
	}
	if !xs.Hit.Normal.Equal(math.NewVector(0, 0, -1)) {
		t.Errorf("Normal vector not computed correctly when inside object")
	}
}

func TestMergeIntersectionsNeverNegativeHit(t *testing.T) {
	a := &ray.Intersections{
		All: []*ray.Intersection{
			{T: -4},
			{T: -2},
		},
		Hit: nil,
	}
	b := &ray.Intersections{
		All: []*ray.Intersection{
			{T: -5},
			{T: -6},
		},
		Hit: nil,
	}
	xs := a.Merge(b)
	if xs.Hit != nil {
		t.Errorf("Merged intersection hit is negative")
	}
}

func TestOverPointAboveSurface(t *testing.T) {
	r := ray.NewRay(
		math.NewPoint(0, 0, -5),
		math.NewVector(0, 0, 1))
	s := entities.NewSphere().Translate(0, 0, 1)
	xs := r.Intersect(s)
	result := xs.Hit.OverPoint.Z()
	if !((result < -utils.Epsilon/2.0) && xs.Hit.Point.Z() > result) {
		t.Errorf("Failed to compute overpoint. Got %v", xs.Hit.OverPoint)
	}

}

func TestIntesectionsHaveReflectVector(t *testing.T) {
	shape := entities.NewPlane()

	r := ray.NewRay(
		math.NewPoint(0, 1, -1),
		math.NewVector(0, -m.Sqrt2/2, m.Sqrt2/2),
	)

	result := r.Hit(shape).ReflectVector
	expected := math.NewVector(0, m.Sqrt2/2, m.Sqrt2/2)

	if !result.Equal(expected) {
		t.Errorf("Failed to calculate reflection vector. Expected %v, got %v", expected, result)
	}

}

func TestReflectionOnNonReflectiveSurface(t *testing.T) {
	s := scene.DefaultScene()

	r := ray.NewRay(
		math.NewPoint(0, 0, 0),
		math.NewVector(0, 0, 1),
	)

	e := s.Entities[1]
	e.GetMaterial().SetAmbient(1.0)

	xs := r.Intersect(e)

	result := s.ReflectedContribution(xs.Hit, 10)
	expected := color.Black

	if !result.Equal(expected) {
		t.Errorf("Reflection on matt object failed. Expected %v, got %v", expected, result)
	}
}

func TestReflectionOnReflectiveSurfaceReturnsReflectedColor(t *testing.T) {
	s := scene.DefaultScene()

	e := entities.NewPlane()
	e.Translate(0, -1, 0)
	e.GetMaterial().SetReflective(0.5)

	s.Add(e)

	r := ray.NewRay(
		math.NewPoint(0, 0, -3),
		math.NewVector(0, -m.Sqrt2/2, m.Sqrt2/2),
	)

	xs := r.Intersect(e)

	result := s.ReflectedContribution(xs.Hit, 10)
	expected := color.New(0.19033, 0.23792, 0.14274)
	if !result.Equal(expected) {
		t.Errorf("Failed to compute reflected color. Expected %v, got %v", expected, result)
	}
}

func TestRefractionIndicesArePresentOnIntersection(t *testing.T) {
	s := scene.NewScene()

	a := entities.NewGlassSphere().Scale(2, 2, 2)
	s.Add(a)

	b := entities.NewGlassSphere().Translate(0, 0, -0.25)
	b.GetMaterial().SetRefractiveIndex(2.0)
	s.Add(b)

	c := entities.NewGlassSphere().Translate(0, 0, 0.25)
	c.GetMaterial().SetRefractiveIndex(2.5)
	s.Add(c)

	r := ray.NewRay(
		math.NewPoint(0, 0, -4),
		math.NewVector(0, 0, 1),
	)

	xs := s.Intersections(r)

	// fmt.Printf("%v\n", xs.All)

	type helper struct {
		N1 float64
		N2 float64
	}

	tests := []helper{
		{1.0, 1.5},
		{1.5, 2.0},
		{2.0, 2.5},
		{2.5, 2.5},
		{2.5, 1.5},
		{1.5, 1.0},
	}
	// fmt.Printf("%v", tests)
	for i, x := range xs.All {
		if !(x.N1 == tests[i].N1 && x.N2 == tests[i].N2) {
			t.Errorf("Intersection does not have correct indices of refraction. Expected %v->%v, got %v->%v", tests[i].N1, tests[i].N2, x.N1, x.N2)
		}
	}

}

func TestUnderPointBelowSurface(t *testing.T) {
	r := ray.NewRay(
		math.NewPoint(0, 0, -5),
		math.NewVector(0, 0, 1),
	)
	shape := entities.NewGlassSphere().Translate(0, 0, 1)
	xs := r.Intersect(shape)

	if !(xs.Hit.UnderPoint.Z() > utils.Epsilon/2) && (xs.Hit.Point.Z() < xs.Hit.UnderPoint.Z()) {
		t.Errorf("Failed to compute underpoint. Got %v. Expected Z to be less than Z in %v", xs.Hit.UnderPoint, xs.Hit.Point)
	}

}
