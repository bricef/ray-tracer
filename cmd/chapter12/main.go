package main

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/materials"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/shaders"
)

func main() {

	width, height := 1000, 500

	// set up scene
	s := scene.NewScene()

	floorMaterial := material.NewMaterial().
		SetAmbient(0.1).
		SetDiffuse(1.0).
		SetShader(
			shaders.With(
				math.Scale(1.5, 1.5, 1.5), //.RotateY(m.Pi/4),
				shaders.Cubes(
					shaders.Pigment(color.Black),
					shaders.Pigment(color.White),
				),
			),
		)

	wallMaterial := material.NewMaterial()
	wallMaterial.SetColor(color.New(1, 0.9, 0.9))
	wallMaterial.SetSpecular(0.0)

	s.Add( // floors
		entities.NewCube().
			AddComponent(floorMaterial).
			Scale(20, 20, 5),
	)

	s.Add( // walls
		entities.NewCube().
			AddComponent(wallMaterial).
			Scale(10, 10, 20),
	)

	s.Add(
		entities.NewSphere().
			AddComponent(materials.Glass()).
			Translate(-0.5, 1, 0.5),
	)

	cube := entities.NewCube()
	cube.Translate(-2, -5, 0)
	cube.RotateZ(m.Pi / 4.0)
	cube.RotateX(m.Pi / 4.0)
	cube.RotateY(m.Pi / 4.0)
	cube.GetMaterial().
		SetColor(color.Red).
		SetAmbient(0.2).
		SetDiffuse(0.6)
	s.Add(cube)

	s.Add( // light
		lighting.NewPointLight(color.White).Translate(-5, 5, 2),
	)

	c := camera.
		CameraFromFOV(width, height, m.Pi/3.0).
		SetTransform(
			math.ViewTransform(
				math.NewPoint(9, 9, 1),
				math.NewPoint(0, 0, 0),
				math.NewVector(0, 0, 1)),
		)

	c.SaveFrame(s, "output/chapter12.png")
}
