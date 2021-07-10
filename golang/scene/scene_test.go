package scene

import (
	"testing"

	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/material"
	q "github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/ray"
	"github.com/bricef/ray-tracer/transform"
)

func TestSceneCreation(t *testing.T) {
	s := NewScene()
	lights := s.Lights
	entities := s.Entities
	if len(lights) > 0 || len(entities) > 0 {
		t.Errorf("Scene %v not created empty. Has %v and %v", s, lights, entities)
	}
}

func TestDefaultScene(t *testing.T) {
	s := DefaultScene()

	if !(len(s.Lights) > 0 && len(s.Entities) > 0) {
		t.Errorf("Default scene should have some lights and entities.")
	}

	if !s.Lights[0].Position.Equal(q.NewPoint(-10, 10, -10)) {
		t.Errorf("Default scene not created with expected default light source at right location")
	}

	if !s.Lights[0].Intensity.Equal(color.New(1, 1, 1)) {
		t.Errorf("Default scene not created with expected default light source of right color")
	}

	if !s.Entities[1].Transform.Equal(transform.NewTransform().Scale(0.5, 0.5, 0.5)) {
		t.Errorf("Default scene not created with default scaled object")
	}

	m := s.Entities[0].Material

	if !m.Color.Equal(color.New(0.8, 1.0, 0.6)) {
		t.Errorf("Default object has the wrong color material")
	}

	if !(m.Diffuse == 0.7 && m.Specular == 0.2) {
		t.Errorf("Material property incorrect for default object")
	}

	m2 := s.Entities[1].Material

	if !m2.Equal(material.NewMaterial()) {
		t.Errorf("Unexpected material on default sphere. Expected %v, got %v", material.NewMaterial(), m2)
	}

}

func TestIntersectWorld(t *testing.T) {
	s := DefaultScene()
	r := ray.NewRay(
		q.NewPoint(0, 0, -5),
		q.NewVector(0, 0, 1),
	)
	xs := s.Intersections(r)

	if len(xs.All) != 4 {
		t.Errorf("Expected 4 intersection in default scene but go %v", len(xs.All))
	} else if !(xs.All[0].T == 4.0 && xs.All[1].T == 4.5 && xs.All[2].T == 5.5 && xs.All[3].T == 6.0) {
		t.Errorf("Scene intersections are not sorted!")
	}
}

func TestSceneShadingOnRayMiss(t *testing.T) {
	s := DefaultScene()
	r := ray.NewRay(
		q.NewPoint(0, 0, -5),
		q.NewVector(0, 1, 0),
	)

	c := s.Shade(r)

	if !c.Equal(color.Black) {
		t.Errorf("Failed to shade a ray that doesn't intersect")
	}
}

func TestSceneShadingOnRayHit(t *testing.T) {
	s := DefaultScene()
	r := ray.NewRay(
		q.NewPoint(0, 0, -5),
		q.NewVector(0, 0, 1),
	)

	c := s.Shade(r)
	expected := color.New(0.38066, 0.47583, 0.2855)
	if !c.Equal(expected) {
		t.Errorf("Scene failed to shade. Expected %v, got %v", c, expected)
	}
}

func TestShadingOnRayHitFirst(t *testing.T) {
	s := DefaultScene()

	s.Entities[0].Material.SetAmbient(1)
	s.Entities[1].Material.SetAmbient(1)

	// Between spheres, looking at inner sphere.
	r := ray.NewRay(
		q.NewPoint(0, 0, .75),
		q.NewVector(0, 0, -1),
	)

	c := s.Shade(r)
	expected := s.Entities[1].Material.Color

	if !c.Equal(expected) {
		t.Errorf("Scene failed to shade. Expected %v, got %v", c, expected)
	}
}
