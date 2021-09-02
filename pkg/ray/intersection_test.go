package ray_test

import (
	"testing"

	m "math"

	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/utils"
)

func TestShclickUnderTotalInternalRefraction(t *testing.T) {
	shape := entities.NewGlassSphere()
	r := ray.NewRay(
		math.NewPoint(0, 0, m.Sqrt2/2.0),
		math.NewVector(0, 1, 0),
	)
	xs := r.GetIntersections([]core.Entity{shape})
	reflectance := xs.Hit.Schlick()
	expected := 1.0
	if reflectance != expected {
		t.Errorf("Reflectance under total internal refraction should be %v. Got %v.", expected, reflectance)
	}

}

func TestSchlickForPerpendicularRay(t *testing.T) {
	shape := entities.NewGlassSphere()
	r := ray.NewRay(
		math.NewPoint(0, 0, 0),
		math.NewVector(0, 1, 0),
	)
	xs := r.GetIntersections([]core.Entity{shape})
	reflectance := xs.Hit.Schlick()
	expected := 0.04

	if !utils.AlmostEqual(reflectance, expected) {
		t.Errorf("Perpendicular reflectance incorrect. Expected %v, got %v.", expected, reflectance)
	}
}

func TestSchlickForSmallAngles(t *testing.T) {
	shape := entities.NewGlassSphere()
	r := ray.NewRay(
		math.NewPoint(0, 0.99, -2),
		math.NewVector(0, 0, 1),
	)
	xs := r.GetIntersections([]core.Entity{shape})

	reflectance := xs.Hit.Schlick()
	expected := 0.48881 // Differs from book: 0.48873

	if !utils.AlmostEqual(reflectance, expected) {
		t.Errorf("Incorrect Reflectance for small angles. Expected %v, got %v.", expected, reflectance)
	}

}
