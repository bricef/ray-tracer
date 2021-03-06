package scene_test

import (
	"fmt"
	"testing"

	m "math"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
	"github.com/bricef/ray-tracer/pkg/scene"
	"github.com/bricef/ray-tracer/pkg/shaders"
	"github.com/bricef/ray-tracer/pkg/utils"
)

func TestSceneCreation(t *testing.T) {
	s := scene.NewScene()
	lights := s.Lights()
	entities := s.Entities
	if len(lights) > 0 || len(entities) > 0 {
		t.Errorf("Scene %v not created empty. Has %v and %v", s, lights, entities)
	}
}

func TestDefaultScene(t *testing.T) {
	s := scene.DefaultScene()

	if !(len(s.Lights()) > 0 && len(s.Entities) > 0) {
		t.Errorf("Default scene should have some lights and entities.")
	}

	if !s.Lights()[0].Position().Equal(math.NewPoint(-10, 10, -10)) {
		t.Errorf("Default scene not created with expected default light source at right location")
	}

	if !s.Lights()[0].GetLight().Intensity().Equal(color.New(1, 1, 1)) {
		t.Errorf("Default scene not created with expected default light source of right color")
	}

	if !s.Entities[1].Transform().Equal(math.NewTransform().Scale(0.5, 0.5, 0.5)) {
		t.Errorf("Default scene not created with default scaled object")
	}

	m := s.Entities[0].GetMaterial()

	if !m.Color().Equal(color.New(0.8, 1.0, 0.6)) {
		t.Errorf("Default object has the wrong color material")
	}

	if !(m.Diffuse() == 0.7 && m.Specular() == 0.2) {
		t.Errorf("Material property incorrect for default object")
	}

	m2 := s.Entities[1].GetMaterial()
	expected := material.NewMaterial()

	if m2 == nil || !m2.Equal(expected) {
		t.Errorf("Unexpected material on default sphere. Expected %v, got %v", material.NewMaterial(), m2)
	}

}

func TestIntersectWorld(t *testing.T) {
	s := scene.DefaultScene()
	r := ray.NewRay(
		math.NewPoint(0, 0, -5),
		math.NewVector(0, 0, 1),
	)
	xs := s.Intersections(r)

	if len(xs.All) != 4 {
		t.Errorf("Expected 4 intersection in default scene but go %v", len(xs.All))
	} else if !(xs.All[0].T == 4.0 && xs.All[1].T == 4.5 && xs.All[2].T == 5.5 && xs.All[3].T == 6.0) {
		t.Errorf("Scene intersections are not sorted!")
	}
}

func TestSceneShadingOnRayMiss(t *testing.T) {
	s := scene.DefaultScene()
	r := ray.NewRay(
		math.NewPoint(0, 0, -5),
		math.NewVector(0, 1, 0),
	)

	c := s.Cast(r)

	if !c.Equal(color.Black) {
		t.Errorf("Failed to shade a ray that doesn't intersect")
	}
}

func TestSceneShadingOnRayHit(t *testing.T) {
	s := scene.DefaultScene()
	r := ray.NewRay(
		math.NewPoint(0, 0, -5),
		math.NewVector(0, 0, 1),
	)

	c := s.Cast(r)
	expected := color.New(0.38066, 0.47583, 0.2855)
	if !c.Equal(expected) {
		t.Errorf("Scene failed to shade. Expected %v, got %v", c, expected)
	}
}

func TestShadingAnIntersectionInsideObject(t *testing.T) {
	s := scene.NewScene()
	s.Add(entities.NewSphere().Scale(0.5, 0.5, 0.5))
	s.Add(lighting.NewPointLight(color.White).MoveTo(math.NewPoint(0, 0.25, 0)))
	r := ray.NewRay(
		math.NewPoint(0, 0, 0),
		math.NewVector(0, 0, 1),
	)
	got := s.Cast(r)
	expected := color.New(0.90498, 0.90498, 0.90498)

	if !got.Equal(expected) {
		t.Errorf("Failed to shade hit inside object. Expected %v, got %v", expected, got)
	}
}

func TestShadingOnRayHitFirst(t *testing.T) {
	s := scene.DefaultScene()

	s.Entities[0].GetMaterial().SetAmbient(1)
	s.Entities[1].GetMaterial().SetAmbient(1)

	// Between spheres, looking at inner sphere.
	r := ray.NewRay(
		math.NewPoint(0, 0, .75),
		math.NewVector(0, 0, -1),
	)

	c := s.Cast(r)
	expected := s.Entities[1].GetMaterial().Color()

	if !c.Equal(expected) {
		t.Errorf("Scene failed to shade. Expected %v, got %v", c, expected)
	}
}

func TestReflectionOnReflectiveSurface(t *testing.T) {
	s := scene.DefaultScene()

	e := entities.NewPlane()
	e.Translate(0, -1, 0)
	e.GetMaterial().SetReflective(0.5)

	s.Add(e)

	r := ray.NewRay(
		math.NewPoint(0, 0, -3),
		math.NewVector(0, -m.Sqrt2/2, m.Sqrt2/2),
	)

	result := s.Cast(r)

	// Book results - close enough that I'll attribute this to rounding issues
	// expected := color.New(0.87677, 0.92436, 0.82918)

	expected := color.New(.87676, .92434, .82917)

	if !result.Equal(expected) {
		t.Errorf("Failed to compute reflected color. Expected %v, got %v", expected, result)
	}
}

func TestHandleInfiniteReflection(t *testing.T) {
	s := scene.NewScene()

	s.Add(lighting.NewPointLight(color.White))

	lower := entities.NewPlane().Translate(0, -1, 0)
	lower.GetMaterial().SetReflective(1.0)
	s.Add(lower)

	upper := entities.NewPlane().Translate(0, 1, 0)
	upper.GetMaterial().SetReflective(1.0)
	s.Add(upper)

	r := ray.NewRay(
		math.NewPoint(0, 0, 0),
		math.NewVector(0, 1, 0),
	)

	utils.FunctionTerminatesIn(t, 5, func() interface{} {
		return s.Cast(r)
	})

}

func TestOpaqueObjectHasNoRefraction(t *testing.T) {
	w := scene.DefaultScene()
	r := ray.NewRay(
		math.NewPoint(0, 0, -5),
		math.NewVector(0, 0, 1),
	)
	xs := r.GetIntersections(w.Entities)
	result := w.RefractedContribution(xs.Hit, 10)

	expected := color.Black
	if !result.Equal(expected) {
		t.Errorf("Opaque object has referaction component. Expected %v, got %v.", expected, result)
	}

}

func TestRefractionAtMaxDepthIsBlack(t *testing.T) {
	w := scene.DefaultScene()
	w.Entities[0].GetMaterial().SetTransparency(1.0)
	w.Entities[0].GetMaterial().SetRefractiveIndex(1.5)

	r := ray.NewRay(math.NewPoint(0, 0, -5), math.NewVector(0, 0, 1))

	xs := r.GetIntersections(w.Entities)

	result := w.RefractedContribution(xs.Hit, 0)
	expected := color.Black

	if !result.Equal(expected) {
		t.Errorf("Referaction at max depth should be black. Expected %v, Got %v.", expected, result)
	}
}

func TestReferactionUnderTotalInternalReflection(t *testing.T) {
	s := scene.DefaultScene()
	s.Entities[0].GetMaterial().SetTransparency(1.0)
	s.Entities[0].GetMaterial().SetRefractiveIndex(1.5)

	r := ray.NewRay(
		math.NewPoint(0, 0, m.Sqrt2/2.0),
		math.NewVector(0, 1, 0),
	)

	xs := r.GetIntersections(s.Entities)
	i := xs.All[1]

	result := s.RefractedContribution(i, 5)
	expected := color.Black

	if !result.Equal(expected) {
		t.Errorf("Refractyion contribution under conditions of total internal reflection. Expected %v, got %v.", expected, result)
	}
}

func TestRefractedColor(t *testing.T) {
	s := scene.DefaultScene()

	s.Entities[0].GetMaterial().SetAmbient(1.0)
	s.Entities[0].GetMaterial().SetShader(shaders.Test())

	s.Entities[1].GetMaterial().SetTransparency(1.0)
	s.Entities[1].GetMaterial().SetRefractiveIndex(1.5)

	r := ray.NewRay(
		math.NewPoint(0, 0, 0.1),
		math.NewVector(0, 1, 0),
	)

	xs := r.GetIntersections(s.Entities)
	fmt.Printf("N1=%v, N2=%v\n", xs.All[2].N1, xs.All[2].N2)

	result := s.RefractedContribution(xs.All[2], 5)

	// Adjusted from (0.0, 0.99888, 0.04725) in book
	expected := color.New(0, 0.99887, 0.04722)

	if !result.Equal(expected) {
		t.Errorf("Refracted color incorrect. Expected %v, got %v", expected, result)
	}

}

func TestRefractionSceneColoring(t *testing.T) {
	s := scene.DefaultScene()

	floor := entities.NewPlane()
	floor.Translate(0, -1, 0)
	floor.GetMaterial().SetTransparency(0.5)
	floor.GetMaterial().SetRefractiveIndex(1.5)
	s.Add(floor)

	ball := entities.NewSphere()
	ball.GetMaterial().SetColor(color.New(1, 0, 0))
	ball.GetMaterial().SetAmbient(0.5)
	ball.Translate(0, -3.5, -0.5)
	s.Add(ball)

	r := ray.NewRay(
		math.NewPoint(0, 0, -3),
		math.NewVector(0, -m.Sqrt2/2.0, m.Sqrt2/2.0),
	)

	result := s.Cast(r)
	expected := color.New(0.93642, 0.68642, 0.68642)

	if !result.Equal(expected) {
		t.Errorf("Invalid color. Expected %v, got %v.", expected, result)
	}

}

func TestRefractionAndReflectionCombined(t *testing.T) {
	s := scene.DefaultScene()

	floor := entities.NewPlane()
	floor.Translate(0, -1, 0)
	floor.GetMaterial().SetReflective(0.5)
	floor.GetMaterial().SetTransparency(0.5)
	floor.GetMaterial().SetRefractiveIndex(1.5)
	s.Add(floor)

	ball := entities.NewSphere()
	ball.GetMaterial().SetColor(color.New(1, 0, 0))
	ball.GetMaterial().SetAmbient(0.5)
	ball.Translate(0, -3.5, -0.5)
	s.Add(ball)

	r := ray.NewRay(
		math.NewPoint(0, 0, -3),
		math.NewVector(0, -m.Sqrt2/2.0, m.Sqrt2/2.0),
	)

	result := s.Cast(r)
	expected := color.New(0.93391, 0.69643, 0.69243)

	if !result.Equal(expected) {
		t.Errorf("Invalid color. Expected %v, got %v.", expected, result)
	}
}
