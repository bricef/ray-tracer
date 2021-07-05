package entity

import (
	"github.com/bricef/ray-tracer/transform"
)

type Entity struct {
	Transform transform.Transform
}

func New() *Entity {
	return &Entity{
		Transform: transform.NewTransform(),
	}
}

func (e *Entity) SetTransform(t transform.Transform) *Entity {
	e.Transform = t
	return e
}
