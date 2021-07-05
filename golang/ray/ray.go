package ray

import (
	q "github.com/bricef/ray-tracer/quaternion"
)

type Ray struct {
	origin    q.Point
	direction q.Vector
}

func New(o q.Point, d q.Vector) Ray {
	return Ray{o, d}
}

func (r Ray) Position(t float64) q.Point {
	return q.Point{r.origin.Add(r.direction.Scale(t))}
}
