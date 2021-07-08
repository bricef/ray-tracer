package light

import (
	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/quaternion"
)

type PointLight struct {
	Intensity color.Color
	Position  quaternion.Quaternion
}

func NewPointLight(intensity color.Color, position quaternion.Quaternion) PointLight {
	return PointLight{
		Intensity: intensity,
		Position:  position,
	}
}
