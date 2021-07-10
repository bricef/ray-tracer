package camera

import (
	"github.com/bricef/ray-tracer/canvas"
	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/entity"
	"github.com/bricef/ray-tracer/light"
	"github.com/bricef/ray-tracer/quaternion"
	. "github.com/bricef/ray-tracer/raytracer"
)

type DeprecatedCamera struct {
	Position         quaternion.Quaternion
	Direction        quaternion.Quaternion
	ViewportDistance float64
	Viewport         Viewport
}

func NewDeprecatedCamera(position quaternion.Quaternion, direction quaternion.Quaternion, distance float64, viewport Viewport) DeprecatedCamera {
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

func (c *DeprecatedCamera) Render(canvas canvas.Canvas, scene []*entity.Entity, lights []*light.PointLight) {

	pixels := canvas.Pixels()
	for pixels.Next() {
		x, y := pixels.Get()

		vx, vy := c.Viewport.FrameXYToViewportXY(canvas, x, y)
		frameCenterToPixel := Vector(vx, vy, 0)
		originToFrame := c.Direction.Normalize().Scale(c.ViewportDistance)
		originToPixel := originToFrame.Add(frameCenterToPixel)

		r := Ray(
			c.Position,
			originToPixel.Normalize(),
		)

		for _, e := range scene {
			hit := r.Hit(e)

			if hit != nil {
				hitPoint := r.Position(hit.T)
				pixelColor := color.New(0, 0, 0)
				for _, l := range lights {
					pixelColor = light.Phong(
						e.Material,
						l,
						hitPoint,
						r.Direction.Invert(),
						e.Normal(hitPoint),
					)

				}
				canvas.Set(x, y, pixelColor)
			}
		}

	}
}
