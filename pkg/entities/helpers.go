package entities

import (
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entity"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/meshes"
)

func NewSphere() core.Entity {
	return entity.NewEntity().
		AddComponent(meshes.SphereMesh()).
		AddComponent(material.NewMaterial())
}

func NewPlane() core.Entity {
	return entity.NewEntity().
		AddComponent(meshes.PlaneMesh()).
		AddComponent(material.NewMaterial())
}

func NewGlassSphere() core.Entity {
	mat := material.NewMaterial().
		SetTransparency(1.0).
		SetRefractiveIndex(1.5)

	return entity.NewEntity().
		AddComponent(meshes.SphereMesh()).
		AddComponent(mat)
}

func NewCube() core.Entity {
	return entity.NewEntity().
		AddComponent(meshes.CubeMesh()).
		AddComponent(material.NewMaterial())
}

func NewCylinder() core.Entity {
	return entity.NewEntity().
		AddComponent(meshes.CylinderMeshLimited(0, 1)).
		AddComponent(material.NewMaterial())
}
func NewCappedCylinder() core.Entity {
	return entity.NewEntity().
		AddComponent(meshes.CylinderClosedMesh(0, 1)).
		AddComponent(material.NewMaterial())
}
