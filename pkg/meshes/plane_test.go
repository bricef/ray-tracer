package meshes_test

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/entity"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/meshes"
	"github.com/bricef/ray-tracer/pkg/ray"
)

func TestPlaneNormals(t *testing.T) {
	p := meshes.PlaneMesh()

	type Case struct {
		Result   math.Vector
		Expected math.Vector
	}

	cases := []Case{
		{
			p.Normal(math.NewPoint(0, 0, 0)),
			math.NewVector(0, 1, 0),
		},
		{
			p.Normal(math.NewPoint(10, 0, -10)),
			math.NewVector(0, 1, 0),
		},
		{
			p.Normal(math.NewPoint(-5, 0, 150)),
			math.NewVector(0, 1, 0),
		},
	}

	for _, c := range cases {
		if !c.Result.Equal(c.Expected) {
			t.Errorf("Plane normal miscalculated. Expected %v, got %v", c.Expected, c.Result)
		}
	}

}

func TestPlaneIntersectionsParallel(t *testing.T) {
	p := entity.NewEntity().AddComponent(meshes.PlaneMesh())

	r := ray.NewRay(
		math.NewPoint(0, 10, 0),
		math.NewVector(0, 0, 1),
	)

	xs := r.Intersect(p)

	if len(xs.All) != 0 {
		t.Errorf("plane Expected no interscetion with parallel ray")
	}
}
func TestPlaneIntersectionsCoplanar(t *testing.T) {
	p := entity.NewEntity().AddComponent(meshes.PlaneMesh())

	r := ray.NewRay(
		math.NewPoint(0, 0, 0),
		math.NewVector(0, 0, 1),
	)

	xs := r.Intersect(p)

	if len(xs.All) != 0 {
		t.Errorf("plane Expected no interscetion with coplanar ray. Got %v", len(xs.All))
	}
}

func TestPlaneIntersectionsFromAbove(t *testing.T) {
	p := entity.NewEntity().AddComponent(meshes.PlaneMesh())

	r := ray.NewRay(
		math.NewPoint(0, 1, 0),
		math.NewVector(0, -1, 0),
	)

	xs := r.Intersect(p)

	if len(xs.All) != 1 || xs.Hit.T != 1.0 || xs.Hit.Entity != p {
		t.Errorf("plane expected one intersection from above. Got %v", xs)
	}
}

func TestPlaneIntersectionsFromBelow(t *testing.T) {
	p := entity.NewEntity().AddComponent(meshes.PlaneMesh())

	r := ray.NewRay(
		math.NewPoint(0, -1, 0),
		math.NewVector(0, 1, 0),
	)

	xs := r.Intersect(p)

	if len(xs.All) != 1 || xs.Hit.T != 1.0 || xs.Hit.Entity != p {
		t.Errorf("plane expected one intersection from above. Got %v", xs)
	}
}
