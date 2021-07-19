package entity_test

import (
	m "math"
	"testing"

	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/math"
)

func TestSphereHasDefaultTransform(t *testing.T) {
	s := entities.NewSphere()
	if !s.Transform().GetMatrix().Equal(math.Identity(4)) {
		t.Errorf("Default transform for %v is not %v", s, math.Identity(4))
	}
}

func TestSphereTransformCanChaneg(t *testing.T) {
	s := entities.NewSphere().Translate(2, 3, 4)
	tr := math.NewTransform().Translate(2, 3, 4)
	if !s.Transform().GetMatrix().Equal(tr.GetMatrix()) {
		t.Errorf("Failed to set a sphere's trasnform")
	}
}

func TestSphereEntityNormalsWhenTranslated(t *testing.T) {
	s := entities.NewSphere().Translate(0, 1, 0)
	p := math.NewPoint(0, 1+1/m.Sqrt2, -1/m.Sqrt2)

	got := s.Normal(p)
	expected := math.NewVector(0, 1/m.Sqrt2, -1/m.Sqrt2)

	if !got.Equal(expected) {
		t.Errorf("Normal not translated. \n\tMesh %v, point %v. \n\tExpected %v, \n\tgot %v", s.GetMesh(), p, expected, got)
	}
}

func TestSphereEntityNormalsWhenTransformed(t *testing.T) {
	s := entities.NewSphere().Scale(1, 0.5, 1).RotateZ(m.Pi / 5)
	p := math.NewPoint(0, m.Sqrt2/2.0, -m.Sqrt2/2.0)

	got := s.Normal(p)
	expected := math.NewVector(0, 0.970142500, -0.242535625)

	if !got.Equal(expected) {
		t.Errorf("Normal not tranformed. \n\tMesh %v, point %v. \n\tExpected %v, \n\tgot %v", s.GetMesh(), p, expected, got)
	}
}
