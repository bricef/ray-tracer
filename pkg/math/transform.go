package math

import (
	"math"
)

type MatrixTransform struct {
	Matrix
}

func NewTransform() *MatrixTransform {
	return &MatrixTransform{Identity(4)}
}

func (a MatrixTransform) GetMatrix() Matrix {
	return a.Matrix
}

func (a MatrixTransform) Equal(b Transform) bool {
	return a.Matrix.Equal(b.GetMatrix())
}

func (t MatrixTransform) Raw(raw [][]float64) Transform {
	m, _ := t.Matrix.Mult(NewMatrix(raw))
	t.Matrix = m
	return t
}

func (t MatrixTransform) Translate(x, y, z float64) Transform {
	return t.Raw([][]float64{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	})
}

func (t MatrixTransform) Apply(q Quaternion) Quaternion {
	m := MatrixFromQuaternion(q) // Quaternion to matrix
	result, _ := t.Matrix.Mult(m)

	x, _ := result.Get(0, 0)
	y, _ := result.Get(1, 0)
	z, _ := result.Get(2, 0)
	w, _ := result.Get(3, 0)

	return NewQuaternion(x, y, z, w)
}

func (t MatrixTransform) Inverse() Transform {
	m, _ := t.Matrix.Inverse()
	t.Matrix = m
	return t
}

func (t MatrixTransform) Transpose() Transform {
	t.Matrix = t.Matrix.Transpose()
	return t
}

func (t MatrixTransform) Scale(x, y, z float64) Transform {
	return t.Raw([][]float64{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	})
}

func (t MatrixTransform) ReflectX() Transform {
	return t.Scale(-1, 1, 1)
}
func (t MatrixTransform) ReflectY() Transform {
	return t.Scale(1, -1, 1)
}
func (t MatrixTransform) ReflectZ() Transform {
	return t.Scale(1, 1, -1)
}

func (t MatrixTransform) RotateX(r float64) Transform {
	return t.Raw([][]float64{
		{1, 0, 0, 0},
		{0, math.Cos(r), -math.Sin(r), 0},
		{0, math.Sin(r), math.Cos(r), 0},
		{0, 0, 0, 1},
	})
}

func (t MatrixTransform) RotateY(r float64) Transform {
	return t.Raw([][]float64{
		{math.Cos(r), 0, math.Sin(r), 0},
		{0, 1, 0, 0},
		{-math.Sin(r), 0, math.Cos(r), 0},
		{0, 0, 0, 1},
	})
}

func (t MatrixTransform) RotateZ(r float64) Transform {
	return t.Raw([][]float64{
		{math.Cos(r), -math.Sin(r), 0, 0},
		{math.Sin(r), math.Cos(r), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

func (t MatrixTransform) Shear(xy, xz, yx, yz, zx, zy float64) Transform {
	return t.Raw([][]float64{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	})
}

func ViewTransform(from Point, to Point, up Vector) Transform {
	forward := to.Sub(from).AsVector().Normalize()
	left := forward.Cross(up.Normalize())
	trueUp := left.Cross(forward)

	orientation := MatrixTransform{
		NewMatrix(
			[][]float64{
				{left.X(), left.Y(), left.Z(), 0},
				{trueUp.X(), trueUp.Y(), trueUp.Z(), 0},
				{-forward.X(), -forward.Y(), -forward.Z(), 0},
				{0, 0, 0, 1},
			},
		),
	}
	return orientation.Translate(-from.X(), -from.Y(), -from.Z())
}

func (t MatrixTransform) MoveTo(p Point) Transform {
	t.Set(0, 3, p.X())
	t.Set(1, 3, p.Y())
	t.Set(2, 3, p.Z())
	return t
}

func (t MatrixTransform) Position() Point {
	return NewPoint(t.Matrix.Values[0][3], t.Matrix.Values[1][3], t.Matrix.Values[2][3])
}

func Translate(x, y, z float64) Transform {
	return NewTransform().Translate(x, y, z)
}
func Scale(x, y, z float64) Transform {
	return NewTransform().Scale(x, y, z)
}
func Shear(xy, xz, yx, yz, zx, zy float64) Transform {
	return NewTransform().Shear(xy, xz, yx, yz, zx, zy)
}
func RotateX(r float64) Transform {
	return NewTransform().RotateX(r)
}
func RotateY(r float64) Transform {
	return NewTransform().RotateY(r)
}
func RotateZ(r float64) Transform {
	return NewTransform().RotateZ(r)
}
func ReflectX() Transform {
	return NewTransform().ReflectX()
}
func ReflectY() Transform {
	return NewTransform().ReflectY()
}
func ReflectZ() Transform {
	return NewTransform().ReflectZ()
}
