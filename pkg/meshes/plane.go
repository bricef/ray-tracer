package meshes

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type planeMesh struct {
	normal math.Vector
}

func PlaneMesh() *planeMesh {
	return &planeMesh{math.NewVector(0, 1, 0)}
}

func (s *planeMesh) Type() core.ComponentType {
	return component.Mesh
}

func (plane *planeMesh) Normal(p math.Point) math.Vector {
	return plane.normal
}

func (plane *planeMesh) Intersect(r core.Ray) []float64 {
	// Parallel ray to plane
	if m.Abs(r.Direction().Y()) < utils.Epsilon {
		return []float64{}
	}
	t := -r.Origin().Y() / r.Direction().Y()
	return []float64{t}
}

func (plane *planeMesh) String() string {
	return "PlaneMesh()"
}
