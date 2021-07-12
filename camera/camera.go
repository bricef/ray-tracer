package camera

import (
	"fmt"
	"math"
	"time"

	"github.com/bricef/ray-tracer/canvas"
	"github.com/bricef/ray-tracer/entity"
	"github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/ray"
	"github.com/bricef/ray-tracer/scene"
	"github.com/bricef/ray-tracer/transform"
	"github.com/bricef/ray-tracer/utils"
)

type Camera struct {
	Transform   transform.Transform
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
	halfView := math.Tan(fov / 2)
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
		Transform:   transform.NewTransform(),
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

	pixel := c.Transform.Inverse().Apply(quaternion.NewPoint(worldX, worldY, -c.Distance))
	origin := c.Transform.Inverse().Apply(quaternion.NewPoint(0, 0, 0))
	direction := pixel.Sub(origin).Normalize()
	return ray.NewRay(
		origin,
		direction,
	)
}

func (c *Camera) SetTransform(t transform.Transform) *Camera {
	c.Transform = t
	return c
}

func (c *Camera) Render(s *scene.Scene, frame canvas.Canvas) {
	defer utils.TimeTrack(time.Now(), "Render")
	pixels := frame.Pixels()
	for pixels.Next() {
		u, v := pixels.Get()
		r := c.ProjectPixelRay(u, v)
		pix := s.Shade(r)
		frame.Set(u, v, pix)
	}
}

func (c *Camera) LookAt(e interface{}) *Camera {
	var target quaternion.Quaternion
	switch t := e.(type) {
	case entity.Entity:
		target = t.Transform.Apply(quaternion.NewPoint(0, 0, 0))
	case quaternion.Quaternion:
		if quaternion.IsPoint(t) {
			target = t
		}
	default:
		panic(fmt.Errorf("Camera cannot look at %v", e))
	}

	self := c.Transform.Apply(quaternion.NewPoint(0, 0, 0))
	c.Transform = transform.ViewTransform(
		self,
		target,
		quaternion.NewVector(0, 1, 0),
	)
	return c
}

func (c *Camera) MoveTo(q quaternion.Quaternion) *Camera {
	c.Transform = c.Transform.MoveTo(q)
	return c
}