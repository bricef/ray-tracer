package main

import (
	"fmt"
	m "math"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/utils"
)

const Pi = m.Pi

func corner(mat core.Material) core.Entity {
	corner := entities.NewSphere().
		// Translate(0, 0, 0).
		Scale(0.5, 0.5, 0.5).
		AddComponent(mat)
	return corner
}

func edge(mat core.Material) core.Entity {
	edge := entities.NewCube().
		Scale(0.25, 1, 0.25).
		// Translate(0, 0, -1)
		// RotateZ(-m.Pi / 2.0).
		// RotateY(-m.Pi / 2.0).
		AddComponent(mat)
	return edge
}

func side(mat core.Material) core.Entity {
	return entities.NewGroup(corner(mat), edge(mat))
}

// func hexagon(mat core.Material) core.Entity {
// 	h := entities.NewGroup()

// 	for n := 0.0; n <= 5; n++ {
// 		s := side(mat)
// 		s.RotateY(n * m.Pi / 3.0)
// 		h.AddChild(s)
// 	}
// 	return h
// }

func main() {

	width, height := 500, 250

	// set up scene
	s := scene.NewScene()

	s.Add( // light
		lighting.NewPointLight(color.White).Translate(-5, 5, 2),
	)

	mat := material.NewMaterial()
	mat.SetColor(color.New(1, 0.9, 0.9))
	mat.SetSpecular(0.0)
	mat.SetAmbient(0.7)

	// s.Add(side(mat))
	// for n := range [1]int{} {
	// 	hs := side(mat)
	// 	hs.RotateY(float64(n) * m.Pi / 3.0)
	// 	child := hs.Children()[0]
	// 	fmt.Printf("CHILD: %v,\nPARENT: %v\n\n", child, child.Parent().Transform())
	// 	s.Add(hs)
	// }

	cube := entities.NewCube().Scale(0.25, 1, 0.25).AddComponent(mat)
	wrappedCube := entities.NewGroup(entities.NewCube().Scale(0.25, 1, 0.25).AddComponent(mat))
	s.Add(cube)
	s.Add(wrappedCube)

	fmt.Printf("Group transform: %v", wrappedCube.Transform())

	s.Show()
	fmt.Println()
	// g1 := entities.NewGroup(
	// 	entities.NewSphere().AddComponent(mat),
	// 	entities.NewCylinder().
	// 		Translate(0, 1, 0).
	// 		Scale(1, 1.1, 1).
	// 		AddComponent(mat),
	// ).Translate(-0.5, 1, 0.5)

	// s.Add(g1)
	// s.Add(corner)

	c := camera.
		CameraFromFOV(width, height, utils.DegressToRadians(70)).
		SetTransform(
			math.ViewTransform(
				math.NewPoint(5, 2, -5),
				math.NewPoint(0, 0, 0),
				math.NewVector(0, 1, 0)),
		)

	c.SaveFrame(s, "output/chapter14.png")
}
