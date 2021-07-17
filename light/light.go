package light

import (
	"math"

	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/material"
	"github.com/bricef/ray-tracer/quaternion"
)

type PointLight struct {
	Intensity color.Color
	Position  quaternion.Quaternion
}

func NewPointLight(intensity color.Color, position quaternion.Quaternion) *PointLight {
	return &PointLight{
		Intensity: intensity,
		Position:  position,
	}
}

func Phong(
	m *material.Material,
	l *PointLight,
	point quaternion.Quaternion,
	eye quaternion.Quaternion,
	normal quaternion.Quaternion,
) color.Color {
	// surface color + light color
	effectiveColor := m.Color.Mult(l.Intensity)

	// Direction to light
	lightVector := l.Position.Sub(point).Normalize()

	// Ambient contribution
	ambient := effectiveColor.Scale(m.Ambient)

	var diffuse color.Color
	var specular color.Color

	dot := lightVector.Dot(normal)
	if dot < 0 { // light on other side of surface
		diffuse = color.Black
		specular = color.Black
	} else {
		diffuse = effectiveColor.Scale(m.Diffuse * dot)

		reflectVector := lightVector.Invert().Reflect(normal)
		reflectFactor := reflectVector.Dot(eye)

		if reflectFactor <= 0 { //light reflects away from eye
			specular = color.Black
		} else {
			factor := math.Pow(reflectFactor, m.Shininess)
			specular = l.Intensity.Scale(m.Specular * factor)
		}
	}

	// fmt.Printf("Ambient: %v\nDiffuse: %v\nSpecular: %v\n", ambient, diffuse, specular)
	// fmt.Printf("Total: %v\n\n", ambient.Add(diffuse).Add(specular))

	return specular.Add(ambient).Add(diffuse)
}

func PhongShadow(
	m *material.Material,
	l *PointLight,
	point quaternion.Quaternion,
	eye quaternion.Quaternion,
	normal quaternion.Quaternion,
) color.Color {
	effectiveColor := m.Color.Mult(l.Intensity)
	ambient := effectiveColor.Scale(m.Ambient)
	return ambient
}
