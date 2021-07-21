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
	ReflectVector math.Vector
}

func (i *Intersection) String() string {
	return fmt.Sprintf("Intersection(%v)", i.T)
}
