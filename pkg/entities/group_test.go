package entities_test

import (
	m "math"
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

func TestEmptyGroupDoesNotHaveMesh(t *testing.T) {
	g := entities.NewGroup()
	// r := ray.NewRay(
	// 	math.NewPoint(0, 0, 0),
	// 	math.NewVector(0, 0, 1),
	// )
	m := g.GetMesh()
	if m != nil {
		t.Errorf("An empty group should not have any mesh.")
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
	s := entities.NewSphere().Translate(5, 0, 0)
	g := entities.NewGroup()
	g.AddChild(s)
	g.Scale(2, 2, 2)

	r := ray.NewRay(
		math.NewPoint(10, 0, -10),
		math.NewVector(0, 0, 1),
	)

	xs := r.GetIntersections([]core.Entity{g})

	if len(xs.All) != 2 {
		t.Errorf("Expected transformed group to apply transform to children. Got %v", xs)
	}
}

func TestConvertingWorldPointToObjectPoint(t *testing.T) {
	g1 := entities.NewGroup().RotateY(m.Pi / 2.0)
	g2 := entities.NewGroup().Scale(2, 2, 2)

	g1.AddChild(g2)

	s := entities.NewSphere().Translate(5, 0, 0)

	g2.AddChild(s)

	p := s.WorldPointToObjectPoint(math.NewPoint(-2, 0, -10))
	expected := math.NewPoint(0, 0, -1)

	if !p.Equal(expected) {
		t.Errorf("Failed to convert world point to object point. Expected %v, got %v.", expected, p)
	}

}

func TestConvertingWorldNormalToObjectNormal(t *testing.T) {
	g1 := entities.NewGroup().RotateY(m.Pi / 2.0)
	g2 := entities.NewGroup().Scale(1, 2, 3)
	g1.AddChild(g2)
	s := entities.NewSphere().Translate(5, 0, 0)
	g2.AddChild(s)

	k := m.Sqrt(3) / 3.0
	n := s.ObjectNormalToWorldNormal(math.NewVector(k, k, k))

	expected := math.NewVector(0.28571, 0.42857, -0.85714)

	if !n.Equal(expected) {
		t.Errorf("Failed to convert object normal to world normal. Expected %v, got %v", expected, n)
	}

}

func TestNormalOfChildInGroup(t *testing.T) {
	g1 := entities.NewGroup().RotateY(m.Pi / 2.0)
	g2 := entities.NewGroup().Scale(1, 2, 3)
	g1.AddChild(g2)
	s := entities.NewSphere().Translate(5, 0, 0)
	g2.AddChild(s)

	n := s.Normal(math.NewPoint(1.7321, 1.1547, -5.5774))

	expected := math.NewVector(0.28570, 0.42854, -0.85716)

	if !n.Equal(expected) {
		t.Errorf("Failed to get normal of entity in group. Expected %v, got %v", expected, n)
	}

}
