package lighting

import (
	"fmt"
	m "math"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

func Phong(
	mat core.Material,
	le core.Entity,
	point math.Point,
	eye math.Vector,
	normal math.Vector,
) color.Color {
	// surface color + light color
	l := le.GetLight()
	effectiveColor := mat.Color().Mult(l.Intensity())

	// Direction to light
	lightVector := le.Position().Sub(point).AsVector().Normalize()

	// Ambient contribution
	ambient := effectiveColor.Scale(mat.Ambient())

	var diffuse color.Color
	var specular color.Color

	dot := lightVector.Dot(normal)
	if dot < 0 { // light on other side of surface
		diffuse = color.Black
		specular = color.Black
	} else {
		diffuse = effectiveColor.Scale(mat.Diffuse() * dot)

		reflectVector := lightVector.Invert().Reflect(normal)
		reflectFactor := reflectVector.Dot(eye)

		if reflectFactor <= 0 { //light reflects away from eye
			specular = color.Black
		} else {
			factor := m.Pow(reflectFactor, mat.Shininess())
			specular = l.Intensity().Scale(mat.Specular() * factor)
		}
	}

	// fmt.Printf("Ambient: %v\nDiffuse: %v\nSpecular: %v\n", ambient, diffuse, specular)
	// fmt.Printf("Total: %v\n\n", ambient.Add(diffuse).Add(specular))

	return specular.Add(ambient).Add(diffuse)
}

func PhongShadow(
	m core.Material,
	le core.Entity,
	point math.Point,
	eye math.Vector,
	normal math.Vector,
) color.Color {
	l := le.GetLight()
	if l == nil {
		panic(fmt.Errorf("nil light passed to PhongShadow()"))
	}
	if m == nil {
		panic(fmt.Errorf("nil material passed to PhongShadow()"))
	}
	effectiveColor := m.Color().Mult(l.Intensity())
	ambient := effectiveColor.Scale(m.Ambient())
	return ambient
}
