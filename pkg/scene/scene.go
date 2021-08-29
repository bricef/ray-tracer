package scene

import (
	"fmt"

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

func (s *Scene) Intersections(r ray.Ray) *ray.Intersections {
	return r.GetIntersections(s.Entities)
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
	return s.LimitedCast(r, 5)
}

func (s *Scene) LimitedCast(r ray.Ray, depth int) color.Color {
	if depth <= 0 { //Abort recursion after depth reached.
		return color.Black
	}

	c := color.New(0, 0, 0)
	xs := s.Intersections(r)
	if xs.Hit != nil {

		// Get lighting contributions
		c = c.Add(s.LightingContribution(xs.Hit, depth))

		// Get reflected contributions
		c = c.Add(s.ReflectedContribution(xs.Hit, depth))

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

func (s *Scene) LightingContribution(hit *ray.Intersection, depth int) color.Color {
	c := color.New(0, 0, 0)
	for _, l := range s.lights {
		c = c.Add(s.LightContribution(l, hit))
	}
	return c
}

func (s *Scene) LightContribution(l core.Entity, hit *ray.Intersection) color.Color {
	mat := hit.Entity.GetMaterial()
	if s.Obstructed(hit.OverPoint, l.Position()) {
		return lighting.PhongShadow(mat, l, hit.OverPoint, hit.EyeVector, hit.Normal)
	} else {
		return lighting.Phong(mat, l, hit.OverPoint, hit.EyeVector, hit.Normal)
	}
}

func (s *Scene) ReflectedContribution(i *ray.Intersection, depth int) color.Color {
	mat := i.Entity.GetMaterial()
	if mat == nil { // No material
		return color.Black
	}

	if mat.Reflective() == 0.0 { // Not reflective
		return color.Black
	}

	r := ray.NewRay(
		i.OverPoint,
		i.ReflectVector,
	)
	return s.LimitedCast(r, depth-1).Scale(mat.Reflective())

}

func (s *Scene) RefractedContribution(i *ray.Intersection, depth int) color.Color {
	mat := i.Entity.GetMaterial()
	// Max depth, no refraction
	if depth <= 0 {
		return color.Black
	}

	// No material, no refraction
	if mat == nil {
		return color.Black
	}

	// Material is opaque, no refraction
	if mat.Transparency() == 0.0 {
		return color.Black
	}

	n_ratio := i.N1 / i.N2
	cos_i := i.EyeVector.Dot(i.Normal)
	sin2_t := n_ratio * n_ratio * (1 - (cos_i * cos_i))
	if sin2_t > 1.0 {
		fmt.Printf("Total Internal Relfection. ration = %v, Sin2_t = %v\n", n_ratio, sin2_t)
		return color.Black
	}

	return color.White
}
