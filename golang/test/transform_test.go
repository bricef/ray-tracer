package test

import (
	"testing"

	m "github.com/bricef/ray-tracer/matrix"
	. "github.com/bricef/ray-tracer/raytracer"
)

func TestTranslateDoesntChangOriginal(t *testing.T) {
	original := Transform()
	i := m.Identity(4)
	original.Translate(1, 2, 3)
	if !original.Equal(i) {
		t.Errorf("Invalid mutation of matrix %v. Shuld be %v", original, i)
	}

}

func TestTransfromTranslatesPoint(t *testing.T) {
	tr := Transform()
	p := Point(-3, 4, 5)
	result := tr.Translate(5, -3, 2).Apply(p)
	expected := Point(2, 1, 7)
	if !result.Equal(expected) {
		t.Errorf("Failed to translate point %v with %v. Expected %v, got %v.", p, t, expected, result)
	}
}

func TestTransformFromInverse(t *testing.T) {
	tr := Transform().Translate(5, -3, 2).Inverse()
	p := Point(-3, 4, 5)
	result := tr.Apply(p)
	expected := Point(-8, 7, 3)

	if !result.Equal(expected) {
		t.Errorf("Failed apply inverse transform %v to %v. Expected %v, got %v", tr, p, expected, result)
	}
}

func TestTranslationDoesntAffectVectors(t *testing.T) {
	tr := Transform().Translate(5, -3, 2)
	v := Vector(-3, 4, 5)
	result := tr.Apply(v)
	expected := Vector(-3, 4, 5)
	if !result.Equal(expected) {
		t.Errorf("Translation %v should not affect vector %v. Expected %v, got %v", tr, v, expected, result)
	}
}

// func TestScalingPoint(t *testing.T) {
// 	tr := Transform().Scale(2, 3, 4)
// 	p := Point(-4, 5, 8)
// 	result := tr.Apply(p)
// 	expected := Point(-8, 18, 32)
// 	if !result.Equal(expected) {
// 		t.Errorf("Translation %v should not affect vector %v. Expected %v, got %v", tr, v, expected, result)
// 	}
// }
