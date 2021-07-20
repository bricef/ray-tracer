package shaders

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

func With(t math.Transform, s core.Shader) core.Shader {
	return func(p math.Point) color.Color {
		return s(t.Inverse().Apply(p))
	}
}

func Pigment(c color.Color) core.Shader {
	return func(p math.Point) color.Color {
		return c
	}
}

func Striped(a core.Shader, b core.Shader) core.Shader {
	return func(p math.Point) color.Color {
		var c color.Color
		if m.Mod(m.Floor(p.X()), 2) == 0 {
			c = a(p)
		} else {
			c = b(p)
		}
		return c
	}
}

func Test() core.Shader {
	return func(p math.Point) color.Color {
		return color.New(p.X(), p.Y(), p.Z())
	}
}
