package shaders

import (
	m "math"
	"math/rand"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
	opensimplex "github.com/ojrac/opensimplex-go"
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

func Stripes(a core.Shader, b core.Shader) core.Shader {
	return func(p math.Point) color.Color {
		var c color.Color
		if utils.AlmostEqual(m.Mod(m.Floor(p.X()), 2), 0) {
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

func LinearGradient(a, b color.Color) core.Shader {
	return func(p math.Point) color.Color {
		ratio := p.X() - m.Floor(p.X())
		return b.Sub(a).Scale(ratio).Add(a)
	}
}

func Rings(a, b core.Shader) core.Shader {
	return func(p math.Point) color.Color {
		distance := m.Sqrt(p.X()*p.X() + p.Z() + p.Z())
		if utils.AlmostEqual(m.Mod(m.Floor(distance), 2), 0.0) {
			return a(p)
		}
		return b(p)
	}
}
func Cubes(a, b core.Shader) core.Shader {
	return func(p math.Point) color.Color {
		sumfloors := m.Floor(p.X()) + m.Floor(p.Y()) + m.Floor(p.Z())
		if utils.AlmostEqual(m.Mod(sumfloors, 2.0), 0.0) {
			return a(p)
		}
		return b(p)
	}
}

func BlendBias(a, b core.Shader, bias float64) core.Shader {
	if bias <= 0.0 {
		return a
	}
	if bias >= 1.0 {
		return b
	}
	return func(p math.Point) color.Color {
		return a(p).Scale(1 - bias).Add(b(p).Scale(bias))
	}
}

func Blend(a, b core.Shader) core.Shader {
	return BlendBias(a, b, 0.5)
}

func OpenSimplex() core.Shader {
	r := opensimplex.NewNormalized(rand.Int63())
	g := opensimplex.NewNormalized(rand.Int63())
	b := opensimplex.NewNormalized(rand.Int63())
	return func(p math.Point) color.Color {
		return color.New(
			r.Eval3(p.X(), p.Y(), p.Z()),
			g.Eval3(p.X(), p.Y(), p.Z()),
			b.Eval3(p.X(), p.Y(), p.Z()),
		)
	}
}
