package color

import "testing"

func TestColorCreation(t *testing.T) {
	c := New(-0.5, 0.4, 1.7)

	if !(c.R == -0.5 && c.G == 0.4 && c.B == 1.7) {
		t.Errorf("Failed to create a color.")
	}
}

func TestColorAdd(t *testing.T) {
	a := New(0.9, 0.6, 0.75)
	b := New(0.7, 0.1, 0.25)
	result := a.Add(b)
	expected := New(1.6, 0.7, 1.0)

	if !result.Equal(expected) {
		t.Errorf("Failed to add color %v and %v. Expected %v, got %v.", a, b, expected, result)
	}
}

func TestColorSubstract(t *testing.T) {
	a := New(0.9, 0.6, 0.75)
	b := New(0.7, 0.1, 0.25)
	result := a.Sub(b)
	expected := New(0.2, 0.5, 0.5)

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

func TestalternativeConstructors(t *testing.T) {
	type cases struct {
		result   Color
		expected Color
	}

	tests := []cases{
		{Hex(0xff0000), New(1, 0, 0)},
		{Hex(0x00ff00), New(0, 1, 0)},
		{Hex(0x0000ff), New(0, 0, 1)},
		{Hex(0xff0000), New(1, 0, 0)},
		{Hex(0x5F9EA0), New(0.37, 0.62, 0.63)},
		{Bytes(255, 0, 0), New(1, 0, 0)},
		{Bytes(0, 255, 0), New(0, 1, 0)},
		{Bytes(0, 0, 255), New(0, 0, 1)},
		{Bytes(95, 158, 160), New(0.37, 0.62, 0.63)},
	}
	for _, c := range tests {
		if !c.result.EqualToTolerance(c.expected, 0.005) {
			t.Errorf("Hex color creation failed. Expected %v, got %v", c.expected, c.result)
		}
	}
}
