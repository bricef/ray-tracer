package main

import (
	"fmt"
	m "math"
	"path"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/utils"
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

	earth := entities.NewSphere().
		AddComponent(
			material.NewMaterial().
				SetColor(color.New(0.1, 0.4, 0.8)).
				SetAmbient(0.2),
		).
		Scale(10, 10, 10)

	sun := lighting.NewPointLight(color.White).Translate(0, 0, 100)

	moonOffset := 6
	moon := entities.NewSphere().
		AddComponent(
			material.NewMaterial().
				SetColor(color.New(0.3, 0.3, 0.3)).
				SetSpecular(0.0),
		).
		Translate(float64(moonOffset), 0, 25)

	moon.GetKinematic().
		SetVelocity(
			math.NewVector(-float64(moonOffset*2)/float64(MAX_TICKS), 0, 0),
		)

	s.Add(earth)
	s.Add(moon)
	s.Add(sun)

	c := camera.
		CameraFromFOV(width, height, m.Pi/2.0).
		SetTransform(
			math.ViewTransform(
				math.NewPoint(0, 0, 30),
				math.NewPoint(0, 0, 0),
				math.NewVector(0, 1, 0)),
		)

	for tick := 0; tick <= MAX_TICKS; tick += 1 {
		s.Tick()
		filepath := path.Join(OUTPUT_DIR, fmt.Sprintf("frame-%v.png", tick))
		saveFrame(frame, c, s, filepath)
	}

	// Write out to file

}
