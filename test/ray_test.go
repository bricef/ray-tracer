package test

import (
	"testing"

	q "github.com/bricef/ray-tracer/quaternion"
	r "github.com/bricef/ray-tracer/ray"
	. "github.com/bricef/ray-tracer/raytracer"
)

func PositionTestHelper(t *testing.T, r r.Ray, d float64, expected q.Quaternion) {
	result := r.Position(d)
	if !result.Equal(expected) {
		t.Errorf("Position failed. Ray %v with t=%v. Expected %v, got %v", r, d, expected, result)
	}
}

func TestRayPosition(t *testing.T) {
	r := Ray(Point(2, 3, 4), Vector(1, 0, 0))

	PositionTestHelper(t, r, 0.0, Point(2, 3, 4))
	PositionTestHelper(t, r, 1.0, Point(3, 3, 4))
	PositionTestHelper(t, r, -1.0, Point(1, 3, 4))
	PositionTestHelper(t, r, 2.5, Point(4.5, 3, 4))
}
