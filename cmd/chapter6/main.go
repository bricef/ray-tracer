package main

import (
	"fmt"
	"path"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/utils"
)

func main() {
	// Set our canvas up for rendering
	frame := canvas.NewImageCanvas(100, 100)

	camera := camera.NewDeprecatedCamera(
		math.NewPoint(0, 0, 25),
		math.NewVector(0, 0, -1),
		8.0,
		camera.NewViewport(8, 8),
	)

	mat := material.NewMaterial()
	mat.SetColor(color.New(1, 0.2, 1))

	sphere := entities.NewSphere()
	sphere.Scale(10, 10, 10)
	sphere.AddComponent(mat)

	scene := []core.Entity{
		sphere,
	}

	light := lighting.NewPointLight(color.New(1, 1, 1))
	light.Translate(-25, -25, 25)

	lights := []core.Entity{
		light,
	}

	camera.Render(frame, scene, lights)

	// Write out to file
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)
	outputFilename := path.Join(OUTPUT_DIR, "chapter6.png")
	frame.WritePNG(outputFilename)
	fmt.Printf("Wrote output to %v", outputFilename)
}
