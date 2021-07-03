package quaternion

import (
	"math"
	"testing"
)

func TestQuaternionAsPoint(t *testing.T) {
	q := Quaternion{0.0, 0.0, 0.0, 1.0}
	if !IsPoint(q) || IsVector(q) {
		t.Errorf("Point vector was not recognised as a point. Expected point but got %v.", q)
	}
}

func TestQuaternionAsVector(t *testing.T) {
	q := Quaternion{0.0, 0.0, 0.0, 0.0}
	if !IsVector(q) || IsPoint(q) {
		t.Errorf("Point vector was not recognised as a point. Expected point but got %v.", q)
	}
}

func TestPointConstructor(t *testing.T) {
	if !IsPoint(Point(0.0, 0.0, 0.0)) {
		t.Errorf("Failed to construct Point")
	}
}

func TestVectorConstructor(t *testing.T) {
	if !IsVector(Vector(0.0, 0.0, 0.0)) {
		t.Errorf("Failed to construct Vector")
	}
}

func TestQuaternionAddition(t *testing.T) {
	A := Quaternion{3, -2, 5, 1}
	B := Quaternion{-2, 3, 1, 0}

	C := Add(A, B)
	expected := Quaternion{1, 1, 6, 1}

	if !Equal(C, expected) {
		t.Errorf("Could not add %v and %v. Got %v expected %v", A, B, C, expected)
	}
}

func TestPointSubstraction(t *testing.T) {
	A := Point(3, 2, 1)
	B := Point(5, 6, 7)
	expected := Vector(-2, -4, -6)

	C := Sub(A, B)

	if !Equal(C, expected) {
		t.Errorf("Could not substract point %v from point %v. Got %v expected %v", B, A, C, expected)
	}

}

func TestSubstractingVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	expected := Point(-2, -4, -6)

	result := Sub(p, v)

	if !Equal(result, expected) {
		t.Errorf("Could not substract vector %v from point %v. Got %v expected %v", v, p, result, expected)
	}

}

func TestVectorSubstraction(t *testing.T) {
	a := Vector(3, 2, 1)
	b := Vector(5, 6, 7)
	expected := Vector(-2, -4, -6)
	result := Sub(a, b)
	if !Equal(result, expected) {
		t.Errorf("Could not substract vector %v from vector %v. Got %v expected %v", b, a, result, expected)
	}

}

func TestQuaternionNegation(t *testing.T) {
	a := Quaternion{1, -2, 3, -4}
	expected := Quaternion{-1, 2, -3, 4}
	result := a.Negate()

	if !Equal(result, expected) {
		t.Errorf("Failed to negate %v. Got %v, expected %v", a, result, expected)
	}
}

func TestQuaternionScale(t *testing.T) {
	a := Quaternion{1, -2, 3, -4}
	scalar := 3.5
	result := a.Scale(scalar)
	expected := Quaternion{3.5, -7, 10.5, -14}

	if !Equal(result, expected) {
		t.Errorf("Failed to scale %v by %v. Got %v, expected %v", a, scalar, result, expected)
	}
}

func TestQuaternionDiv(t *testing.T) {
	a := Quaternion{1, -2, 3, -4}
	scalar := 2.0
	result := a.Divide(scalar)
	expected := Quaternion{0.5, -1, 1.5, -2}

	if !Equal(result, expected) {
		t.Errorf("Failed to divide %v by %v. Got %v, expected %v", a, scalar, result, expected)
	}
}

func TestMagnitude(t *testing.T) {
	tests := []Quaternion{
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
	tests := []Quaternion{
		Vector(4, 0, 0),
		Vector(1, 2, 3),
	}
	results := []Quaternion{
		Vector(1, 0, 0),
		Vector(1, 2, 3).Divide(math.Sqrt(14)),
	}
	for i, v := range tests {
		result := v.Normalize()
		expected := results[i]
		if result != expected {
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
