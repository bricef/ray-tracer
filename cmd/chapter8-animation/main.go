package main

import (
	"fmt"
	"math"
	"path"

	"github.com/bricef/ray-tracer/camera"
	"github.com/bricef/ray-tracer/canvas"
	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/entity"
	"github.com/bricef/ray-tracer/light"
	"github.com/bricef/ray-tracer/material"
	"github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/scene"
	"github.com/bricef/ray-tracer/transform"
	"github.com/bricef/ray-tracer/utils"
)

func saveFrame(frame *canvas.ImageCanvas, c *camera.Camera, s *scene.Scene, filepath string) {

	c.Render(s, frame)
	frame.WritePNG(filepath)
	fmt.Printf("Wrote output to %v\n", filepath)
}

func main() {
	width, height := 500, 250
	MAX_TICKS := 100

	// Set up output dir
	OUTPUT_DIR := "output/chapter8/animation"
	utils.EnsureDir(OUTPUT_DIR)

	frame := canvas.NewImageCanvas(width, height)
	s := scene.NewScene()

	earth := entity.
		NewSphere().
		SetMaterial(
			material.NewMaterial().
				SetColor(color.New(0.1, 0.4, 0.8)).
				SetAmbient(0.2),
		).
		SetTransform(
			transform.NewTransform().Scale(10, 10, 10),
		)

	sun := light.NewPointLight(
		color.White,
		quaternion.NewPoint(0, 0, 100))

	moonOffset := 6
	moon := entity.
		NewSphere().
		SetMaterial(
			material.NewMaterial().
				SetColor(color.New(0.3, 0.3, 0.3)).
				SetSpecular(0.0),
		).
		SetTransform(
			transform.NewTransform().Translate(float64(moonOffset), 0, 25),
		).
		SetVelocity(
			quaternion.NewVector(-float64(moonOffset*2)/float64(MAX_TICKS), 0, 0),
		)

	s.Add(earth)
	s.Add(moon)
	s.Add(sun)

	c := camera.
		CameraFromFOV(width, height, math.Pi/2.0).
		SetTransform(
			transform.ViewTransform(
				quaternion.NewPoint(0, 0, 30),
				quaternion.NewPoint(0, 0, 0),
				quaternion.NewVector(0, 1, 0)),
		)

	for tick := 0; tick <= MAX_TICKS; tick += 1 {
		s.Tick()
		filepath := path.Join(OUTPUT_DIR, fmt.Sprintf("frame-%v.png", tick))
		saveFrame(frame, c, s, filepath)
	}

	// Write out to file

}
