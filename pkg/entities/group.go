package entities

import (
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entity"
)

func NewGroup(es ...core.Entity) core.Entity {
	e := entity.NewEntity()
	for _, ei := range es {
		e.AddChild(ei)
	}
	return e
}
