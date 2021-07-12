package material

import "github.com/bricef/ray-tracer/color"

type Material struct {
	Color     color.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() *Material {
	return &Material{
		Color:     color.New(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m *Material) Equal(o *Material) bool {
	return m.Color.Equal(o.Color) && m.Ambient == o.Ambient && m.Diffuse == o.Diffuse && m.Specular == o.Specular && m.Shininess == o.Shininess
}

func (m *Material) SetAmbient(v float64) *Material {
	m.Ambient = v
	return m
}

func (m *Material) SetDiffuse(v float64) *Material {
	m.Diffuse = v
	return m
}

func (m *Material) SetSpecular(v float64) *Material {
	m.Specular = v
	return m
}

func (m *Material) SetShininess(v float64) *Material {
	m.Shininess = v
	return m
}

func (m *Material) SetColor(c color.Color) *Material {
	m.Color = c
	return m
}
