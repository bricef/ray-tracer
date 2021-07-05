package ray

import (
	"fmt"
	"math"

	"github.com/bricef/ray-tracer/entity"
	q "github.com/bricef/ray-tracer/quaternion"
)

type Ray struct {
	Origin    q.Quaternion
	Direction q.Quaternion
}

func New(o q.Quaternion, d q.Quaternion) Ray {
	return Ray{o, d}
}

func (r Ray) Position(t float64) q.Quaternion {
	return r.Origin.Add(r.Direction.Scale(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(%v,%v,%v -> %v,%v,%v)",
		r.Origin.X, r.Origin.Y, r.Origin.Z,
		r.Direction.X, r.Direction.Y, r.Direction.Z,
	)
}

type Intersection struct {
	T      float64
	Entity *entity.Entity
}
type Intersections []Intersection

func intersects(r Ray, e entity.Entity) []float64 {
	sphere_to_ray := r.Origin.Sub(q.NewPoint(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1.0
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []float64{}
	} else {
		return []float64{
			(-b - math.Sqrt(discriminant)) / (2 * a),
			(-b + math.Sqrt(discriminant)) / (2 * a),
		}
	}
}
