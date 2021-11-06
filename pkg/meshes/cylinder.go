package meshes

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type cylinder struct {
	min float64
	max float64
}

func CylinderMesh() core.Mesh {
	return &cylinder{
		min: m.Inf(-1),
		max: m.Inf(+1),
	}
}

func CylinderMeshLimited(min float64, max float64) core.Mesh {
	return &cylinder{
		min: min,
		max: max,
	}
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
	if t0 > t1 { // swap
		t0, t1 = t1, t0
	}

	ts := []float64{}

	roy := r.Origin().Y()
	rdy := r.Direction().Y()

	y0 := roy + t0*rdy
	if cy.min < y0 && y0 < cy.max {
		ts = append(ts, t0)
	}

	y1 := roy + t1*rdy
	if cy.min < y1 && y1 < cy.max {
		ts = append(ts, t1)
	}

	return ts
}

func (c *cylinder) Normal(p math.Point) math.Vector {
	return math.NewVector(p.X(), 0, p.Z()).Normalize()
}
