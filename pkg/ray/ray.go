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
			T:          t,
			Entity:     e,
			Point:      p,
			OverPoint:  p.Add(n.Scale(utils.Epsilon)),
			UnderPoint: p.Sub(n.Scale(utils.Epsilon)),
			EyeVector:  eye,
			Normal:     n,
			Inside:     inside,

			ReflectVector: r.direction.Reflect(n),
			N1:            0.0,
			N2:            0.0,
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

func (r Ray) GetIntersections(es []core.Entity) *Intersections {
	xs := NewIntersections()
	for _, e := range es {
		mat := e.GetMaterial()
		mesh := e.GetMesh()
		if mat != nil && mesh != nil { // Ignore entities without mesh or material
			xs = xs.Merge(r.Intersect(e))
		}
		if e.HasChildren() {
			for _, child := range e.Children() {
				tr := r.Transform(e.Transform().Inverse()).
					Transform(child.Transform().Inverse()).(Ray)
				xs = xs.Merge(tr.Intersect(child))
			}
		}
	}

	for i, x := range xs.All {
		if i == 0 { // first item. assume 1.0 refraction incident
			x.N1 = 1.0
			x.N2 = x.Entity.GetMaterial().RefractiveIndex()
		} else if len(xs.All) > 1 && i == (len(xs.All)-1) { // last item
			x.N1 = xs.All[i-i].Entity.GetMaterial().RefractiveIndex()
			x.N2 = 1.0
		} else {
			if xs.All[i-1].Inside && x.Inside { // Both inside
				x.N1 = x.Entity.GetMaterial().RefractiveIndex()
				x.N2 = xs.All[i+1].Entity.GetMaterial().RefractiveIndex()
			} else if xs.All[i-1].Inside && !x.Inside { // Coming out of previous entity
				x.N1 = xs.All[i-1].Entity.GetMaterial().RefractiveIndex()
				x.N2 = 1.0
			} else if !xs.All[i-1].Inside && x.Inside { // ???
				x.N1 = xs.All[i-1].Entity.GetMaterial().RefractiveIndex()
				x.N2 = xs.All[i+1].Entity.GetMaterial().RefractiveIndex()
			} else { // both outside
				x.N1 = xs.All[i-1].Entity.GetMaterial().RefractiveIndex()
				x.N2 = x.Entity.GetMaterial().RefractiveIndex()
			}
		}
	}

	// objects := []core.Entity{}
	// hit := xs.Hit
	// for _, x := range xs.All {
	// 	if x == hit {
	// 		if len(objects) > 0 {
	// 			e := objects[len(objects)-1]
	// 			if e.GetMaterial() != nil {
	// 				x.N1 = e.GetMaterial().RefractiveIndex()
	// 			} else {
	// 				x.N1 = 1.0
	// 			}
	// 		} else {
	// 			x.N1 = 1.0
	// 		}
	// 	}

	// 	if core.Contains(objects, x.Entity) {
	// 		objects = core.Remove(objects, x.Entity)
	// 	} else {
	// 		objects = append(objects, x.Entity)
	// 	}

	// 	if x == hit {
	// 		if len(objects) == 0 {
	// 			x.N2 = 1.0
	// 		} else {
	// 			o := objects[len(objects)-1]
	// 			if o.GetMaterial() != nil {
	// 				x.N2 = o.GetMaterial().RefractiveIndex()
	// 			} else {
	// 				x.N2 = 1.0
	// 			}
	// 		}
	// 		break
	// 	}
	// }
	return xs

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
