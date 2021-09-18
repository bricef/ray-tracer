package entities_test

import (
	"testing"

	"github.com/bricef/ray-tracer/pkg/core"
	"github.com/bricef/ray-tracer/pkg/entities"
	"github.com/bricef/ray-tracer/pkg/entity"
	"github.com/bricef/ray-tracer/pkg/math"
	"github.com/bricef/ray-tracer/pkg/ray"
)

func TestGroupInitialisation(t *testing.T) {
	g := entities.NewGroup()
	result := g.Transform()

	expected := math.NewTransform()

	if !result.Equal(expected) {
		t.Errorf("Group did not initialise with indenty transform. Expected %v, got %v.", expected, t)
	}
}

func TestGroupHasParent(t *testing.T) {
	g := entities.NewGroup()
	result := g.Parent()
	if result != nil {
		t.Errorf("Expect entity default parent set to be empty.")
	}
}

func TestGroupCanHaveChild(t *testing.T) {
	g := entities.NewGroup()
	e := entity.NewEntity()
	g.AddChild(e)

	if g.Children()[0] != e || e.Parent() != g {
		t.Errorf("Failed to add a child entity to a group.")
	}
}

func TestGroupConstructionWithChildren(t *testing.T) {
	e := entity.NewEntity()
	g := entities.NewGroup(e)

	if g.Children()[0] != e || e.Parent() != g {
		t.Errorf("Failed to add a child entity to a group.")
	}
}

func TestEmptyGroupHasEmptyIntersections(t *testing.T) {
	g := entities.NewGroup()
	r := ray.NewRay(
		math.NewPoint(0, 0, 0),
		math.NewVector(0, 0, 1),
	)
	ts := g.GetMesh().Intersect(r)
	if len(ts) != 0 {
		t.Errorf("An empty group should have not intersection.")
	}

}

func TestGroupWithChildrenShouldHaveChildrenIntersect(t *testing.T) {
	e1 := entities.NewSphere()
	e2 := entities.NewSphere().Translate(0, 0, -3)
	e3 := entities.NewSphere().Translate(5, 0, 0)
	g := entities.NewGroup(e1, e2, e3)
	r := ray.NewRay(
		math.NewPoint(0, 0, -5),
		math.NewVector(0, 0, 1),
	)

	xs := r.GetIntersections([]core.Entity{g})

	if len(xs.All) != 4 {
		t.Errorf("Unexpected number of intersections for test group. Expected 4, got %v", len(xs.All))
	}

	type Case struct {
		expected core.Entity
		result   core.Entity
	}
	tests := []Case{
		{e2, xs.All[0].Entity},
		{e2, xs.All[1].Entity},
		{e1, xs.All[2].Entity},
		{e1, xs.All[3].Entity},
	}

	for i, test := range tests {
		if test.result != test.expected {
			t.Errorf("Expected group intersection for index %v: %v. Got %v instead", i, test.expected, test.result)
		}
	}
}

func TestGroupTransformOnIntersect(t *testing.T) {
	g := entities.NewGroup()
	g.Transform().Scale(2, 2, 2)

	s := entities.NewSphere().Translate(5, 0, 0)
	g.AddChild(s)

	r := ray.NewRay(
		math.NewPoint(10, 0, -10),
		math.NewVector(0, 0, 1),
	)

	xs := r.GetIntersections([]core.Entity{g})

	if len(xs.All) != 2 {
		t.Errorf("Expected transformed group to apply transform to children. Got %v", xs)
	}
}
