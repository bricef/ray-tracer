package entities

import (
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entity"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/meshes"
)

func NewSphere() core.Entity {
	return entity.NewEntity().AddComponent(meshes.SphereMesh()).AddComponent(material.NewMaterial())
}
