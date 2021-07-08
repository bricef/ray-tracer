package entity

import (
	"github.com/bricef/ray-tracer/mesh"
)

func NewSphere() *Entity {
	mesh := mesh.NewSphere()
	return New(&mesh)
}
