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
