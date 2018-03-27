package triforce

import "testing"

func TestMeshFromBinary(t *testing.T) {
	m := MeshFromBinary("binary.stl")
	if m.Name != "binary.stl" {
		t.Error("Unable to load mesh from binary.")
	}
	if len(m.Triangles) != 12 {
		t.Error("Unable to load mesh from binary.")
	}
}

func TestMeshMaximums(t *testing.T) {
	m := MeshFromBinary("binary.stl")
	if m.MaximumX() != 0.5 {
		t.Error("Unable to find maximum X from binary file.")
	}
	if m.MaximumY() != 0.5 {
		t.Error("Unable to find maximum Y from binary file.")
	}
	if m.MaximumZ() != 1 {
		t.Error("Unable to find maximum Z from binary file.")
	}
}

func TestMeshMinimums(t *testing.T) {
	m := MeshFromBinary("binary.stl")
	if m.MinimumX() != -0.5 {
		t.Error("Unable to find maximum X from binary file.")
	}
	if m.MinimumY() != -0.5 {
		t.Error("Unable to find maximum Y from binary file.")
	}
	if m.MinimumZ() != 0 {
		t.Error("Unable to find maximum Z from binary file.")
	}
}
