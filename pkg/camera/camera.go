package camera

import (
	"fmt"
	m "math"
	"time"

	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type Camera struct {
	Transform   math.Transform
	Distance    float64
	FrameWidth  int
	FrameHeight int
	FOV         float64
	PixelSize   float64
	Aspect      float64
	HalfWidth   float64
	HalfHeight  float64
}

func CameraFromFOV(w int, h int, fov float64) *Camera {
	halfView := m.Tan(fov / 2)
	aspect := float64(w) / float64(h)
	var halfHeight, halfWidth float64
	if aspect >= 1.0 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}
	return &Camera{
		Transform:   math.NewTransform(),
		Distance:    1.0,
		FrameWidth:  w,
		FrameHeight: h,
		FOV:         fov,
		PixelSize:   (2.0 * halfWidth) / float64(w),
		Aspect:      aspect,
		HalfWidth:   halfWidth,
		HalfHeight:  halfHeight,
	}
}

func (c *Camera) ProjectPixelRay(u, v int) ray.Ray {
	xoff := (float64(u) + 0.5) * c.PixelSize
	yoff := (float64(v) + 0.5) * c.PixelSize

	worldX := c.HalfWidth - xoff
	worldY := c.HalfHeight - yoff

	pixel := c.Transform.Inverse().Apply(math.NewPoint(worldX, worldY, -c.Distance))
	origin := c.Transform.Inverse().Apply(math.NewPoint(0, 0, 0)).AsPoint()
	direction := pixel.Sub(origin).AsVector().Normalize()
	return ray.NewRay(
		origin,
		direction,
	)
}

func (c *Camera) SetTransform(t math.Transform) *Camera {
	c.Transform = t
	return c
}

func (c *Camera) Render(s *scene.Scene, frame canvas.Canvas) {
	defer utils.TimeTrack(time.Now(), "Render")
	pixels := frame.Pixels()
	for pixels.Next() {
		u, v := pixels.Get()
		// if u == 0 {
		// 	fmt.Printf("Rendering Row %v of %v. (%v,%v)\n", v, frame.Height(), u, v)
		// }
		r := c.ProjectPixelRay(u, v)
		pix := s.Cast(r)
		frame.Set(u, v, pix)
	}
}

func (c *Camera) LookAt(e interface{}) *Camera {
	var target math.Quaternion
	switch t := e.(type) {
	case core.Entity:
		target = t.Transform().Apply(math.NewPoint(0, 0, 0))
	case math.Quaternion:
		if t.IsPoint() {
			target = t
		}
	default:
		panic(fmt.Errorf("Camera cannot look at %v", e))
	}

	self := c.Transform.Apply(math.NewPoint(0, 0, 0))
	c.Transform = math.ViewTransform(
		self,
		target,
		math.NewVector(0, 1, 0),
	)
	return c
}

func (c *Camera) MoveTo(p math.Point) *Camera {
	c.Transform = c.Transform.MoveTo(p)
	return c
}
