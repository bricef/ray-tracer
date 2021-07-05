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
	q := quaternion.From(a)
	return q
}
