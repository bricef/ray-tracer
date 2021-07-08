package entity

import (
	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/transform"
)

type Entity struct {
	Transform transform.Transform
	Color     color.Color
}

func New() *Entity {
	return &Entity{
		Transform: transform.NewTransform(),
		Color:     color.New(1, 0, 1),
	}
}

func (e *Entity) SetTransform(t transform.Transform) *Entity {
	e.Transform = t
	return e
}
