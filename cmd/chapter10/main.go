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
	"github.com/bricef/ray-tracer/pkg/shaders"
	"github.com/bricef/ray-tracer/pkg/utils"
)

func main() {
	width, height := 1000, 500

	// Set up output dir
	OUTPUT_DIR := "output"
	utils.EnsureDir(OUTPUT_DIR)

	// set up scene
	s := scene.NewScene()

	floorMaterial := material.NewMaterial().
		SetAmbient(0.1).
		SetDiffuse(1.0).
		SetShader(
			shaders.With(
				math.Scale(1.5, 1.5, 1.5), //.RotateY(m.Pi/4),
				shaders.Perturbed(
					0.2,
					10,
					shaders.Cubes(
						shaders.Pigment(color.Black),
						shaders.Pigment(color.White),
					),
				),
			),
		)

	s.Add(
		entities.NewPlane().
			AddComponent(floorMaterial),
	)

	wallMaterial1 := material.NewMaterial().
		SetAmbient(0.1).
		SetDiffuse(1.0).
		SetShader(
			shaders.With(
				math.Scale(0.5, 0.5, 0.5).RotateY(m.Pi/4),
				shaders.Cubes(
					shaders.Pigment(color.White),
					shaders.Pigment(color.Black),
				),
			),
		)

	wallMaterial2 := material.NewMaterial().
		SetAmbient(0.1).
		SetDiffuse(1.0).
		SetShader(
			shaders.With(
				math.RotateZ(m.Pi/4).Scale(1, 1, 1),
				shaders.Blend(
					shaders.Stripes(
						shaders.Pigment(color.White),
						shaders.Pigment(color.Red),
					),
					shaders.With(
						math.RotateZ(m.Pi/2),
						shaders.Stripes(
							shaders.Pigment(color.White),
							shaders.Pigment(color.Green),
						),
					),
				),
			),
		)

	s.Add(
		entities.NewPlane().
			Translate(5, 0, 5).
			RotateY(m.Pi / 4).
			RotateX(-m.Pi / 2).
			AddComponent(wallMaterial1),
	)

	s.Add(
		entities.NewPlane().
			Translate(-5, 0, 5).
			RotateY(-m.Pi / 4).
			RotateX(-m.Pi / 2).
			AddComponent(wallMaterial2),
	)

	s.Add(
		entities.NewSphere().
			AddComponent(
				material.NewMaterial().SetShader(
					shaders.With(
						math.Scale(0.1, 0.1, 0.1),
						shaders.OpenSimplex(),
					),
				),
			).
			Translate(0, 0.5, 0),
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
				math.NewPoint(0, 3, -7.0),
				math.NewPoint(0, 1, 0),
				math.NewVector(0, 1, 0)),
		)

	filepath := path.Join(OUTPUT_DIR, "chapter10.png")
	c.SaveFrame(s, filepath)
}
