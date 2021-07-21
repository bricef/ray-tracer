package ray

import "sort"

type Intersections struct {
	All []*Intersection
	Hit *Intersection
}

func NewIntersections() *Intersections {
	return &Intersections{}
}

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
