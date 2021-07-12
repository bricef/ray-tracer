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

func main() {
	width, height := 1000, 500

	// set up frame
	frame := canvas.NewImageCanvas(width, height)

	// set up scene
	s := scene.NewScene()

	wallMaterial := material.NewMaterial().
		SetColor(color.New(1, 0.9, 0.9)).
		SetSpecular(0.0)

	s.Add( // floor
		entity.NewSphere().
			SetMaterial(wallMaterial).
			SetTransform(
				transform.NewTransform().Scale(10, 0.01, 10),
			),
	)

	s.Add( // left floor
		entity.NewSphere().
			SetMaterial(wallMaterial).
			SetTransform(
				transform.NewTransform().
					Translate(0, 0, 5).
					RotateY(-math.Pi/4.0).
					RotateX(math.Pi/2.0).
					Scale(10, 0.01, 10),
			),
	)

	s.Add( // right wall
		entity.NewSphere().
			SetMaterial(wallMaterial).
			SetTransform(
				transform.NewTransform().
					Translate(0, 0, 5).
					RotateY(math.Pi/4.0).
					RotateX(math.Pi/2.0).
					Scale(10, 0.01, 10),
			),
	)

	s.Add( // middle
		entity.NewSphere().
			SetMaterial(
				material.NewMaterial().SetColor(color.New(0.1, 1.0, 0.5)).SetDiffuse(0.7).SetSpecular(0.3),
			).
			SetTransform(
				transform.NewTransform().Translate(-0.5, 1, 0.5),
			),
	)

	s.Add( // right
		entity.NewSphere().
			SetMaterial(
				material.NewMaterial().SetColor(color.New(0.1, 1.0, 0.5)).SetDiffuse(0.7).SetSpecular(0.3),
			).
			SetTransform(
				transform.NewTransform().Translate(1.5, 0.5, -0.5).Scale(0.5, 0.5, 0.5),
			),
	)

	s.Add( // left
		entity.NewSphere().
			SetMaterial(
				material.NewMaterial().SetColor(color.New(1, 0.8, 0.1)).SetDiffuse(0.7).SetSpecular(0.3),
			).
			SetTransform(
				transform.NewTransform().Translate(-1.5, 0.33, -0.75).Scale(0.33, 0.33, 0.33),
			),
	)

	// fmt.Printf("len(s.Entities)=%v\n", len(s.Entities))

	s.Add( // light
		light.NewPointLight(
			color.White,
			quaternion.NewPoint(-10, 10, -10),
		),
	)

	c := camera.
		CameraFromFOV(width, height, math.Pi/3.0).
		SetTransform(
			transform.ViewTransform(
				quaternion.NewPoint(0, 1.5, -5.0),
				quaternion.NewPoint(0, 1, 0),
				quaternion.NewVector(0, 1, 0)),
		)

	c.Render(s, frame)

	// Write out to file
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)
	outputFilename := path.Join(OUTPUT_DIR, "chapter7.png")
	frame.WritePNG(outputFilename)
	fmt.Printf("Wrote output to %v", outputFilename)
}
