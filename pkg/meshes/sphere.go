package meshes

import (
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
