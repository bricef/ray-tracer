package meshes

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

type sphere struct {
}

func SphereMesh() core.Mesh {
	return &sphere{}
}

func (s *sphere) Type() core.ComponentType {
	return component.Mesh
}

func (s *sphere) Normal(p math.Point) math.Vector {
	return p.Sub(math.NewPoint(0, 0, 0)).AsVector().Normalize()
}

func (s *sphere) Intersect(r core.Ray) []float64 {
	sphere_to_ray := r.Origin().Sub(math.NewPoint(0, 0, 0)).AsVector()
	a := r.Direction().Dot(r.Direction())
	b := 2 * r.Direction().Dot(sphere_to_ray)
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
