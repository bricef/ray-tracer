package color

import "testing"

func TestColorCreation(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)

	if !(c.R == -0.5 && c.G == 0.4 && c.B == 1.7) {
		t.Errorf("Failed to create a color.")
	}
}

func TestColorAdd(t *testing.T) {
	a := NewColor(0.9, 0.6, 0.75)
	b := NewColor(0.7, 0.1, 0.25)
	result := a.Add(b)
	expected := NewColor(1.6, 0.7, 1.0)

	if !result.Equal(expected) {
		t.Errorf("Failed to add color %v and %v. Expected %v, got %v.", a, b, expected, result)
	}
}

func TestColorSubstract(t *testing.T) {
	a := NewColor(0.9, 0.6, 0.75)
	b := NewColor(0.7, 0.1, 0.25)
	result := a.Sub(b)
	expected := NewColor(0.2, 0.5, 0.5)

	if !result.Equal(expected) {
		t.Errorf("Failed to substract color %v from %v. Expected %v, got %v.", b, a, expected, result)
	}
}

func TestColormultiply(t *testing.T) {
	a := Color{1, 0.2, 0.4}
	b := Color{0.9, 1, 0.1}
	result := a.Mult(b)
	expected := Color{0.9, 0.2, 0.04}

	if !result.Equal(expected) {
		t.Errorf("Failed to multiply color %v and %v. Expected %v, got %v.", a, b, expected, result)
	}
}
