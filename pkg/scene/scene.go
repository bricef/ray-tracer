package scene

import (
	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/lighting"
	"github.com/bricef/ray-tracer/pkg/material"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
)

type Scene struct {
	lights          []core.Entity
	Entities        []core.Entity
	BackgroundColor color.Color
}

func (s *Scene) Lights() []core.Entity {
	return s.lights
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) Add(o core.Entity) {
	if l := o.GetLight(); l != nil {
		s.lights = append(s.lights, o)
	} else {
		s.Entities = append(s.Entities, o)
	}

}

func DefaultScene() *Scene {
	s := NewScene()
	s.Add(
		lighting.NewPointLight(
			color.New(1, 1, 1),
		).Translate(-10, 10, -10),
	)
	s.Add(
		entities.NewSphere().AddComponent(
			material.NewMaterial().
				SetColor(color.New(0.8, 1.0, 0.6)).
				SetDiffuse(0.7).
				SetSpecular(0.2),
		),
	)
	s.Add(
		entities.NewSphere().Scale(0.5, 0.5, 0.5),
	)
	return s
}

func (s *Scene) Intersections(r ray.Ray) ray.Intersections {
	xs := ray.Intersections{}
	for _, e := range s.Entities {
		mat := e.GetMaterial()
		mesh := e.GetMesh()
		if mat != nil && mesh != nil { // Ignore entities without mesh or material
			xs = xs.Merge(r.Intersect(e))
		}
	}
	return xs
}

func (s *Scene) Shade(r ray.Ray) color.Color {
	xs := s.Intersections(r)
	if xs.Hit != nil {
		// fmt.Printf("Hit\n")
		return xs.Hit.ShadeAll(s.Lights())
	}
	return s.BackgroundColor
}

func (s *Scene) Obstructed(a math.Point, b math.Point) bool {
	path := a.Sub(b).AsVector()
	distance := path.Magnitude()
	direction := path.Normalize()
	r := ray.NewRay(b, direction)
	xs := s.Intersections(r)
	if xs.Hit != nil && xs.Hit.T <= distance {
		return true
	}
	return false
}

func (s *Scene) Cast(r ray.Ray) color.Color {
	c := color.New(0, 0, 0)
	xs := s.Intersections(r)
	if xs.Hit != nil {
		mat := xs.Hit.Entity.GetMaterial()
		for _, l := range s.lights {
			if s.Obstructed(xs.Hit.OverPoint, l.Position()) {
				contribution := lighting.PhongShadow(mat, l, xs.Hit.Point, xs.Hit.EyeVector, xs.Hit.Normal)
				c = c.Add(contribution)
			} else {
				contribution := lighting.Phong(mat, l, xs.Hit.Point, xs.Hit.EyeVector, xs.Hit.Normal)
				c = c.Add(contribution)
			}

		}
		return c
	}
	return s.BackgroundColor
}

func (s *Scene) Tick() *Scene {
	for _, e := range s.Entities {
		e.Tick(s.Entities)
	}
	return s
}
