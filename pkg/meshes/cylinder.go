package meshes

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type cylinder struct {
	min    float64
	max    float64
	capped bool
}

func CylinderMesh() core.Mesh {
	return &cylinder{m.Inf(-1), m.Inf(+1), false}
}

func CylinderMeshLimited(min float64, max float64) core.Mesh {
	return &cylinder{min, max, false}
}

func CylinderClosedMesh(min, max float64) core.Mesh {
	return &cylinder{min, max, true}
}

func (c *cylinder) Type() core.ComponentType {
	return component.Mesh
}

func (cy *cylinder) checkCap(r core.Ray, t float64) bool {
	x := r.Origin().X() + t*r.Direction().X()
	z := r.Origin().Z() + t*r.Direction().Z()
	return (x*x + z*z) <= 1
}

func (cy *cylinder) intersectCaps(r core.Ray) []float64 {
	if !cy.capped || utils.AlmostEqual(r.Direction().Y(), 0) {
		return []float64{}
	}
	ts := []float64{}

	t1 := (cy.min - r.Origin().Y()) / r.Direction().Y()
	if cy.checkCap(r, t1) {
		ts = append(ts, t1)
	}

	t2 := (cy.max - r.Origin().Y()) / r.Direction().Y()
	if cy.checkCap(r, t2) {
		ts = append(ts, t2)
	}

	return ts
}

func (cy *cylinder) Intersect(r core.Ray) []float64 {
	rdn := r.Direction().Normalize()
	rdx := rdn.X()
	rdz := rdn.Z()
	a := rdx*rdx + rdz*rdz

	ts := []float64{}

	if utils.AlmostEqual(a, 0.0) {
		return cy.intersectCaps(r)
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

	ts = append(ts, cy.intersectCaps(r)...)

	return ts
}

func (c *cylinder) Normal(p math.Point) math.Vector {
	dist := m.Pow(p.X(), 2) + m.Pow(p.Z(), 2)

	if dist < 1 && p.Y() >= (c.max-utils.Epsilon) {
		return math.NewVector(0, 1, 0)
	}

	if dist < 1 && p.Y() <= (c.min+utils.Epsilon) {
		return math.NewVector(0, -1, 0)
	}

	return math.NewVector(p.X(), 0, p.Z()).Normalize()
}
