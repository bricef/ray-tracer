package entity

import (
	"testing"

	"github.com/bricef/ray-tracer/matrix"
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
