package material

import (
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
)

type Material struct {
	color     color.Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

func NewMaterial() *Material {
	return &Material{
		color:     color.New(1, 1, 1),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}

func (m *Material) Type() core.ComponentType {
	return component.Material
}

func (m *Material) Equal(o core.Material) bool {
	return m.color.Equal(o.Color()) && m.ambient == o.Ambient() && m.diffuse == o.Diffuse() && m.specular == o.Specular() && m.shininess == o.Shininess()
}

func (m *Material) SetAmbient(v float64) core.Material {
	m.ambient = v
	return m
}

func (m *Material) SetDiffuse(v float64) core.Material {
	m.diffuse = v
	return m
}

func (m *Material) SetSpecular(v float64) core.Material {
	m.specular = v
	return m
}

func (m *Material) SetShininess(v float64) core.Material {
	m.shininess = v
	return m
}

func (m *Material) SetColor(c color.Color) core.Material {
	m.color = c
	return m
}

func (m *Material) Color() color.Color {
	return m.color
}
func (m *Material) Ambient() float64 {
	return m.ambient
}
func (m *Material) Diffuse() float64 {
	return m.diffuse
}
func (m *Material) Specular() float64 {
	return m.specular
}
func (m *Material) Shininess() float64 {
	return m.shininess
}
