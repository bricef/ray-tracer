package meshes

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

type cube struct {
}

func CubeMesh() core.Mesh {
	return &cube{}
}

func (c *cube) Type() core.ComponentType {
	return component.Mesh
}

func check_axis(origin float64, direction float64) (float64, float64) {
	tmin := (-1 - origin) / direction
	tmax := (1 - origin) / direction
	if tmin > tmax {
		return tmax, tmin
	} else {
		return tmin, tmax
	}
}

func (c *cube) Intersect(r core.Ray) []float64 {
	txmin, txmax := check_axis(r.Origin().X(), r.Direction().X())
	tymin, tymax := check_axis(r.Origin().Y(), r.Direction().Y())
	tzmin, tzmax := check_axis(r.Origin().Z(), r.Direction().Z())
	tmin := m.Max(txmin, m.Max(tymin, tzmin))
	tmax := m.Min(txmax, m.Min(tymax, tzmax))
	if tmin > tmax {
		return []float64{}
	}
	return []float64{tmin, tmax}
}

func (c *cube) Normal(p math.Point) math.Vector {
	maxc := m.Max(m.Abs(p.X()), m.Max(m.Abs(p.Y()), m.Abs(p.Z())))

	if maxc == m.Abs(p.X()) {
		return math.NewVector(p.X(), 0, 0)
	} else if maxc == m.Abs(p.Y()) {
		return math.NewVector(0, p.Y(), 0)
	}
	return math.NewVector(0, 0, p.Z())

}
