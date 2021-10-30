package ray

import (
	"sort"
)

type Intersections struct {
	All []*Intersection
	Hit *Intersection
}

func NewIntersections() *Intersections {
	return &Intersections{}
}

// func (is *Intersections) All() []*Intersection {
// 	// Haiku: A code lament
// 	//
// 	// If Thing implements Interface,
// 	// Why []Thing not implement []Interface?
// 	// Memory layout! Autumn leaves fall sadly.
// 	//
// 	xs := make([]*Intersection, len(is.all))
// 	for i, x := range is.all {
// 		xs[i] = x
// 	}
// 	return xs
// }

// func (is *Intersections) Hit() *Intersection {
// 	return is.hit
// }

func (is *Intersections) Merge(xs *Intersections) *Intersections {
	// Short circuit the empty case.
	if (len(is.All) + len(xs.All)) == 0 {
		return &Intersections{
			All: []*Intersection{},
			Hit: nil,
		}
	}

	newAll := []*Intersection{}

	newAll = append(newAll, is.All...)
	newAll = append(newAll, xs.All...)

	sort.Slice(newAll, func(i, j int) bool {
		return newAll[i].T < newAll[j].T
	})

	var hit *Intersection
	for i, x := range newAll {
		if x.T > 0 && ((hit == nil) || x.T < hit.T) {
			hit = newAll[i]
		}
	}

	return &Intersections{
		All: newAll,
		Hit: hit,
	}
}
