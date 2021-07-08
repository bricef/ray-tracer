package entity

import (
	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/material"
	"github.com/bricef/ray-tracer/mesh"
	"github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/transform"
)

type Entity struct {
	Transform transform.Transform
	Color     color.Color
	Mesh      *mesh.Mesh
	Material  *material.Material
}

func New(m *mesh.Mesh) *Entity {
	return &Entity{
		Transform: transform.NewTransform(),
		Mesh:      m,
		Material:  material.NewMaterial(),
	}
}

func (e *Entity) SetMaterial(m *material.Material) *Entity {
	e.Material = m
	return e
}

func (e *Entity) SetTransform(t transform.Transform) *Entity {
	e.Transform = t
	return e
}

func (e *Entity) Normal(worldPoint quaternion.Quaternion) quaternion.Quaternion {

	objectPoint := e.Transform.Inverse().Apply(worldPoint)
	objectNormal := (*e.Mesh).Normal(objectPoint)
	worldNormal := e.Transform.Inverse().Transpose().Apply(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()

}
