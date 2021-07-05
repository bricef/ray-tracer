package transform

import (
	"testing"

	m "github.com/bricef/ray-tracer/matrix"
	// . "github.com/bricef/ray-tracer/raytracer"
)

func TestTranslateDoesntChangOriginal(t *testing.T) {
	original := New()
	i := m.Identity(4)
	original.Translate(1, 2, 3)
	if !original.Equal(i) {
		t.Errorf("Invalid mutation of matrix %v. Shuld be %v", original, i)
	}

}

// func TestTransfromTranslatesPoint(t *testing.T) {
// 	tr := New()
// 	p := Point(-3, 4, 5)
// 	result := tr.Translate(5, -3, 2).Apply(p)
// 	expected := Point(2, 1, 7)
// 	if !result.Equal(expected) {
// 		t.Errorf("Failed to translate point %v with %v. Expected %v, got %v.", p, t, expected, result)
// 	}

// }
