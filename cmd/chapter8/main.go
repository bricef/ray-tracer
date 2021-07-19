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

func main() {
	width, height := 100, 50

	// set up frame
	frame := canvas.NewImageCanvas(width, height)

	// set up scene
	s := scene.NewScene()

	wallMaterial := material.NewMaterial()
	wallMaterial.SetColor(color.New(1, 0.9, 0.9))
	wallMaterial.SetSpecular(0.0)

	s.Add( // floor
		entities.NewSphere().
			AddComponent(wallMaterial).
			Scale(10, 0.01, 10),
	)

	s.Add( // left floor
		entities.NewSphere().
			AddComponent(wallMaterial).
			Translate(0, 0, 5).
			RotateY(-m.Pi/4.0).
			RotateX(m.Pi/2.0).
			Scale(10, 0.01, 10),
	)

	s.Add( // right wall
		entities.NewSphere().
			AddComponent(wallMaterial).
			Translate(0, 0, 5).
			RotateY(m.Pi/4.0).
			RotateX(m.Pi/2.0).
			Scale(10, 0.01, 10),
	)

	s.Add( // middle
		entities.NewSphere().
			AddComponent(
				material.NewMaterial().SetColor(color.New(0.1, 1.0, 0.5)).SetDiffuse(0.7).SetSpecular(0.3),
			).
			Translate(-0.5, 1, 0.5),
	)

	s.Add( // right
		entities.NewSphere().
			AddComponent(
				material.NewMaterial().SetColor(color.New(0.1, 1.0, 0.5)).SetDiffuse(0.7).SetSpecular(0.3),
			).
			Translate(1.5, 0.5, -0.5).
			Scale(0.5, 0.5, 0.5),
	)

	s.Add( // left
		entities.NewSphere().
			AddComponent(
				material.NewMaterial().SetColor(color.New(1, 0.8, 0.1)).SetDiffuse(0.7).SetSpecular(0.3),
			).
			Translate(-1.5, 0.33, -0.75).
			Scale(0.33, 0.33, 0.33),
	)

	s.Add( // light
		lighting.NewPointLight(color.White).Translate(-10, 10, -10),
	)

	c := camera.
		CameraFromFOV(width, height, m.Pi/3.0).
		SetTransform(
			math.ViewTransform(
				math.NewPoint(0, 1.5, -5.0),
				math.NewPoint(0, 1, 0),
				math.NewVector(0, 1, 0)),
		)

	c.Render(s, frame)

	// Write out to file
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)
	outputFilename := path.Join(OUTPUT_DIR, "chapter8.png")
	frame.WritePNG(outputFilename)
	fmt.Printf("Wrote output to %v", outputFilename)
}
