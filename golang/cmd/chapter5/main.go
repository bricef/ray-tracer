package main

import (
	"fmt"
	"path"

	"github.com/bricef/ray-tracer/camera"
	"github.com/bricef/ray-tracer/canvas"
	"github.com/bricef/ray-tracer/entity"

	. "github.com/bricef/ray-tracer/raytracer"
	"github.com/bricef/ray-tracer/shapes"
	"github.com/bricef/ray-tracer/utils"
)

func main() {
	// Set our canvas up for rendering
	frame := canvas.NewImageCanvas(100, 100)

	camera := camera.NewCamera(
		Point(0, 0, 25),
		Vector(0, 0, -1),
		8.0,
		camera.NewViewport(8, 8),
	)

	sphere := shapes.Sphere()
	sphere.SetTransform(Transform().Scale(3, 3, 3).Shear(0.5, 0, 0, 0, 0, 0))
	scene := []*entity.Entity{sphere}

	camera.Render(&frame, scene)

	// Write out to file
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)
	outputFilename := path.Join(OUTPUT_DIR, "chapter5.png")
	frame.WritePNG(outputFilename)
	fmt.Printf("Wrote output to %v", outputFilename)
}
