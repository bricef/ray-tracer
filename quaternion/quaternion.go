package quaternion

import (
	"fmt"
	"math"

	"github.com/bricef/ray-tracer/utils"
)

type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

func New(x, y, z, w float64) Quaternion {
	return Quaternion{x, y, z, w}
}

func IsPoint(q Quaternion) bool {
	return q.W == 1.0
}

func IsVector(q Quaternion) bool {
	return q.W == 0.0
}

func (a Quaternion) Add(q Quaternion) Quaternion {
	return New(a.X+q.X, a.Y+q.Y, a.Z+q.Z, a.W+q.W)
}

func (a Quaternion) Equal(q Quaternion) bool {
	return utils.AlmostEqual(a.X, q.X) && utils.AlmostEqual(a.Y, q.Y) && utils.AlmostEqual(a.Z, q.Z) && utils.AlmostEqual(a.W, q.W)
}

func (a Quaternion) Sub(q Quaternion) Quaternion {
	return New(a.X-q.X, a.Y-q.Y, a.Z-q.Z, a.W-q.W)
}

func (q Quaternion) Negate() Quaternion {
	return New(-q.X, -q.Y, -q.Z, -q.W)
}

func (q Quaternion) Scale(s float64) Quaternion {
	return New(s*q.X, s*q.Y, s*q.Z, s*q.W)
}

func (q Quaternion) Divide(s float64) Quaternion {
	return New(q.X/s, q.Y/s, q.Z/s, q.W/s)
}

func (q Quaternion) Magnitude() float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

func (q Quaternion) Normalize() Quaternion {
	return q.Divide(q.Magnitude())
}

func (v Quaternion) Dot(o Quaternion) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z + v.W*o.W
}

func (v Quaternion) Cross(o Quaternion) Quaternion {
	return New(
		v.Y*o.Z-v.Z*o.Y,
		v.Z*o.X-v.X*o.Z,
		v.X*o.Y-v.Y*o.X,
		0.0,
	)
}

func NewPoint(x, y, z float64) Quaternion {
	return New(x, y, z, 1.0)
}

func NewVector(x, y, z float64) Quaternion {
	return New(x, y, z, 0.0)
}

func (q Quaternion) String() string {
	var t string = "Quaternion"
	if IsPoint(q) {
		t = "Point"
	} else if IsVector(q) {
		t = "Vector"
	}

	return fmt.Sprintf("%s(%v,%v,%v)", t, q.X, q.Y, q.Z)
}

func (q Quaternion) Reflect(n Quaternion) Quaternion {
	return q.Sub(n.Scale(2 * q.Dot(n)))
}

func (q Quaternion) Invert() Quaternion {
	return q.Scale(-1)
}
