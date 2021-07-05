package quaternion

import (
	"math"
)

type Quaternion struct {
	x float64
	y float64
	z float64
	w float64
}

func New(x, y, z, w float64) Quaternion {
	return Quaternion{x, y, z, w}
}

func IsPoint(q Quaternion) bool {
	return q.w == 1.0
}

func IsVector(q Quaternion) bool {
	return q.w == 0.0
}

func From(b interface{}) Quaternion {
	q := Quaternion{}
	switch b := b.(type) {
	case Point:
		q = b.Quaternion
	case Vector:
		q = b.Quaternion
	case Quaternion:
		q = b
	default:
		return Quaternion{}
	}
	return q
}

func (a Quaternion) Add(b interface{}) Quaternion {
	q := From(b)
	return New(a.x+q.x, a.y+q.y, a.z+q.z, a.w+q.w)
}

func (a Quaternion) Equal(b interface{}) bool {
	q := From(b)
	return a.x == q.x && a.y == q.y && a.z == q.z && a.w == q.w
}

func (a Quaternion) Sub(b interface{}) Quaternion {
	q := From(b)
	return New(a.x-q.x, a.y-q.y, a.z-q.z, a.w-q.w)
}

func (q Quaternion) Negate() Quaternion {
	return New(-q.x, -q.y, -q.z, -q.w)
}

func (q Quaternion) Scale(s float64) Quaternion {
	return New(s*q.x, s*q.y, s*q.z, s*q.w)
}

func (q Quaternion) Divide(s float64) Quaternion {
	return New(q.x/s, q.y/s, q.z/s, q.w/s)
}

/*
 * Points
 */

type Point struct {
	Quaternion
}

func NewPoint(x, y, z float64) Point {
	return Point{Quaternion{x, y, z, 1.0}}
}

func (p Point) Sub(o interface{}) Quaternion {
	q := From(o)
	return p.Quaternion.Sub(q)
}

/*
 * Vectors
 */
type Vector struct {
	Quaternion
}

func NewVector(x, y, z float64) Vector {
	return Vector{Quaternion{x, y, z, 0.0}}
}

func (q Vector) Magnitude() float64 {
	return math.Sqrt(q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w)
}

func (q Vector) Normalize() Vector {
	return Vector{q.Divide(q.Magnitude())}
}

func (v Vector) Dot(o Vector) float64 {
	return v.x*o.x + v.y*o.y + v.z*o.z + v.w*o.w
}

func (v Vector) Cross(o Vector) (Vector, error) {
	return Vector{New(
		v.y*o.z-v.z*o.y,
		v.z*o.x-v.x*o.z,
		v.x*o.y-v.y*o.x,
		0.0,
	)}, nil
}

func (v Vector) Equal(o Vector) bool {
	return v.Quaternion.Equal(o.Quaternion)
}
