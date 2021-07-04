package matrix

import "testing"

func TestMatrixCreation(t *testing.T) {
	m := New([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
	})

	if !(m.Rows == 3 && m.Columns == 4) {
		t.Errorf("Failed to initialise matrix. Expected 3x4, got %vx%v", m.Rows, m.Columns)
	}

}

func TestZeroMatrixCreation(t *testing.T) {
	m := Zero(3, 4)

	if !(m.Rows == 3 && m.Columns == 4) {
		t.Errorf("Failed to initialise matrix. Expected 3x4, got %vx%v", m.Rows, m.Columns)
	}

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			v, _ := m.Get(i, j)
			if v != 0.0 {
				t.Errorf("Zero matrix %v has non-zero elements", m)
			}
		}
	}

}

func TestMatrixEquality(t *testing.T) {
	a := New([][]float64{
		{1, 2, 3},
		{3, 2, 1},
	})
	b := New([][]float64{
		{1, 2, 3},
		{3, 2, 1},
	})

	if !(a.Equal(a) && a.Equal(b) && b.Equal(a)) {
		t.Errorf("Equality failed between %v and %v", a, b)
	}
}

func TestMatrixInequality(t *testing.T) {
	a := New([][]float64{
		{1, 2, 3},
		{3, 2, 1},
	})
	b := New([][]float64{
		{3, 2, 1},
		{1, 2, 3},
	})
	c := New([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})

	if a.Equal(b) || a.Equal(c) || b.Equal(c) {
		t.Errorf("different matrices %v, %v, %v are incorrectly equal", a, b, c)
	}
}

func TestMatrixMult2x2(t *testing.T) {
	a := New([][]float64{
		{1, 2},
		{3, 4},
	})
	b := New([][]float64{
		{5, 6},
		{7, 8},
	})
	expected := New([][]float64{
		{1*5 + 2*7, 1*6 + 2*8},
		{3*5 + 4*7, 3*6 + 4*8},
	})
	result, err := a.Mult(b)
	if !result.Equal(expected) || err != nil {
		t.Errorf("%v Mult %v. Expected %v, Got %v (%v)", a, b, expected, result, err)
	}
}

func TestMatrixMult4x4(t *testing.T) {
	a := New([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	b := New([][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})
	expected := New([][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})
	result, err := a.Mult(b)
	if !result.Equal(expected) || err != nil {
		t.Errorf("%v Mult %v. Expected %v, Got %v (%v)", a, b, expected, result, err)
	}
}

func TestMultByIdentity(t *testing.T) {
	m := New([][]float64{
		{1, 2},
		{3, 4},
	})
	i := Identity(2)
	result, err := m.Mult(i)
	if !result.Equal(m) || err != nil {
		t.Errorf("Multiplying %v by identity %v led to incorrect result %v (%v)", m, i, result, err)
	}
	result, err = i.Mult(m)
	if !result.Equal(m) || err != nil {
		t.Errorf("Multiplying identity %v by %v led to incorrect result %v (%v)", i, m, result, err)
	}
}

func TestTranspose4x4(t *testing.T) {
	m := New([][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	})

	result := m.Transpose()
	expected := New([][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})

	if !result.Equal(expected) {
		t.Errorf("Bad transpose of %v. Expected %v, Got %v", m, expected, result)
	}
}

func TestTranspose2x3(t *testing.T) {
	m := New([][]float64{
		{0, 9, 3},
		{9, 8, 0},
	})

	result := m.Transpose()
	expected := New([][]float64{
		{0, 9},
		{9, 8},
		{3, 0},
	})

	if !result.Equal(expected) {
		t.Errorf("Bad transpose of %v. Expected %v, Got %v", m, expected, result)
	}
}

func TestSubmatrix3x3(t *testing.T) {
	m := New([][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})
	result, err := m.Submatrix(0, 2)
	expected := New([][]float64{
		{-3, 2},
		{0, 6},
	})
	if !result.Equal(expected) {
		t.Errorf("Failed to get submatrix 0,2 of %v. Expected %v, got %v (%v)", m, expected, result, err)
	}
}

func TestSubmatrix4x4(t *testing.T) {
	m := New([][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	result, err := m.Submatrix(2, 1)
	expected := New([][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})
	if !result.Equal(expected) {
		t.Errorf("Failed to get submatrix 0,2 of %v. Expected %v, got %v (%v)", m, expected, result, err)
	}
}

func TestDeterminant2x2(t *testing.T) {
	m := New([][]float64{
		{1, 5},
		{-3, 2},
	})
	result, err := m.Determinant()
	expected := 17.0
	if result != expected {
		t.Errorf("Determinant of %v incorrect. Expected %v, got %v (%v)", m, expected, result, err)
	}
}

func TestMinor3x3(t *testing.T) {
	m := New([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	minor, err := m.Minor(1, 0)
	expected := 25.0
	if minor != expected {
		t.Errorf("Minor 1,0 of %v failed. Expected %v, got %v (%v)", m, expected, minor, err)
	}
}

func TestCofactor(t *testing.T) {
	m := New([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	c, err := m.Cofactor(0, 0)
	expected := -12.0
	if c != expected {
		t.Errorf("Failed to calculate cofactor 0,0 for %v. Expected %v, got %v (%v)", m, expected, c, err)
	}

	c, err = m.Cofactor(1, 0)
	expected = -25.0
	if c != expected {
		t.Errorf("Failed to calculate cofactor 1,0 for %v. Expected %v, got %v (%v)", m, expected, c, err)
	}
}

func TestDeterminant3x3(t *testing.T) {
	m := New([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})
	det, err := m.Determinant()
	expected := -196.0
	if det != expected {
		t.Errorf("Failed to calculate determinant of %v. Expected %v, got %v (%v)", m, expected, det, err)
	}
}
func TestDeterminant4x4(t *testing.T) {
	m := New([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})
	det, err := m.Determinant()
	expected := -4071.0
	if det != expected {
		t.Errorf("Failed to calculate determinant of %v. Expected %v, got %v (%v)", m, expected, det, err)
	}
}

func TestInvertible(t *testing.T) {
	invertible := New([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})

	if !invertible.IsInvertible() {
		t.Errorf("Invertible matrix %v labeled as not invertible.", invertible)
	}

	notInvertible := New([][]float64{
		{2, 6},
		{1, 3},
	})
	if notInvertible.IsInvertible() {
		t.Errorf("Not invertible matrix %v labeled as invertible.", notInvertible)
	}
}

func TestInvert2x2(t *testing.T) {
	m := New([][]float64{
		{4, 3},
		{1, 1},
	})

	expected := New([][]float64{
		{1, -3},
		{-1, 4},
	})

	i, err := m.Inverse()

	if !i.Equal(expected) {
		t.Errorf("Failed to invert matrix %v. Expected %v, got %v (%v)", m, expected, i, err)
	}
}
func TestInvert3x3(t *testing.T) {
	m := New([][]float64{
		{3, 0, 2},
		{2, 0, -2},
		{0, 1, 1},
	})

	expected := New([][]float64{
		{.2, .2, 0},
		{-.2, .3, 1},
		{.2, -.3, 0},
	})

	i, err := m.Inverse()

	if !i.Equal(expected) {
		t.Errorf("Failed to invert matrix %v. Expected %v, got %v (%v)", m, expected, i, err)
	}
}

func TestInvert4x4(t *testing.T) {
	m := New([][]float64{
		{1, 1, 1, -1},
		{1, 1, -1, 1},
		{1, -1, 1, 1},
		{-1, 1, 1, 1},
	})

	expected := New([][]float64{
		{.25, .25, .25, -.25},
		{.25, .25, -.25, .25},
		{.25, -.25, .25, .25},
		{-.25, .25, .25, .25},
	})

	i, err := m.Inverse()

	if !i.Equal(expected) {
		t.Errorf("Failed to invert matrix %v. Expected %v, got %v (%v)", m, expected, i, err)
	}
}

func TestMultInverseYieldsIdentity(t *testing.T) {
	m := New([][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})
	inverse, _ := m.Inverse()
	result, _ := m.Mult(inverse)
	if !result.Equal(Identity(4)) {
		t.Errorf("Multiplying %v by its inverse does not yield the identity. Expected %v, got %v", m, Identity(4), result)
	}
}

func TestProductByInverse(t *testing.T) {
	a := New([][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})
	b := New([][]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})
	c, _ := a.Mult(b)
	i, _ := b.Inverse()
	result, _ := c.Mult(i)
	if !result.Equal(a) {
		t.Errorf("Failed to get a sensible value when multiplying product by inverse.")
	}
}
