package mesh

import (
	"github.com/bricef/ray-tracer/quaternion"
)

type Sphere struct{}

func NewSphere() Mesh {
	return &Sphere{}
}

func (s *Sphere) Normal(p quaternion.Quaternion) quaternion.Quaternion {
	return p.Sub(quaternion.NewPoint(0, 0, 0)).Normalize()
}

func (s *Sphere) String() string {
	return "Mesh[Sphere]"
}
