package quaternion

import (
	"errors"
	"math"
)

type Quaternion struct {
	x float64
	y float64
	z float64
	w float64
}

func IsPoint(q Quaternion) bool {
	return q.w == 1.0
}

func IsVector(q Quaternion) bool {
	return q.w == 0.0
}

func Point(x float64, y float64, z float64) Quaternion {
	return Quaternion{x, y, z, 1.0}
}

func Vector(x float64, y float64, z float64) Quaternion {
	return Quaternion{x, y, z, 0.0}
}

func Add(a Quaternion, b Quaternion) Quaternion {
	return Quaternion{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func Equal(a Quaternion, b Quaternion) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z && a.w == b.w
}

func Sub(a Quaternion, b Quaternion) Quaternion {
	return Quaternion{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w}
}

func (q Quaternion) Negate() Quaternion {
	return Quaternion{-q.x, -q.y, -q.z, -q.w}
}

func (q Quaternion) Scale(s float64) Quaternion {
	return Quaternion{s * q.x, s * q.y, s * q.z, s * q.w}
}

func (q Quaternion) Divide(s float64) Quaternion {
	return Quaternion{q.x / s, q.y / s, q.z / s, q.w / s}
}

func (q Quaternion) Magnitude() float64 {
	return math.Sqrt(q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w)
}

func (q Quaternion) Normalize() Quaternion {
	return q.Divide(q.Magnitude())
}

func (q Quaternion) Dot(o Quaternion) float64 {
	return q.x*o.x + q.y*o.y + q.z*o.z + q.w*o.w
}

func (q Quaternion) Cross(o Quaternion) (Quaternion, error) {
	if !(IsVector(q) && IsVector(o)) {
		return Quaternion{}, errors.New("typeError: Cannot take the cross product of non-vectors")
	}
	return Vector(
		q.y*o.z-q.z*o.y,
		q.z*o.x-q.x*o.z,
		q.x*o.y-q.y*o.x,
	), nil
}
