package ray

import (
	"fmt"

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
