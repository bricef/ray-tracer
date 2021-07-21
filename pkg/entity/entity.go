package entity

import (
	"fmt"

	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

type EntityNode struct {
	transform  math.Transform
	components map[core.ComponentType]core.Component
	children   []core.Entity
}

func NewEntity() *EntityNode {
	return &EntityNode{
		transform:  math.NewTransform(),
		components: make(map[core.ComponentType]core.Component, 5),
		children:   make([]core.Entity, 0),
	}
}

// Transform proxy

func (e *EntityNode) Translate(x, y, z float64) core.Entity {
	e.transform = e.transform.Translate(x, y, z)
	return e
}

func (e *EntityNode) Scale(x, y, z float64) core.Entity {
	e.transform = e.transform.Scale(x, y, z)
	return e
}

func (e *EntityNode) RotateX(r float64) core.Entity {
	e.transform = e.transform.RotateX(r)
	return e
}

func (e *EntityNode) RotateY(r float64) core.Entity {
	e.transform = e.transform.RotateY(r)
	return e
}
func (e *EntityNode) RotateZ(r float64) core.Entity {
	e.transform = e.transform.RotateZ(r)
	return e
}
func (e *EntityNode) Shear(xy, xz, yx, yz, zx, zy float64) core.Entity {
	e.transform = e.transform.Shear(xy, xz, yx, yz, zx, zy)
	return e
}

func (e *EntityNode) MoveTo(p math.Point) core.Entity {
	e.transform = e.transform.MoveTo(p)
	return e
}

func (e *EntityNode) Transform() math.Transform {
	return e.transform
}

func (e *EntityNode) Position() math.Point {
	return e.Transform().Apply(math.NewPoint(0, 0, 0)).AsPoint()
}

// Composition

func (e *EntityNode) Components() []core.Component {
	vs := make([]core.Component, 0)
	for _, v := range e.components {
		vs = append(vs, v)
	}
	return vs
}
func (e *EntityNode) Children() []core.Entity {
	return e.children
}
func (e *EntityNode) AddComponent(c core.Component) core.Entity {
	e.components[c.Type()] = c
	return e
}

func (e *EntityNode) AddChild(c core.Entity) core.Entity {
	e.children = append(e.children, c)
	return e
}

func (e *EntityNode) GetComponent(t core.ComponentType) core.Component {
	if v, ok := e.components[t]; ok {
		return v
	} else {
		return nil
	}
}
func (e *EntityNode) RemoveComponent(t core.ComponentType) core.Entity {
	delete(e.components, t)
	return e
}

func (e *EntityNode) HasComponent(t core.ComponentType) bool {
	_, ok := e.components[t]
	return ok
}

func (e *EntityNode) Normal(worldPoint math.Point) math.Vector {
	if mesh := e.GetMesh(); mesh != nil {
		objectPoint := e.transform.Inverse().Apply(worldPoint)
		objectNormal := mesh.Normal(objectPoint)
		worldNormal := e.transform.Inverse().Transpose().Apply(objectNormal)
		worldNormalVector := math.NewVector(worldNormal.X(), worldNormal.Y(), worldNormal.Z())
		return worldNormalVector.Normalize()
	}
	return math.NewVector(0, 0, 0)
}

func (e *EntityNode) Tick(scene []core.Entity) {

	// Tick all dynamic components
	for _, comp := range e.components {
		switch ct := comp.(type) {
		case core.Dynamic:
			ct.Tick(e)
		}
	}

	// Tick all children
	for _, child := range e.children {
		child.Tick(scene)
	}

}

// Helpers

func (e *EntityNode) GetKinematic() core.Kinematic {
	if c := e.GetComponent(component.Kinematic); c != nil {
		return c.(core.Kinematic)
	}
	return nil
}

func (e *EntityNode) GetMaterial() core.Material {
	if c := e.GetComponent(component.Material); c != nil {
		return c.(core.Material)
	}
	return nil
}

func (e *EntityNode) GetMesh() core.Mesh {
	if c := e.GetComponent(component.Mesh); c != nil {
		return c.(core.Mesh)
	}
	return nil
}

func (e *EntityNode) GetLight() core.PointLight {
	if c := e.GetComponent(component.PointLight); c != nil {
		return c.(core.PointLight)
	}
	return nil
}

// Good practice

func (e *EntityNode) String() string {
	return fmt.Sprintf(`Entity(position: %v, components(%v): %v, children(%v): %v)`,
		e.transform.Position(),
		len(e.components), e.Components(),
		len(e.children), e.children,
	)
}
