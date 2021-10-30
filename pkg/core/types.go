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

type Shader func(p math.Point) color.Color

type Material interface {
	Component
	Equal(o Material) bool
	SetAmbient(v float64) Material
	SetDiffuse(v float64) Material
	SetSpecular(v float64) Material
	SetShininess(v float64) Material
	SetReflective(v float64) Material
	SetRefractiveIndex(v float64) Material
	SetTransparency(v float64) Material
	SetColor(c color.Color) Material
	SetShader(s Shader) Material
	ColorAt(math.Point) color.Color
	ColorOn(Entity, math.Point) color.Color
	Color() color.Color
	Ambient() float64
	Diffuse() float64
	Specular() float64
	Shininess() float64
	Reflective() float64
	Transparency() float64
	RefractiveIndex() float64
}

type Ray interface {
	Origin() math.Point
	Direction() math.Vector
	Position(t float64) math.Point
	Transform(math.Transform) Ray
	Equal(Ray) bool
}

type Mesh interface {
	Component
	Normal(meshPoint math.Point) math.Vector
	Intersect(Ray) []float64
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
	HasChildren() bool
	Parent() Entity
	AddComponent(c Component) Entity
	AddChild(e Entity) Entity
	SetParent(e Entity) Entity
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

	// Utilities
	String() string
}

func Contains(es []Entity, e Entity) bool {
	for _, a := range es {
		if a == e {
			return true
		}
	}
	return false
}
func Remove(es []Entity, e Entity) []Entity {
	for i, a := range es {
		if a == e {
			return append(es[:i], es[i+1:]...)
		}
	}
	return es
}
