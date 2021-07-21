package ray

import (
	"fmt"

	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type Ray struct {
	origin    math.Point
	direction math.Vector
}

func NewRay(o math.Point, d math.Vector) Ray {
	return Ray{o, d}
}

func (r Ray) Origin() math.Point {
	return r.origin
}

func (r Ray) Direction() math.Vector {
	return r.direction
}

func (r Ray) Position(t float64) math.Point {
	return r.origin.Add(r.direction.Scale(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(%v,%v,%v -> %v,%v,%v)",
		r.origin.X(), r.origin.Y(), r.origin.Z(),
		r.direction.X(), r.direction.Y(), r.direction.Z(),
	)
}

func (r Ray) Hit(e core.Entity) *Intersection {
	return r.Intersect(e).Hit
}

func (r Ray) Intersect(e core.Entity) *Intersections {
	tray := r.Transform(e.Transform().Inverse())

	if m := e.GetMesh(); m == nil {
		return &Intersections{}
	}

	// icoords := intersects(tray, e) // []float64

	icoords := e.GetMesh().Intersect(tray)

	// Short circuit on miss
	if len(icoords) == 0 {
		return &Intersections{}
	}

	xs := make([]*Intersection, len(icoords))
	var hit *Intersection
	hit = nil
	for i, t := range icoords {
		p := r.Position(t)
		n := e.Normal(p)
		eye := r.direction.Invert()
		inside := false
		if n.Dot(eye) < 0 { // inside entity check
			n = n.Invert()
			inside = true
		}
		x := &Intersection{
			T:             t,
			Entity:        e,
			Point:         p,
			EyeVector:     eye,
			Normal:        n,
			Inside:        inside,
			OverPoint:     n.Scale(utils.Epsilon).Add(p),
			ReflectVector: r.direction.Reflect(n),
		}
		xs[i] = x
		if t >= 0 && ((hit == nil) || t < hit.T) {
			hit = x
		}
	}
	if hit != nil && hit.T < 0 {
		panic(fmt.Errorf("Hit.T < 0: %v", hit))
	}
	return &Intersections{All: xs, Hit: hit}
}

func (r Ray) Transform(t math.Transform) core.Ray {
	return NewRay(
		t.Apply(r.origin).AsPoint(),
		t.Apply(r.direction).AsVector(),
	)
}

func (a Ray) Equal(b core.Ray) bool {
	return a.direction.Equal(b.Direction()) && a.origin.Equal(b.Origin())
}
