package main

import (
	m "math"
	"path"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/utils"
)

func main() {
	width, height := 1000, 500

	// Set up output dir
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)

	// set up scene
	s := scene.NewScene()

	wallMaterial := material.NewMaterial().
		SetColor(color.New(1, 0.9, 0.9)).
		SetSpecular(0.0)

	s.Add(
		entities.NewSphere().
			AddComponent(
				material.NewMaterial(),
			).
			Translate(0, 0.5, 0),
	)

	s.Add(
		entities.NewPlane().
			AddComponent(wallMaterial),
	)

	s.Add( // light 1
		lighting.NewPointLight(color.New(0.2, 0.6, 0.3)).Translate(-10, 10, -10),
	)

	s.Add( // light 2
		lighting.NewPointLight(color.New(0.6, 0.2, 0.3)).Translate(10, 10, -10),
	)

	c := camera.
		CameraFromFOV(width, height, m.Pi/3.0).
		SetTransform(
			math.ViewTransform(
				math.NewPoint(0, 1.5, -5.0),
				math.NewPoint(0, 1, 0),
				math.NewVector(0, 1, 0)),
		)

	filepath := path.Join(OUTPUT_DIR, "chapter9.png")
	c.SaveFrame(s, filepath)
}
