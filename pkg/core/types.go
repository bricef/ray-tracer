package core

import (
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/math"
)

type ComponentType int

type Component interface {
	Type() ComponentType
}

type Dynamic interface {
	Tick(owner Entity)
}

type Material interface {
	Component
	Equal(o Material) bool
	SetAmbient(v float64) Material
	SetDiffuse(v float64) Material
	SetSpecular(v float64) Material
	SetShininess(v float64) Material
	SetColor(c color.Color) Material
	Color() color.Color
	Ambient() float64
	Diffuse() float64
	Specular() float64
	Shininess() float64
}

type Mesh interface {
	Component
	Normal(meshPoint math.Point) math.Vector
}

type Kinematic interface {
	Dynamic
	Component
	SetAcceleration(math.Vector) Kinematic
	SetVelocity(math.Vector) Kinematic
}

type PointLight interface {
	Component
	Intensity() color.Color
}

type Entity interface {
	// Transform proxy
	Translate(x float64, y float64, z float64) Entity
	Scale(x float64, y float64, z float64) Entity
	RotateX(r float64) Entity
	RotateY(r float64) Entity
	RotateZ(r float64) Entity
	Shear(xy, xz, yx, yz, zx, zy float64) Entity
	MoveTo(math.Point) Entity
	Transform() math.Transform
	Position() math.Point

	// Composition
	Components() []Component
	Children() []Entity
	AddComponent(c Component) Entity
	AddChild(e Entity) Entity
	HasComponent(t ComponentType) bool
	GetComponent(t ComponentType) Component
	RemoveComponent(t ComponentType) Entity

	// Proxy to mesh (requires entity transform)
	Normal(worldPoint math.Point) math.Vector

	// Update method
	Tick(scene []Entity)

	// Helpers
	GetMesh() Mesh
	GetMaterial() Material
	GetKinematic() Kinematic
	GetLight() PointLight

	// Good practices
	String() string
}
