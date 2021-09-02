package ray

import (
	"fmt"
	m "math"

	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

type Intersection struct {
	T             float64
	Entity        core.Entity
	Point         math.Point
	EyeVector     math.Vector
	Normal        math.Vector
	Inside        bool
	OverPoint     math.Point
	UnderPoint    math.Point
	ReflectVector math.Vector
	N1            float64
	N2            float64
}

func (i *Intersection) String() string {
	return fmt.Sprintf("Intersection(%v\n%v)\n", i.T, i.Entity)
}

// Schlick approximation to Fresnel equations
func (i *Intersection) Schlick() float64 {
	cos := i.EyeVector.Dot(i.Normal)
	if i.N1 > i.N2 {
		r := i.N1 / i.N2
		sin2_t := r * r * (1 - (cos * cos))
		if sin2_t > 1.0 {
			return 1.0
		}
		cos_t := m.Sqrt(1 - sin2_t)
		cos = cos_t
	}
	r0 := m.Pow(((i.N1 - i.N2) / (i.N1 + i.N2)), 2)
	return r0 + (1-r0)*m.Pow((1-cos), 5)
}
