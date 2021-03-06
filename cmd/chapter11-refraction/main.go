package main

import (
	m "math"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
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

	s.Add( // floor
		entities.NewSphere().
			AddComponent(floorMaterial).
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

	s.Add(
		entities.NewSphere().
			AddComponent(
				material.NewMaterial().
					SetColor(color.New(0.1, 0.1, 0.1)).
					SetDiffuse(0.0).
					SetSpecular(1.0).
					SetShininess(300).
					SetTransparency(0.9).
					SetReflective(0.9).
					SetRefractiveIndex(1.5),
			).
			Translate(-0.5, 1, 0.5),
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

	c.SaveFrame(s, "output/chapter11-refraction.png")
}
