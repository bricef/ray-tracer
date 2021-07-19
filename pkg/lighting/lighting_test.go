package lighting_test

import (
	"fmt"
	m "math"
	"testing"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/scene"
)

type TestCase struct {
	Material *material.Material
	Position math.Point
	Eye      math.Vector
	Normal   math.Vector
	Light    core.Entity
	Expected color.Color
}

func TestPhongLightingScenarios(t *testing.T) {
	cases := []TestCase{
		{
			Material: material.NewMaterial(),
			Position: math.NewPoint(0, 0, 0),
			Eye:      math.NewVector(0, 0, -1),
			Normal:   math.NewVector(0, 0, -1),
			Light:    lighting.NewPointLight(color.White).Translate(0, 0, -10),
			Expected: color.New(1.9, 1.9, 1.9),
		},
		{
			Material: material.NewMaterial(),
			Position: math.NewPoint(0, 0, 0),
			Eye:      math.NewVector(0, m.Sqrt2/2, -m.Sqrt2/2),
			Normal:   math.NewVector(0, 0, -1),
			Light:    lighting.NewPointLight(color.White).Translate(0, 0, -10),
			Expected: color.New(1, 1, 1),
		},
		{
			Material: material.NewMaterial(),
			Position: math.NewPoint(0, 0, 0),
			Eye:      math.NewVector(0, 0, -1),
			Normal:   math.NewVector(0, 0, -1),
			Light:    lighting.NewPointLight(color.White).Translate(0, 10, -10),
			Expected: color.New(0.736396103, 0.736396103, 0.736396103),
		},
		{
			Material: material.NewMaterial(),
			Position: math.NewPoint(0, 0, 0),
			Eye:      math.NewVector(0, -m.Sqrt2/2, -m.Sqrt2/2),
			Normal:   math.NewVector(0, 0, -1),
			Light:    lighting.NewPointLight(color.White).Translate(0, 10, -10),
			Expected: color.New(1.6363961031, 1.6363961031, 1.6363961031),
		},
		{
			Material: material.NewMaterial(),
			Position: math.NewPoint(0, 0, 0),
			Eye:      math.NewVector(0, 0, -10),
			Normal:   math.NewVector(0, 0, -1),
			Light:    lighting.NewPointLight(color.White).Translate(0, 0, 10),
			Expected: color.New(0.1, 0.1, 0.1),
		},
	}

	for _, c := range cases {
		fmt.Printf("[CASE]: %v\n", c)
		result := lighting.Phong(
			c.Material,
			c.Light,
			c.Position,
			c.Eye,
			c.Normal,
		)
		if !result.Equal(c.Expected) {
			t.Errorf("Lighting failed [CASE]: %v. \nExpected %v, \ngot %v", c, c.Expected, result)
		}
	}
}

func TestLightingWithShadow(t *testing.T) {
	mat := material.NewMaterial()
	pos := math.NewPoint(0, 0, 0)
	eyev := math.NewVector(0, 0, -1)
	normal := math.NewVector(0, 0, -1)
	l := lighting.NewPointLight(color.White).Translate(0, 0, -10)

	result := lighting.PhongShadow(mat, l, pos, eyev, normal)
	expected := color.New(.1, .1, .1)
	if !result.Equal(expected) {
		t.Errorf("Shadowed Lighting failed. Expected %v, got %v", expected, result)
	}
}

type ShadowTestCase struct {
	Scene    *scene.Scene
	Point    math.Quaternion
	Light    core.Entity
	Expected bool
}

func TestIsShadowed(t *testing.T) {
	defaultScene := scene.DefaultScene()
	cases := []ShadowTestCase{
		{
			defaultScene,
			math.NewPoint(0, 10, 0),
			defaultScene.Lights()[0],
			false,
		},
		{
			defaultScene,
			math.NewPoint(10, -10, 10),
			defaultScene.Lights()[0],
			true,
		},
		{
			defaultScene,
			math.NewPoint(-20, 20, -20),
			defaultScene.Lights()[0],
			false,
		},
		{
			defaultScene,
			math.NewPoint(-2, 2, -2),
			defaultScene.Lights()[0],
			false,
		},
	}

	for _, c := range cases {
		result := c.Scene.Obstructed(c.Light.Position(), c.Point)
		if result != c.Expected {
			t.Errorf("ShadowTracing Failed [CASE]: %v. Expected %v, got %v", c, c.Expected, result)
		}
	}
}

func TestShadowHitRenderedCorrectly(t *testing.T) {
	s := scene.NewScene()
	s.Add(
		lighting.NewPointLight(color.White).Translate(0, 0, -10),
	)
	s.Add(entities.NewSphere())
	s.Add(
		entities.NewSphere().Translate(0, 0, 10),
	)

	r := ray.NewRay(
		math.NewPoint(0, 0, 5),
		math.NewVector(0, 0, 1),
	)

	result := s.Cast(r)

	expected := color.New(0.1, 0.1, 0.1)
	if !result.Equal(expected) {
		t.Errorf("Shading shadow failed. Expected %v, got %v", expected, result)
	}
}
