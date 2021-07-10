package ray

import (
	"fmt"
	"math"
	"sort"

	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/entity"
	"github.com/bricef/ray-tracer/light"
	q "github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/transform"
)

type Ray struct {
	Origin    q.Quaternion
	Direction q.Quaternion
}

func NewRay(o q.Quaternion, d q.Quaternion) Ray {
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
	T         float64
	Entity    *entity.Entity
	Point     q.Quaternion
	EyeVector q.Quaternion
	Normal    q.Quaternion
	Inside    bool
}

func (i Intersection) String() string {
	return fmt.Sprintf("Intersection(%v)", i.T)
}

func (i Intersection) Shade(l *light.PointLight) color.Color {
	return light.Phong(i.Entity.Material, l, i.Point, i.EyeVector, i.Normal)
}

func (i Intersection) ShadeAll(ls []*light.PointLight) color.Color {
	c := color.New(0, 0, 0)
	for _, l := range ls {
		c = c.Add(i.Shade(l))
	}
	return c
}

type Intersections struct {
	All []Intersection
	Hit *Intersection
}

func (is Intersections) Merge(xs Intersections) Intersections {
	newAll := []Intersection{}

	newAll = append(newAll, is.All...)
	newAll = append(newAll, xs.All...)

	sort.Slice(newAll, func(i, j int) bool {
		return newAll[i].T < newAll[j].T
	})

	return Intersections{
		All: newAll,
		Hit: &newAll[0],
	}
}

func (r Ray) Hit(e *entity.Entity) *Intersection {
	return r.Intersect(e).Hit
}

func (r Ray) Intersect(e *entity.Entity) Intersections {
	tray := r.Transform(e.Transform.Inverse())

	xs := make([]Intersection, 2)
	var hit *Intersection
	for i, v := range intersects(tray, e) {
		p := r.Position(v)
		n := e.Normal(p)
		eye := r.Direction.Invert()
		inside := false
		if n.Dot(eye) < 0 { // inside entity check
			n = n.Invert()
			inside = true
		}
		x := Intersection{
			T:         v,
			Entity:    e,
			Point:     p,
			EyeVector: eye,
			Normal:    n,
			Inside:    inside,
		}
		xs[i] = x
		if v >= 0 && ((hit == nil) || v < hit.T) {
			hit = &x
		}
	}
	return Intersections{All: xs, Hit: hit}
}

func intersects(r Ray, e *entity.Entity) []float64 {
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

func (r Ray) Transform(t transform.Transform) Ray {
	return NewRay(
		t.Apply(r.Origin),
		t.Apply(r.Direction),
	)
}
