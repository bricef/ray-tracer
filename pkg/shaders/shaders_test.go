package shaders

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

type ShaderTest struct {
	shader   core.Shader
	point    math.Point
	expected color.Color
}

func TestPigmentShader(t *testing.T) {

	c1 := Pigment(color.Red)(math.NewPoint(0, 23, 33))
	c2 := Pigment(color.Green)(math.NewPoint(0, 23, 33))

	if !c1.Equal(color.Red) {
		t.Errorf("Failed to set pigment color. Expected %v, got %v", color.Red, c1)
	}

	if !c2.Equal(color.Green) {
		t.Errorf("Failed to set pigment color. Expected %v, got %v", color.Green, c2)
	}

}

func TestStriped(t *testing.T) {
	shader := Striped(
		Pigment(color.White),
		Pigment(color.Black),
	)

	tests := []ShaderTest{
		// constant in y
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(0, 1, 0), color.White},
		{shader, math.NewPoint(0, 2, 0), color.White},
		// Constant in Z
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(0, 0, 1), color.White},
		{shader, math.NewPoint(0, 0, 2), color.White},
		// Changes in X
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(0.9, 0, 0), color.White},
		{shader, math.NewPoint(1, 0, 0), color.Black},
		{shader, math.NewPoint(-0.1, 0, 0), color.Black},
		{shader, math.NewPoint(-1, 0, 0), color.Black},
		{shader, math.NewPoint(-1.1, 0, 0), color.White},
	}

	for _, test := range tests {
		result := test.shader(test.point)
		if !result.Equal(test.expected) {
			t.Errorf("Shader %v failure at %v. Expected %v, got %v ", test.shader, test.point, test.expected, result)
		}
	}

}

func TestGradientShader(t *testing.T) {
	shader := LinearGradient(color.White, color.Black)

	tests := []ShaderTest{
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(0.25, 0, 0), color.New(0.75, 0.75, 0.75)},
		{shader, math.NewPoint(0.5, 0, 0), color.New(0.5, 0.5, 0.5)},
		{shader, math.NewPoint(0.75, 0, 0), color.New(0.25, 0.25, 0.25)},
	}

	for _, test := range tests {
		result := test.shader(test.point)
		if !result.Equal(test.expected) {
			t.Errorf("Shader %v failure at %v. Expected %v, got %v ", test.shader, test.point, test.expected, result)
		}
	}
}

func TestRingShader(t *testing.T) {
	shader := Rings(Pigment(color.White), Pigment(color.Black))

	tests := []ShaderTest{
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(1, 0, 0), color.Black},
		{shader, math.NewPoint(0, 0, 1), color.Black},
		{shader, math.NewPoint(0.708, 0, 0.708), color.Black},
	}

	for _, test := range tests {
		result := test.shader(test.point)
		if !result.Equal(test.expected) {
			t.Errorf("Shader %v failure at %v. Expected %v, got %v ", test.shader, test.point, test.expected, result)
		}
	}
}

func TestCubeShader(t *testing.T) {
	shader := Cubes(Pigment(color.White), Pigment(color.Black))

	tests := []ShaderTest{
		// in x
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(0.99, 0, 0), color.White},
		{shader, math.NewPoint(1.01, 0, 0), color.Black},
		// in y
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(0, 0.99, 0), color.White},
		{shader, math.NewPoint(0, 1.01, 0), color.Black},
		// in z
		{shader, math.NewPoint(0, 0, 0), color.White},
		{shader, math.NewPoint(0, 0, 0.99), color.White},
		{shader, math.NewPoint(0, 0, 1.01), color.Black},
	}

	for _, test := range tests {
		result := test.shader(test.point)
		if !result.Equal(test.expected) {
			t.Errorf("Shader %v failure at %v. Expected %v, got %v ", test.shader, test.point, test.expected, result)
		}
	}
}

func TestBlendShader(t *testing.T) {
	white := Pigment(color.White)
	black := Pigment(color.Black)

	tests := []ShaderTest{
		{BlendBias(white, black, 0.0), math.NewPoint(1, 2, 3), color.White},
		{BlendBias(white, black, 1.0), math.NewPoint(1, 2, 3), color.Black},
		{BlendBias(white, black, 0.5), math.NewPoint(1, 2, 3), color.New(0.5, 0.5, 0.5)},
		{BlendBias(white, black, 0.75), math.NewPoint(1, 2, 3), color.New(0.25, 0.25, 0.25)},
		// Blend is a shortcut for BlendBias(a,b,0.5)
		{Blend(white, black), math.NewPoint(1, 2, 3), color.New(0.5, 0.5, 0.5)},
	}

	for _, test := range tests {
		result := test.shader(test.point)
		if !result.Equal(test.expected) {
			t.Errorf("Shader %v failure at %v. Expected %v, got %v ", test.shader, test.point, test.expected, result)
		}
	}

}
