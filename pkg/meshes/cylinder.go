package meshes

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type cylinder struct {
}

func CylinderMesh() core.Mesh {
	return &cylinder{}
}

func (c *cylinder) Type() core.ComponentType {
	return component.Mesh
}

func (cy *cylinder) Intersect(r core.Ray) []float64 {
	rdn := r.Direction().Normalize()
	rdx := rdn.X()
	rdz := rdn.Z()
	a := rdx*rdx + rdz*rdz

	if utils.AlmostEqual(a, 0.0) {
		return []float64{}
	}

	rox := r.Origin().X()
	roz := r.Origin().Z()

	b := 2*rox*rdx + 2*roz*rdz
	c := rox*rox + roz*roz - 1.0
	disc := b*b - 4*a*c

	if disc < 0.0 {
		return []float64{}
	}

	t0 := (-b - m.Sqrt(disc)) / (2 * a)
	t1 := (-b + m.Sqrt(disc)) / (2 * a)
	return []float64{t0, t1}
}

func (c *cylinder) Normal(p math.Point) math.Vector {
	return math.NewVector(p.X(), 0, p.Z()).Normalize()
}
