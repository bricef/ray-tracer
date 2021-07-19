package lighting

import (
	"fmt"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entity"
)

type pointLight struct {
	intensity color.Color
}

func NewPointLight(intensity color.Color) core.Entity {
	return entity.NewEntity().AddComponent(&pointLight{intensity})
}

func (p *pointLight) Intensity() color.Color {
	return p.intensity
}

func (l *pointLight) Type() core.ComponentType {
	return component.PointLight
}

func (l *pointLight) String() string {
	return fmt.Sprintf("Light(%v, %v, %v)", l.intensity.R, l.intensity.G, l.intensity.B)
}
