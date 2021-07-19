package math

type Quaternion interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
	Add(Quaternion) Quaternion
	Sub(Quaternion) Quaternion
	Divide(float64) Quaternion
	Scale(float64) Quaternion
	Equal(Quaternion) bool
	Negate() Quaternion
	AsVector() Vector
	AsPoint() Point
	IsVector() bool
	IsPoint() bool
}

type Point interface {
	Quaternion
}

type Vector interface {
	Quaternion
	Magnitude() float64
	Normalize() Vector
	Cross(Vector) Vector
	Dot(Vector) float64
	Reflect(Vector) Vector
	Invert() Vector
}

type Transform interface {
	GetMatrix() Matrix
	Equal(other Transform) bool
	Raw(rawMatrix [][]float64) Transform
	Translate(x float64, y float64, z float64) Transform
	Apply(Quaternion) Quaternion
	Inverse() Transform
	Transpose() Transform
	Scale(x float64, y float64, z float64) Transform
	RotateX(r float64) Transform
	RotateY(r float64) Transform
	RotateZ(r float64) Transform
	ReflectX() Transform
	ReflectY() Transform
	ReflectZ() Transform
	Shear(xy, xz, yx, yz, zx, zy float64) Transform
	MoveTo(Point) Transform
	Position() Point
}
