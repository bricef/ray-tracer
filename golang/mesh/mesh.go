package mesh

import "github.com/bricef/ray-tracer/quaternion"

type Mesh interface {
	Normal(quaternion.Quaternion) quaternion.Quaternion
}
