package raytracer

import (
	m "github.com/bricef/ray-tracer/matrix"
	q "github.com/bricef/ray-tracer/quaternion"
	t "github.com/bricef/ray-tracer/transform"
)

func Vector(x float64, y float64, z float64) q.Vector {
	return q.NewVector(x, y, z)
}

func Point(x, y, z float64) q.Point {
	return q.NewPoint(x, y, z)
}

func Matrix(values [][]float64) m.Matrix {
	return m.New(values)
}

func Quaternion(x, y, z, w float64) q.Quaternion {
	return q.New(x, y, z, w)
}

func Transform() t.Transform {
	return t.New()
}
