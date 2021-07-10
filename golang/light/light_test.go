package light

import (
	"fmt"
	"math"
	"testing"

	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/material"
	"github.com/bricef/ray-tracer/quaternion"
	q "github.com/bricef/ray-tracer/quaternion"
)

type TestCase struct {
	Material *material.Material
	Position quaternion.Quaternion
	Eye      quaternion.Quaternion
	Normal   quaternion.Quaternion
	Light    *PointLight
	Expected color.Color
}

func TestPhongLightingScenarios(t *testing.T) {
	cases := []TestCase{
		{
			Material: material.NewMaterial(),
			Position: q.NewPoint(0, 0, 0),
			Eye:      q.NewVector(0, 0, -1),
			Normal:   q.NewVector(0, 0, -1),
			Light: NewPointLight(
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
			Light: NewPointLight(
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
			Light: NewPointLight(
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
			Light: NewPointLight(
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
			Light: NewPointLight(
				color.New(1, 1, 1),
				q.NewPoint(0, 0, 10),
			),
			Expected: color.New(0.1, 0.1, 0.1),
		},
	}

	for _, c := range cases {
		fmt.Printf("[CASE]: %v\n", c)
		result := Phong(
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
