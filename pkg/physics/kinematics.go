package physics

import (
	"github.com/bricef/ray-tracer/pkg/component"
	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/math"
)

type Kinematic struct {
	Velocity     math.Vector
	Acceleration math.Vector
}

func NewKinematic() core.Kinematic {
	return &Kinematic{}
}

func (c *Kinematic) Type() core.ComponentType {
	return component.Kinematic
}

func (c *Kinematic) SetAcceleration(a math.Vector) core.Kinematic {
	c.Acceleration = a
	return c
}

func (c *Kinematic) SetVelocity(a math.Vector) core.Kinematic {
	c.Velocity = a
	return c
}

func (c *Kinematic) Tick(owner core.Entity) {
	c.Velocity = c.Velocity.Add(c.Acceleration).AsVector()
	owner.Translate(c.Velocity.X(), c.Velocity.Y(), c.Velocity.Z())
}
