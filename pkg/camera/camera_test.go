package camera_test

import (
	m "math"
	"testing"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/scene"
)

const halfPi = m.Pi / 2.0

func TestCameraInitialisation(t *testing.T) {
	c := camera.CameraFromFOV(160, 120, halfPi)
	if !(c.FrameWidth == 160 &&
		c.FrameHeight == 120 &&
		c.FOV == m.Pi/2) {
		t.Errorf("Failed to initialise a camera")
	}
}

func TestHorizontalCameraPixelSize(t *testing.T) {
	c := camera.CameraFromFOV(200, 125, halfPi)
	expected := 0.01
	if c.PixelSize != expected {
		t.Errorf("Failed to calculate pixel size for a horizontal camera")
	}
}

func TestVerticalCameraPixelSize(t *testing.T) {
	c := camera.CameraFromFOV(125, 200, halfPi)
	expected := 0.01
	if c.PixelSize != expected {
		t.Errorf("Failed to calculate pixel size for a horizontal camera")
	}
}

type RayCase struct {
	camera   *camera.Camera
	u        int
	v        int
	expected ray.Ray
}

func TestCameraRayProjection(t *testing.T) {
	c := camera.CameraFromFOV(201, 101, halfPi)

	cases := []RayCase{
		{c, 100, 50, ray.NewRay(
			math.NewPoint(0, 0, 0),
			math.NewVector(0, 0, -1),
		)},
	}

	for _, test := range cases {
		r := test.camera.ProjectPixelRay(test.u, test.v)
		if !r.Equal(test.expected) {
			t.Errorf("Failed to project pixel through camera. Expected %v, got %v", test.expected, r)
		}
	}
}

func TestCameraRender(t *testing.T) {
	s := scene.DefaultScene()
	c := camera.CameraFromFOV(11, 11, halfPi)
	c.SetTransform(
		math.ViewTransform(
			math.NewPoint(0, 0, -5),
			math.NewPoint(0, 0, 0),
			math.NewVector(0, 1, 0),
		),
	)
	frame := canvas.NewImageCanvas(11, 11)
	c.Render(s, frame)
	result, _ := frame.Get(5, 5)
	expected := color.New(0.38066, 0.47583, 0.2855)
	if !result.Equal(expected) {
		t.Errorf("failed to render to canvas. Expected %v, got %v", expected, result)
	}

}
