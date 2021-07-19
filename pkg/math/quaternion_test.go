package math_test

import (
	"fmt"
	"testing"

	m "math"

	"github.com/bricef/ray-tracer/pkg/math"
)

func TestQuaternionAsPoint(t *testing.T) {
	a := math.NewQuaternion(0.0, 0.0, 0.0, 1.0)
	if !a.IsPoint() || a.IsVector() {
		t.Errorf("Point vector was not recognised as a point. Expected point but got %v.", a)
	}
}

func TestQuaternionAsVector(t *testing.T) {
	a := math.NewQuaternion(0.0, 0.0, 0.0, 0.0)
	if !a.IsVector() || a.IsPoint() {
		t.Errorf("Point vector was not recognised as a point. Expected point but got %v.", a)
	}
}

func TestQuaternionAddition(t *testing.T) {
	A := math.NewQuaternion(3, -2, 5, 1)
	B := math.NewQuaternion(-2, 3, 1, 0)

	C := A.Add(B)
	expected := math.NewQuaternion(1, 1, 6, 1)

	if !C.Equal(expected) {
		t.Errorf("Could not add %v and %v. Got %v expected %v", A, B, C, expected)
	}
}

func TestPointSubstraction(t *testing.T) {
	A := math.NewPoint(3, 2, 1)
	B := math.NewPoint(5, 6, 7)
	expected := math.NewVector(-2, -4, -6)

	C := A.Sub(B)

	if !C.Equal(expected) {
		t.Errorf("Could not substract point %v from point %v. Got %v expected %v", B, A, C, expected)
	}

}

func TestSubstractingVectorFromPoint(t *testing.T) {
	p := math.NewPoint(3, 2, 1)
	v := math.NewVector(5, 6, 7)
	expected := math.NewPoint(-2, -4, -6)

	result := p.Sub(v)

	if !result.Equal(expected) {
		t.Errorf("Could not substract vector %v from point %v. Got %v expected %v", v, p, result, expected)
	}

}

func TestVectorSubstraction(t *testing.T) {
	a := math.NewVector(3, 2, 1)
	b := math.NewVector(5, 6, 7)
	expected := math.NewVector(-2, -4, -6)
	result := a.Sub(b)
	if !result.Equal(expected) {
		t.Errorf("Could not substract vector %v from vector %v. Got %v expected %v", b, a, result, expected)
	}

}

func TestQuaternionNegation(t *testing.T) {
	a := math.NewQuaternion(1, -2, 3, -4)
	expected := math.NewQuaternion(-1, 2, -3, 4)
	result := a.Negate()

	if !result.Equal(expected) {
		t.Errorf("Failed to negate %v. Got %v, expected %v", a, result, expected)
	}
}

func TestQuaternionScale(t *testing.T) {
	a := math.NewQuaternion(1, -2, 3, -4)
	scalar := 3.5
	result := a.Scale(scalar)
	expected := math.NewQuaternion(3.5, -7, 10.5, -14)

	if !result.Equal(expected) {
		t.Errorf("Failed to scale %v by %v. Got %v, expected %v", a, scalar, result, expected)
	}
}

func TestQuaternionDiv(t *testing.T) {
	a := math.NewQuaternion(1, -2, 3, -4)
	scalar := 2.0
	result := a.Divide(scalar)
	expected := math.NewQuaternion(0.5, -1, 1.5, -2)

	if !result.Equal(expected) {
		t.Errorf("Failed to divide %v by %v. Got %v, expected %v", a, scalar, result, expected)
	}
}

func TestMagnitude(t *testing.T) {
	tests := []math.Vector{
		math.NewVector(1, 0, 0),
		math.NewVector(0, 1, 0),
		math.NewVector(0, 0, 1),
		math.NewVector(1, 2, 3),
		math.NewVector(-1, -2, -3),
	}
	results := []float64{
		1.0,
		1.0,
		1.0,
		m.Sqrt(14),
		m.Sqrt(14),
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
	tests := []math.Vector{
		math.NewVector(4, 0, 0),
		math.NewVector(1, 2, 3),
	}
	results := []math.Vector{
		math.NewVector(1, 0, 0),
		math.NewVector(1, 2, 3).Divide(m.Sqrt(14)).AsVector(),
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
	a := math.NewVector(1, 2, 3)
	b := math.NewVector(2, 3, 4)
	result := a.Dot(b)
	expected := 20.0
	if result != expected {
		t.Errorf("Failed to calculate dot product %v dot %v. Got %v, expected %v", a, b, result, expected)
	}

}

func TestCrossProduct(t *testing.T) {
	a := math.NewVector(1, 2, 3)
	b := math.NewVector(2, 3, 4)

	result1 := a.Cross(b)
	expected1 := math.NewVector(-1, 2, -1)

	result2 := b.Cross(a)
	expected2 := math.NewVector(1, -2, 1)

	if result1 != expected1 {
		t.Errorf("Failed to calculate %v x %v. Expected %v, got %v", a, b, result1, expected1)
	}

	if result2 != expected2 {
		t.Errorf("Failed to calculate %v x %v. Expected %v, got %v", b, a, result2, expected2)
	}

}

func TestVectorReflection1(t *testing.T) {
	v := math.NewVector(1, -1, 0)
	normal := math.NewVector(0, 1, 0)
	result := v.Reflect(normal)
	expected := math.NewVector(1, 1, 0)
	if !result.Equal(expected) {
		t.Errorf("Reflection failed. Reflecting %v in %v. Expected %v, got %v", v, normal, expected, result)
	}
}

func TestVectorReflection2(t *testing.T) {
	v := math.NewVector(0, -1, 0)
	normal := math.NewVector(m.Sqrt2/2, m.Sqrt2/2, 0)
	result := v.Reflect(normal)
	expected := math.NewVector(1, 0, 0)
	if !result.Equal(expected) {
		t.Errorf("Reflection failed. Reflecting %v in %v. Expected %v, got %v", v, normal, expected, result)
	}
}

func TestMagnitudeSanity(t *testing.T) {
	position := math.NewPoint(0, 0, 0)
	point := math.NewPoint(0, 0, -10)
	v := point.Sub(position).AsVector()
	mag_v := v.Magnitude()
	normal_v := v.Normalize()

	if mag_v != 10 {
		t.Errorf("Magnitude failure. Expected 10, got %v", mag_v)
	}
	if !normal_v.Equal(math.NewVector(0, 0, -1)) {
		t.Errorf("normalisation failure. Expected %v, got %v", math.NewVector(0, 0, -1), normal_v)
	}

	fmt.Printf("mag: %v, normal:%v", mag_v, normal_v)

}
