package entity

import (
	"math"
	"testing"

	"github.com/bricef/ray-tracer/matrix"
	"github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/transform"
)

func TestSphereHasDefaultTransform(t *testing.T) {
	s := NewSphere()
	if !s.Transform.Matrix.Equal(matrix.Identity(4)) {
		t.Errorf("Default transform for %v is not %v", s, matrix.Identity(4))
	}
}

func TestSphereTransformCanChaneg(t *testing.T) {
	s := NewSphere()
	tr := transform.NewTransform().Translate(2, 3, 4)
	s.SetTransform(tr)
	if !s.Transform.Matrix.Equal(tr.Matrix) {
		t.Errorf("Failed to set a sphere's trasnform")
	}
}

func TestSphereEntityNormalsWhenTranslated(t *testing.T) {
	s := NewSphere()
	s.SetTransform(transform.NewTransform().Translate(0, 1, 0))
	p := quaternion.NewPoint(0, 1+1/math.Sqrt2, -1/math.Sqrt2)

	got := s.Normal(p)
	expected := quaternion.NewVector(0, 1/math.Sqrt2, -1/math.Sqrt2)

	if !got.Equal(expected) {
		t.Errorf("Normal not translated. \n\tMesh %v, point %v. \n\tExpected %v, \n\tgot %v", s.Mesh, p, expected, got)
	}
}

func TestSphereEntityNormalsWhenTransformed(t *testing.T) {
	s := NewSphere()
	s.SetTransform(transform.NewTransform().Scale(1, 0.5, 1).RotateZ(math.Pi / 5))
	p := quaternion.NewPoint(0, math.Sqrt2/2.0, -math.Sqrt2/2.0)

	got := s.Normal(p)
	expected := quaternion.NewVector(0, 0.970142500, -0.242535625)

	if !got.Equal(expected) {
		t.Errorf("Normal not tranformed. \n\tMesh %v, point %v. \n\tExpected %v, \n\tgot %v", s.Mesh, p, expected, got)
	}
}
