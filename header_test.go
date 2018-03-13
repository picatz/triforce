package triforce

import "testing"

func TestHeader(t *testing.T) {
	header := ReadBinaryHeader("binary.stl")
	triCount := header.TriangleCount()
	if int(triCount) != 12 {
		t.Error("Unable to successfully extract the triangle count for the binary file in the header section.")
	}
}
