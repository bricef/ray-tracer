package camera

import (
	"fmt"
	m "math"
	"path/filepath"
	"sync"
	"time"

	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/utils"
	"github.com/gosuri/uiprogress"
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

type Result struct {
	Pixel canvas.Pixel
	Color color.Color
}

func merge(cs ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	out := make(chan Result)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan Result) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func RenderGoroutine(pxs <-chan canvas.Pixel, s *scene.Scene, c *Camera) <-chan Result {
	out := make(chan Result)
	go func() {
		for px := range pxs {
			x, y := px.X, px.Y

			r := c.ProjectPixelRay(x, y)
			color := s.Cast(r)

			out <- Result{
				Pixel: px,
				Color: color,
			}
		}
		close(out)
	}()
	return out
}

const PARALLELISM = 16

func (c *Camera) Render(s *scene.Scene, frame canvas.Canvas) {
	defer utils.TimeTrack(time.Now(), "Render")
	pixels := frame.PixelChannel()
	cs := []<-chan Result{}

	for x := 0; x < PARALLELISM; x++ {
		cs = append(cs, RenderGoroutine(pixels, s, c))
	}
	for res := range merge(cs...) {
		frame.Set(res.Pixel.X, res.Pixel.Y, res.Color)
	}
}

func (c *Camera) SaveFrame(s *scene.Scene, filename string) {
	defer utils.TimeTrack(time.Now(), "SaveFrame")
	utils.EnsureDir(filepath.Dir(filename))
	// Set up frame to render to
	frame := canvas.NewImageCanvas(c.FrameWidth, c.FrameHeight)

	// Set up progress bar
	uiprogress.Start()
	bar := uiprogress.AddBar(frame.Width() * frame.Height())
	bar.AppendCompleted()
	bar.PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return filepath.Base(filename)
	})

	pixels := frame.PixelChannel()

	cs := []<-chan Result{}
	for x := 0; x < PARALLELISM; x++ {
		cs = append(cs, RenderGoroutine(pixels, s, c))
	}
	for res := range merge(cs...) {
		frame.Set(res.Pixel.X, res.Pixel.Y, res.Color)
		bar.Incr()
	}
	frame.WritePNG(filename)
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
