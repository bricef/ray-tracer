package materials

import (
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/material"
)

func Glass() core.Material {
	return material.NewMaterial().
		SetColor(color.New(0.1, 0.1, 0.1)).
		SetDiffuse(0.0).
		SetSpecular(1.0).
		SetShininess(300).
		SetTransparency(0.9).
		SetReflective(0.9).
		SetRefractiveIndex(1.5)
}
