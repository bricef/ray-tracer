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
	new := t.Clone()
	new.Set(0, 3, x)
	new.Set(1, 3, y)
	new.Set(2, 3, z)
	return Transform{new}
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
