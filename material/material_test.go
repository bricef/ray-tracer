package material

import "testing"

func TestDefaultMaterial(t *testing.T) {
	m := NewMaterial()

	if !(m.Ambient == 0.1 && m.Diffuse == 0.9 && m.Specular == 0.9 && m.Shininess == 200.0) {
		t.Errorf("Failed to set up default material %v", m)
	}

}
