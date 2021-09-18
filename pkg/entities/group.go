package entities

import (
	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entity"
	"github.com/bricef/ray-tracer/pkg/math"
)

type gmesh struct {
	entity core.Entity
}

func (m *gmesh) Type() core.ComponentType {
	return component.Mesh
}

func (m *gmesh) Normal(meshPoint math.Point) math.Vector {
	return math.NewVector(0, 0, 0)
}

func (m *gmesh) Intersect(r core.Ray) []float64 {
	ts := []float64{}
	for _, e := range m.entity.Children() {
		mesh := e.GetMesh()
		if mesh != nil {
			ts = append(ts, mesh.Intersect(r)...)
		}
	}
	return ts
}

func GroupMesh(e core.Entity) core.Mesh {
	return &gmesh{e}
}

func NewGroup(es ...core.Entity) core.Entity {
	e := entity.NewEntity()
	e.AddComponent(GroupMesh(e))
	for _, ei := range es {
		e.AddChild(ei)
	}
	return e
}
