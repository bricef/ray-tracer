package transform

import (
	"github.com/bricef/ray-tracer/matrix"
	"github.com/bricef/ray-tracer/quaternion"
)

type Transform struct {
	matrix.Matrix
}

func New() Transform {
	return Transform{matrix.Identity(4)}
}

func (t Transform) Translate(x, y, z float64) Transform {
	m, _ := t.Matrix.Mult(matrix.New([][]float64{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}))
	return Transform{m}
}

func (t Transform) Apply(a interface{}) quaternion.Quaternion {
	q := quaternion.From(a)       // Accepts Quaternions, Vectors, Points
	m := matrix.FromQuaternion(q) // Quaternion to matrix
	result, _ := t.Matrix.Mult(m)

	x, _ := result.Get(0, 0)
	y, _ := result.Get(1, 0)
	z, _ := result.Get(2, 0)
	w, _ := result.Get(3, 0)

	return quaternion.New(x, y, z, w)
}

func (t Transform) Inverse() Transform {
	m, _ := t.Matrix.Inverse()
	return Transform{m}
}

func (t Transform) Scale(x, y, z float64) Transform {
	m, _ := t.Matrix.Mult(matrix.New([][]float64{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}))
	return Transform{m}
}

func (t Transform) ReflectX() Transform {
	return t.Scale(-1, 1, 1)
}
func (t Transform) ReflectY() Transform {
	return t.Scale(1, -1, 1)
}
func (t Transform) ReflectZ() Transform {
	return t.Scale(1, 1, -1)
}
