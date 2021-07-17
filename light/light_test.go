package light_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/entity"
	"github.com/bricef/ray-tracer/light"
	"github.com/bricef/ray-tracer/material"
	"github.com/bricef/ray-tracer/quaternion"
	q "github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/ray"
	"github.com/bricef/ray-tracer/scene"
	"github.com/bricef/ray-tracer/transform"
)

type TestCase struct {
	Material *material.Material
	Position quaternion.Quaternion
	Eye      quaternion.Quaternion
	Normal   quaternion.Quaternion
	Light    *light.PointLight
	Expected color.Color
}

func TestPhongLightingScenarios(t *testing.T) {
	cases := []TestCase{
		{
			Material: material.NewMaterial(),
			Position: q.NewPoint(0, 0, 0),
			Eye:      q.NewVector(0, 0, -1),
			Normal:   q.NewVector(0, 0, -1),
			Light: light.NewPointLight(
				color.New(1, 1, 1),
				q.NewPoint(0, 0, -10),
			),
			Expected: color.New(1.9, 1.9, 1.9),
		},
		{
			Material: material.NewMaterial(),
			Position: q.NewPoint(0, 0, 0),
			Eye:      q.NewVector(0, math.Sqrt2/2, -math.Sqrt2/2),
			Normal:   q.NewVector(0, 0, -1),
			Light: light.NewPointLight(
				color.New(1, 1, 1),
				q.NewPoint(0, 0, -10),
			),
			Expected: color.New(1, 1, 1),
		},
		{
			Material: material.NewMaterial(),
			Position: q.NewPoint(0, 0, 0),
			Eye:      q.NewVector(0, 0, -1),
			Normal:   q.NewVector(0, 0, -1),
			Light: light.NewPointLight(
				color.New(1, 1, 1),
				q.NewPoint(0, 10, -10),
			),
			Expected: color.New(0.736396103, 0.736396103, 0.736396103),
		},
		{
			Material: material.NewMaterial(),
			Position: q.NewPoint(0, 0, 0),
			Eye:      q.NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2),
			Normal:   q.NewVector(0, 0, -1),
			Light: light.NewPointLight(
				color.New(1, 1, 1),
				q.NewPoint(0, 10, -10),
			),
			Expected: color.New(1.6363961031, 1.6363961031, 1.6363961031),
		},
		{
			Material: material.NewMaterial(),
			Position: q.NewPoint(0, 0, 0),
			Eye:      q.NewVector(0, 0, -10),
			Normal:   q.NewVector(0, 0, -1),
			Light: light.NewPointLight(
				color.New(1, 1, 1),
				q.NewPoint(0, 0, 10),
			),
			Expected: color.New(0.1, 0.1, 0.1),
		},
	}

	for _, c := range cases {
		fmt.Printf("[CASE]: %v\n", c)
		result := light.Phong(
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
	pos := q.NewPoint(0, 0, 0)
	eyev := quaternion.NewVector(0, 0, -1)
	normal := quaternion.NewVector(0, 0, -1)
	l := light.NewPointLight(color.White, quaternion.NewPoint(0, 0, -10))

	result := light.PhongShadow(mat, l, pos, eyev, normal)
	expected := color.New(.1, .1, .1)
	if !result.Equal(expected) {
		t.Errorf("Shadowed Lighting failed. Expected %v, got %v", expected, result)
	}
}

type ShadowTestCase struct {
	Scene    *scene.Scene
	Point    quaternion.Quaternion
	Light    *light.PointLight
	Expected bool
}

func TestIsShadowed(t *testing.T) {
	defaultScene := scene.DefaultScene()
	cases := []ShadowTestCase{
		{
			defaultScene,
			quaternion.NewPoint(0, 10, 0),
			defaultScene.Lights[0],
			false,
		},
		{
			defaultScene,
			quaternion.NewPoint(10, -10, 10),
			defaultScene.Lights[0],
			true,
		},
		{
			defaultScene,
			quaternion.NewPoint(-20, 20, -20),
			defaultScene.Lights[0],
			false,
		},
		{
			defaultScene,
			quaternion.NewPoint(-2, 2, -2),
			defaultScene.Lights[0],
			false,
		},
	}

	for _, c := range cases {
		result := c.Scene.Obstructed(c.Light.Position, c.Point)
		if result != c.Expected {
			t.Errorf("ShadowTracing Failed [CASE]: %v. Expected %v, got %v", c, c.Expected, result)
		}
	}
}

func TestShadowHitRenderedCorrectly(t *testing.T) {
	s := scene.NewScene()
	s.Add(
		light.NewPointLight(color.White, q.NewPoint(0, 0, -10)),
	)
	s.Add(entity.NewSphere())
	s.Add(
		entity.NewSphere().
			SetTransform(
				transform.NewTransform().Translate(0, 0, 10),
			),
	)

	r := ray.NewRay(
		q.NewPoint(0, 0, 5),
		q.NewVector(0, 0, 1),
	)

	result := s.Cast(r)

	expected := color.New(0.1, 0.1, 0.1)
	if !result.Equal(expected) {
		t.Errorf("Shading shadow failed. Expected %v, got %v", expected, result)
	}
}
