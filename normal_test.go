package triforce

import "testing"

func TestNormalZeroOut(t *testing.T) {
	var triangle *Triangle
	for tri := range ReadBinaryTriangleStream("binary.stl") {
		triangle = tri
		break
	}
	triangle.Normal.ZeroOut()
	if triangle.Normal.I != 0 {
		t.Error("Unable to zero out the normal for a given triangle.")
	}
	if triangle.Normal.J != 0 {
		t.Error("Unable to zero out the normal for a given triangle.")
	}
	if triangle.Normal.K != 0 {
		t.Error("Unable to zero out the normal for a given triangle.")
	}
}
