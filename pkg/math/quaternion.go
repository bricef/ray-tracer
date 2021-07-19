package math

import (
	"fmt"
	"math"

	"github.com/bricef/ray-tracer/pkg/utils"
)

type quaternion struct {
	x float64
	y float64
	z float64
	w float64
}

func (q quaternion) X() float64 {
	return q.x
}

func (q quaternion) Y() float64 {
	return q.y
}

func (q quaternion) Z() float64 {
	return q.z
}
func (q quaternion) W() float64 {
	return q.w
}

func NewQuaternion(x, y, z, w float64) Quaternion {
	return quaternion{x, y, z, w}
}

func (a quaternion) IsVector() bool {
	return a.w == 0.0
}

func (a quaternion) IsPoint() bool {
	return a.w == 1.0
}

func (a quaternion) AsVector() Vector {
	return vector{a}
}

func (a quaternion) AsPoint() Point {
	return point{a}
}

func (a quaternion) Add(q Quaternion) Quaternion {
	return NewQuaternion(a.x+q.X(), a.y+q.Y(), a.z+q.Z(), a.w+q.W())
}

func (a quaternion) Equal(q Quaternion) bool {
	return utils.AlmostEqual(a.x, q.X()) && utils.AlmostEqual(a.y, q.Y()) && utils.AlmostEqual(a.z, q.Z()) && utils.AlmostEqual(a.w, q.W())
}

func (a quaternion) Sub(q Quaternion) Quaternion {
	return NewQuaternion(a.x-q.X(), a.y-q.Y(), a.z-q.Z(), a.w-q.W())
}

func (a quaternion) Scale(s float64) Quaternion {
	return quaternion{a.x * s, a.y * s, a.z * s, a.w * s}
}

func (q quaternion) Negate() Quaternion {
	return quaternion{-q.x, -q.y, -q.z, -q.w}
}

func (q quaternion) Divide(s float64) Quaternion {
	return NewQuaternion(q.X()/s, q.Y()/s, q.Z()/s, q.W()/s)
}

func (q quaternion) String() string {
	var t string = "Quaternion"
	if q.W() == 1.0 {
		t = "Point"
	} else if q.W() == 0.0 {
		t = "Vector"
	}
	return fmt.Sprintf("%s(%v,%v,%v)", t, q.X(), q.Y(), q.Z())
}

// Point
type point struct{ quaternion }

func NewPoint(x, y, z float64) Point {
	return point{quaternion{x, y, z, 1.0}}
}

// Vector
type vector struct{ quaternion }

func NewVector(x, y, z float64) Vector {
	return vector{quaternion{x, y, z, 0.0}}
}

func (q vector) Magnitude() float64 {
	return math.Sqrt(q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w)
}

func (q vector) Normalize() Vector {
	return q.Divide(q.Magnitude()).AsVector()
}

func (v vector) Dot(o Vector) float64 {
	return v.X()*o.X() + v.Y()*o.Y() + v.Z()*o.Z() + v.W()*o.W()
}

func (v vector) Cross(o Vector) Vector {
	return vector{quaternion{
		v.y*o.Z() - v.z*o.Y(),
		v.z*o.X() - v.x*o.Z(),
		v.x*o.Y() - v.y*o.X(),
		0.0,
	}}
}

func (q vector) Invert() Vector {
	return q.Negate().AsVector()
}

func (q vector) Reflect(n Vector) Vector {
	return q.Sub(n.Scale(2 * q.Dot(n))).AsVector()
}
