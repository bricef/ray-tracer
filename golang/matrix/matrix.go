package matrix

import (
	"fmt"
	"math"
)

type Matrix struct {
	Rows    int
	Columns int
	Values  [][]float64
}

func New(values [][]float64) Matrix {
	return Matrix{
		len(values),
		len(values[0]),
		values,
	}
}

func Zero(r, c int) Matrix {
	values := make([][]float64, r)
	for i := range values {
		values[i] = make([]float64, c)
	}
	return Matrix{
		r,
		c,
		values,
	}
}

func (m Matrix) Get(r, c int) (float64, error) {
	if r >= m.Rows || c >= m.Columns {
		return 0.0, fmt.Errorf("out of bounds error. Cannot access %v,%v on %v", r, c, m)
	}
	return m.Values[r][c], nil
}

func (m Matrix) Equal(o Matrix) bool {
	if m.Rows != o.Rows || m.Columns != o.Columns {
		return false
	}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			mv, _ := m.Get(i, j)
			ov, _ := o.Get(i, j)
			if mv != ov {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Mult(o Matrix) (Matrix, error) {
	if m.Columns != o.Rows {
		return Matrix{}, fmt.Errorf("cannot multiply %v by %v", m, o)
	}
	out := Zero(m.Rows, o.Columns)

	// See https://en.wikipedia.org/wiki/Matrix_multiplication_algorithm
	for i := 0; i < out.Rows; i++ {
		for j := 0; j < out.Columns; j++ {
			sum := 0.0
			for k := 0; k < m.Rows; k++ {
				Aij, err := m.Get(i, k)
				if err != nil {
					return Matrix{}, err
				}
				Bkj, err := o.Get(k, j)
				if err != nil {
					return Matrix{}, err
				}
				sum += Aij * Bkj
			}
			out.Values[i][j] = sum
		}
	}
	return out, nil
}

func Identity(n int) Matrix {
	m := Zero(n, n)

	for i := 0; i < n; i++ {
		m.Values[i][i] = 1.0
	}
	return m
}

func (m Matrix) Transpose() Matrix {
	t := Zero(m.Columns, m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			t.Values[j][i] = m.Values[i][j]
		}
	}
	return t
}

func (m Matrix) Submatrix(r, c int) (Matrix, error) {
	if !(c < m.Columns && r < m.Rows) {
		return Matrix{}, fmt.Errorf("out of bounds error. Cannot access %v,%v on %v", r, c, m)
	}
	s := Zero(m.Rows-1, m.Columns-1)

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			if i != r && j != c {
				k := i
				l := j
				if i >= r {
					k = k - 1
				}
				if j >= c {
					l = l - 1
				}
				s.Values[k][l] = m.Values[i][j]
			}
		}
	}
	return s, nil

}

func (m Matrix) Determinant() (float64, error) {
	if m.Rows != m.Columns {
		return 0.0, fmt.Errorf("cannot take the determinant of non-square matrix %v", m)
	}

	if m.Rows == 2 && m.Columns == 2 {
		return m.Values[0][0]*m.Values[1][1] - m.Values[0][1]*m.Values[1][0], nil
	}

	acc := 0.0
	for j := 0; j < m.Columns; j++ {
		c, err := m.Cofactor(0, j)
		if err != nil {
			return 0, err
		}
		acc += m.Values[0][j] * c
	}

	return acc, nil

}

func (m Matrix) Minor(r, c int) (float64, error) {
	s, err := m.Submatrix(r, c)
	if err != nil {
		return 0.0, err
	}
	det, err := s.Determinant()
	if err != nil {
		return 0.0, err
	}
	return det, nil
}

func (m Matrix) Cofactor(r, c int) (float64, error) {
	minor, err := m.Minor(r, c)
	if err != nil {
		return 0, err
	}
	// Fuck you Go for not implementing basic integer exponentiation. What a piece of shit.
	return math.Pow(-1, float64(r+c)) * minor, nil

}

func (m Matrix) IsInvertible() bool {
	if m.Rows != m.Columns {
		return false
	}
	det, err := m.Determinant()
	if err != nil {
		return false
	}
	return det != 0.0

}
