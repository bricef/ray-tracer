package scene

import (
	"fmt"

	"github.com/bricef/ray-tracer/color"
	"github.com/bricef/ray-tracer/entity"
	"github.com/bricef/ray-tracer/light"
	"github.com/bricef/ray-tracer/material"
	q "github.com/bricef/ray-tracer/quaternion"
	"github.com/bricef/ray-tracer/ray"
	"github.com/bricef/ray-tracer/transform"
)

type Scene struct {
	Lights          []*light.PointLight
	Entities        []*entity.Entity
	BackgroundColor color.Color
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) Add(o interface{}) {
	switch t := o.(type) {
	case *entity.Entity:
		s.Entities = append(s.Entities, t)
	case *light.PointLight:
		s.Lights = append(s.Lights, t)
	default:
		fmt.Printf("Scene: Ignoring unknown object %v.", t)
	}

}

func DefaultScene() *Scene {
	s := NewScene()
	s.Add(
		light.NewPointLight(
			color.New(1, 1, 1),
			q.NewPoint(-10, 10, -10),
		),
	)
	s.Add(
		entity.NewSphere().SetMaterial(
			material.NewMaterial().
				SetColor(color.New(0.8, 1.0, 0.6)).
				SetDiffuse(0.7).
				SetSpecular(0.2),
		),
	)
	s.Add(
		entity.NewSphere().SetTransform(
			transform.NewTransform().Scale(0.5, 0.5, 0.5),
		),
	)
	return s
}

func (s *Scene) Intersections(r ray.Ray) ray.Intersections {
	xs := ray.Intersections{}
	for _, e := range s.Entities {
		newxs := r.Intersect(e)
		xs = xs.Merge(newxs)
	}
	return xs
}

func (s *Scene) Shade(r ray.Ray) color.Color {
	xs := s.Intersections(r)
	if xs.Hit != nil {
		return xs.Hit.ShadeAll(s.Lights)
	}
	return s.BackgroundColor
}
