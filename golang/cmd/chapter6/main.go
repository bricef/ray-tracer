package main

import (
	"fmt"
	"path"

	"github.com/bricef/ray-tracer/camera"
	"github.com/bricef/ray-tracer/canvas"
	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/entity"
	"github.com/bricef/ray-tracer/light"
	"github.com/bricef/ray-tracer/material"

	. "github.com/bricef/ray-tracer/raytracer"
	"github.com/bricef/ray-tracer/utils"
)

func main() {
	// Set our canvas up for rendering
	frame := canvas.NewImageCanvas(1000, 1000)

	camera := camera.NewCamera(
		Point(0, 0, 25),
		Vector(0, 0, -1),
		8.0,
		camera.NewViewport(8, 8),
	)

	mat := material.NewMaterial()
	mat.Color = color.New(1, 0.2, 1)

	sphere := entity.NewSphere()
	sphere.SetTransform(
		Transform().Scale(10, 10, 10),
	)
	sphere.SetMaterial(mat)

	scene := []*entity.Entity{
		sphere,
	}

	lights := []*light.PointLight{
		light.NewPointLight(
			color.New(1, 1, 1),
			Point(-25, -25, 25),
		),
	}

	camera.Render(&frame, scene, lights)

	// Write out to file
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)
	outputFilename := path.Join(OUTPUT_DIR, "chapter6.png")
	frame.WritePNG(outputFilename)
	fmt.Printf("Wrote output to %v", outputFilename)
}
