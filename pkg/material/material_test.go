package material_test

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/shaders"
)

func TestDefaultMaterial(t *testing.T) {
	m := material.NewMaterial()

	if !(m.Ambient() == 0.1 && m.Diffuse() == 0.9 && m.Specular() == 0.9 && m.Shininess() == 200.0) {
		t.Errorf("Failed to set up default material %v", m)
	}

}

func TestMaterialLigthingWithTexture(t *testing.T) {
	shader := shaders.Striped(
		shaders.Pigment(color.White),
		shaders.Pigment(color.Black),
	)

	mat := material.NewMaterial().
		SetAmbient(1.0).
		SetDiffuse(0.0).
		SetSpecular(0.0).
		SetShader(shader)

	light := lighting.NewPointLight(color.White)

	eyev := math.NewVector(0, 0, -1)
	normalv := math.NewVector(0, 0, -1)

	c1 := lighting.Phong(mat, light, math.NewPoint(0.9, 0, 0), eyev, normalv)
	c2 := lighting.Phong(mat, light, math.NewPoint(1.1, 0, 0), eyev, normalv)

	if !c1.Equal(color.White) || !c2.Equal(color.Black) {
		t.Errorf("Failed to render a striped texture.")
	}
}

func TestShaderWithObjectTransform(t *testing.T) {
	testShader := shaders.Test()
	mat := material.NewMaterial().SetShader(testShader)

	e := entities.NewSphere().Scale(2, 2, 2).AddComponent(mat)
	result := mat.ColorOn(e, math.NewPoint(2, 3, 4))
	expected := color.New(1, 1.5, 2)

	if !result.Equal(expected) {
		t.Errorf("Failed to render texture on object")
	}

}

func TestShaderWithShaderTransform(t *testing.T) {
	shaderTransform := math.Scale(2, 2, 2)
	testShader := shaders.With(shaderTransform, shaders.Test())
	mat := material.NewMaterial().SetShader(testShader)

	e := entities.NewSphere().AddComponent(mat)
	result := mat.ColorOn(e, math.NewPoint(2, 3, 4))
	expected := color.New(1, 1.5, 2)

	if !result.Equal(expected) {
		t.Errorf("Failed to render texture on object")
	}
}

func TestShaderWithShaderAndObjectTransform(t *testing.T) {
	shaderTransform := math.Translate(0.5, 1, 1.5)
	testShader := shaders.With(shaderTransform, shaders.Test())
	mat := material.NewMaterial().SetShader(testShader)

	e := entities.NewSphere().AddComponent(mat).Scale(2, 2, 2)
	result := mat.ColorOn(e, math.NewPoint(2.5, 3, 3.5))
	expected := color.New(0.75, 0.5, 0.25)

	if !result.Equal(expected) {
		t.Errorf("Failed to render texture on object")
	}
}
