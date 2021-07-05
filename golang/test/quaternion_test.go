package test

import (
	"math"
	"testing"

	q "github.com/bricef/ray-tracer/quaternion"
	. "github.com/bricef/ray-tracer/raytracer"
)

func TestQuaternionAsPoint(t *testing.T) {
	a := Quaternion(0.0, 0.0, 0.0, 1.0)
	if !q.IsPoint(a) || q.IsVector(a) {
		t.Errorf("Point vector was not recognised as a point. Expected point but got %v.", a)
	}
}

func TestQuaternionAsVector(t *testing.T) {
	a := Quaternion(0.0, 0.0, 0.0, 0.0)
	if !q.IsVector(a) || q.IsPoint(a) {
		t.Errorf("Point vector was not recognised as a point. Expected point but got %v.", a)
	}
}

func TestQuaternionAddition(t *testing.T) {
	A := Quaternion(3, -2, 5, 1)
	B := Quaternion(-2, 3, 1, 0)

	C := A.Add(B)
	expected := Quaternion(1, 1, 6, 1)

	if !C.Equal(expected) {
		t.Errorf("Could not add %v and %v. Got %v expected %v", A, B, C, expected)
	}
}

func TestPointSubstraction(t *testing.T) {
	A := Point(3, 2, 1)
	B := Point(5, 6, 7)
	expected := Vector(-2, -4, -6)

	C := A.Sub(B)

	if !C.Equal(expected) {
		t.Errorf("Could not substract point %v from point %v. Got %v expected %v", B, A, C, expected)
	}

}

func TestSubstractingVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	expected := Point(-2, -4, -6)

	result := p.Sub(v)

	if !result.Equal(expected) {
		t.Errorf("Could not substract vector %v from point %v. Got %v expected %v", v, p, result, expected)
	}

}

func TestVectorSubstraction(t *testing.T) {
	a := Vector(3, 2, 1)
	b := Vector(5, 6, 7)
	expected := Vector(-2, -4, -6)
	result := a.Sub(b)
	if !result.Equal(expected) {
		t.Errorf("Could not substract vector %v from vector %v. Got %v expected %v", b, a, result, expected)
	}

}

func TestQuaternionNegation(t *testing.T) {
	a := Quaternion(1, -2, 3, -4)
	expected := Quaternion(-1, 2, -3, 4)
	result := a.Negate()

	if !result.Equal(expected) {
		t.Errorf("Failed to negate %v. Got %v, expected %v", a, result, expected)
	}
}

func TestQuaternionScale(t *testing.T) {
	a := Quaternion(1, -2, 3, -4)
	scalar := 3.5
	result := a.Scale(scalar)
	expected := Quaternion(3.5, -7, 10.5, -14)

	if !result.Equal(expected) {
		t.Errorf("Failed to scale %v by %v. Got %v, expected %v", a, scalar, result, expected)
	}
}

func TestQuaternionDiv(t *testing.T) {
	a := Quaternion(1, -2, 3, -4)
	scalar := 2.0
	result := a.Divide(scalar)
	expected := Quaternion(0.5, -1, 1.5, -2)

	if !result.Equal(expected) {
		t.Errorf("Failed to divide %v by %v. Got %v, expected %v", a, scalar, result, expected)
	}
}

func TestMagnitude(t *testing.T) {
	tests := []q.Vector{
		Vector(1, 0, 0),
		Vector(0, 1, 0),
		Vector(0, 0, 1),
		Vector(1, 2, 3),
		Vector(-1, -2, -3),
	}
	results := []float64{
		1.0,
		1.0,
		1.0,
		math.Sqrt(14),
		math.Sqrt(14),
	}

	for i, v := range tests {
		result := v.Magnitude()
		expected := results[i]
		if result != expected {
			t.Errorf("Failed to get magnitude of %v. Got %v, expected %v", v, result, expected)
		}
	}

}

func TestNormalisation(t *testing.T) {
	tests := []q.Vector{
		Vector(4, 0, 0),
		Vector(1, 2, 3),
	}
	results := []q.Vector{
		Vector(1, 0, 0),
		q.Vector{Vector(1, 2, 3).Divide(math.Sqrt(14))},
	}
	for i, v := range tests {
		result := v.Normalize()
		expected := results[i]
		if !result.Equal(expected) {
			t.Errorf("Failed to normalize %v. Got %v, expected %v", v, result, expected)
		}
	}
}

func TestDotProduct(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	result := a.Dot(b)
	expected := 20.0
	if result != expected {
		t.Errorf("Failed to calculate dot product %v dot %v. Got %v, expected %v", a, b, result, expected)
	}

}

func TestCrossProduct(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)

	result1, err1 := a.Cross(b)
	expected1 := Vector(-1, 2, -1)

	result2, err2 := b.Cross(a)
	expected2 := Vector(1, -2, 1)

	if (result1 != expected1) || err1 != nil {
		t.Errorf("Failed to calculate %v x %v. Expected %v, got %v. (%v)", a, b, result1, expected1, err1)
	}

	if (result2 != expected2) || err2 != nil {
		t.Errorf("Failed to calculate %v x %v. Expected %v, got %v. (%v)", b, a, result2, expected2, err2)
	}

}
