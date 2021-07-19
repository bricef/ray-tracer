package camera

import (
	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
)

type DeprecatedCamera struct {
	Position         math.Point
	Direction        math.Vector
	ViewportDistance float64
	Viewport         Viewport
}

func NewDeprecatedCamera(position math.Point, direction math.Vector, distance float64, viewport Viewport) DeprecatedCamera {
	return DeprecatedCamera{
		Position:         position,
		Direction:        direction,
		ViewportDistance: distance,
		Viewport:         viewport,
	}
}

type Viewport struct {
	Width  float64
	Height float64
}

func NewViewport(width float64, height float64) Viewport {
	return Viewport{
		width,
		height,
	}
}

func (v *Viewport) FrameXYToViewportXY(frame canvas.Canvas, fx int, fy int) (float64, float64) {
	Fx := float64(fx)
	Fy := float64(fy)

	Kx := v.Width / float64(frame.Width())
	Ky := v.Height / float64(frame.Height())

	Vx := (Kx * Fx) - (0.5 * v.Width)
	Vy := (Ky * Fy) - (0.5 * v.Height)

	// fmt.Printf("canvas: %v,%v -> viewport: %v,%v\n ", fx, fy, Vx, Vy)
	return Vx, Vy
}

func (c *DeprecatedCamera) Render(canvas canvas.Canvas, scene []core.Entity, lights []core.Entity) {

	pixels := canvas.Pixels()
	for pixels.More() {
		x, y := pixels.Get()

		vx, vy := c.Viewport.FrameXYToViewportXY(canvas, x, y)
		frameCenterToPixel := math.NewVector(vx, vy, 0)
		originToFrame := c.Direction.Normalize().Scale(c.ViewportDistance)
		originToPixel := originToFrame.Add(frameCenterToPixel).AsVector().Normalize()

		r := ray.NewRay(
			c.Position,
			originToPixel,
		)

		for _, e := range scene {
			hit := r.Hit(e)

			if hit != nil {
				hitPoint := r.Position(hit.T)
				pixelColor := color.New(0, 0, 0)
				for _, l := range lights {
					pixelColor = lighting.Phong(
						e.GetMaterial(),
						l,
						hitPoint,
						r.Direction().Invert(),
						e.Normal(hitPoint),
					)

				}
				canvas.Set(x, y, pixelColor)
			}
		}

	}
}
