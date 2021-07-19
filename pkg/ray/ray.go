package ray

import (
	"fmt"
	m "math"
	"sort"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type Ray struct {
	Origin    math.Point
	Direction math.Vector
}

func NewRay(o math.Point, d math.Vector) Ray {
	return Ray{o, d}
}

func (r Ray) Position(t float64) math.Quaternion {
	return r.Origin.Add(r.Direction.Scale(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(%v,%v,%v -> %v,%v,%v)",
		r.Origin.X(), r.Origin.Y(), r.Origin.Z(),
		r.Direction.X(), r.Direction.Y(), r.Direction.Z(),
	)
}

type Intersection struct {
	T         float64
	Entity    core.Entity
	Point     math.Point
	EyeVector math.Vector
	Normal    math.Vector
	Inside    bool
	OverPoint math.Point
}

func (i Intersection) String() string {
	return fmt.Sprintf("Intersection(%v)", i.T)
}

func (i Intersection) Shade(l core.Entity) color.Color {
	if l == nil {
		panic(fmt.Errorf("trying to shade with nil light"))
	}
	if i.Entity == nil {
		panic(fmt.Errorf("trying to shade with nil entity"))
	}
	mat := i.Entity.GetMaterial()

	return lighting.Phong(mat, l, i.Point, i.EyeVector, i.Normal)
}

func (i Intersection) ShadeAll(ls []core.Entity) color.Color {
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
	// Short circuit the empty case.
	if (len(is.All) + len(xs.All)) == 0 {
		return Intersections{
			All: []Intersection{},
			Hit: nil,
		}
	}

	newAll := []Intersection{}

	newAll = append(newAll, is.All...)
	newAll = append(newAll, xs.All...)

	sort.Slice(newAll, func(i, j int) bool {
		return newAll[i].T < newAll[j].T
	})

	var hit *Intersection
	for i, x := range newAll {
		if x.T > 0 && ((hit == nil) || x.T < hit.T) {
			hit = &newAll[i]
		}
	}

	return Intersections{
		All: newAll,
		Hit: hit,
	}
}

func (r Ray) Hit(e core.Entity) *Intersection {
	return r.Intersect(e).Hit
}

func (r Ray) Intersect(e core.Entity) Intersections {
	tray := r.Transform(e.Transform().Inverse())

	icoords := intersects(tray, e) // []float64

	// Short circuit on miss
	if len(icoords) == 0 {
		return Intersections{}
	}

	xs := make([]Intersection, len(icoords))
	var hit *Intersection
	hit = nil
	for i, t := range icoords {
		p := r.Position(t)
		n := e.Normal(p)
		eye := r.Direction.Invert()
		inside := false
		if n.Dot(eye) < 0 { // inside entity check
			n = n.Invert()
			inside = true
		}
		x := Intersection{
			T:         t,
			Entity:    e,
			Point:     p,
			EyeVector: eye,
			Normal:    n,
			Inside:    inside,
			OverPoint: n.Scale(utils.Epsilon).Add(p),
		}
		xs[i] = x
		if t >= 0 && ((hit == nil) || t < hit.T) {
			hit = &x
		}
	}
	if hit != nil && hit.T < 0 {
		panic(fmt.Errorf("Hit.T < 0: %v", hit))
	}
	return Intersections{All: xs, Hit: hit}
}

func intersects(r Ray, e core.Entity) []float64 {
	sphere_to_ray := r.Origin.Sub(math.NewPoint(0, 0, 0)).AsVector()
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1.0
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []float64{}
	} else {
		return []float64{
			(-b - m.Sqrt(discriminant)) / (2 * a),
			(-b + m.Sqrt(discriminant)) / (2 * a),
		}
	}
}

func (r Ray) Transform(t math.Transform) Ray {
	return NewRay(
		t.Apply(r.Origin).AsPoint(),
		t.Apply(r.Direction).AsVector(),
	)
}

func (a Ray) Equal(b Ray) bool {
	return a.Direction.Equal(b.Direction) && a.Origin.Equal(b.Origin)
}
