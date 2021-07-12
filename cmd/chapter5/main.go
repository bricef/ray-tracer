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
	frame := canvas.NewImageCanvas(100, 100)

	camera := camera.NewDeprecatedCamera(
		Point(0, 0, 25),
		Vector(0, 0, -1),
		8.0,
		camera.NewViewport(8, 8),
	)

	mat := material.NewMaterial()
	mat.Ambient = 1.0
	mat.Specular = 0.0
	mat.Diffuse = 0.0

	sphere := entity.NewSphere()
	sphere.SetTransform(Transform().Scale(6, 6, 6).Shear(0.5, 0, 0, 0, 0, 0))
	sphere.SetMaterial(mat)
	scene := []*entity.Entity{sphere}

	lights := []*light.PointLight{
		light.NewPointLight(color.New(1, 1, 1), Point(0, 0, 25)),
	}

	camera.Render(frame, scene, lights)

	// Write out to file
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)
	outputFilename := path.Join(OUTPUT_DIR, "chapter5.png")
	frame.WritePNG(outputFilename)
	fmt.Printf("Wrote output to %v", outputFilename)
}
