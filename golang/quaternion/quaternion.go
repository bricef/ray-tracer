package quaternion

import (
	"math"
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
	return New(a.X+q.X, a.Y+q.Y, a.Z+q.Z, a.W+q.W)
}

func (a Quaternion) Equal(b interface{}) bool {
	q := From(b)
	return a.X == q.X && a.Y == q.Y && a.Z == q.Z && a.W == q.W
}

func (a Quaternion) Sub(b interface{}) Quaternion {
	q := From(b)
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
	return math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

func (q Vector) Normalize() Vector {
	return Vector{q.Divide(q.Magnitude())}
}

func (v Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z + v.W*o.W
}

func (v Vector) Cross(o Vector) (Vector, error) {
	return Vector{New(
		v.Y*o.Z-v.Z*o.Y,
		v.Z*o.X-v.X*o.Z,
		v.X*o.Y-v.Y*o.X,
		0.0,
	)}, nil
}

func (v Vector) Equal(o Vector) bool {
	return v.Quaternion.Equal(o.Quaternion)
}
