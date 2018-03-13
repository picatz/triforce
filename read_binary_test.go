package triforce

import "testing"
import "fmt"

func TestReadBinaryHeader(t *testing.T) {
	header := ReadBinaryHeader("binary.stl")
	total := len(header.Data)
	want := 84
	if total != 84 {
		t.Errorf("Size of the header data is incorrect, got: %d, want: %d.", total, want)
	}
}

func TestReadBinaryTriangleStream(t *testing.T) {
	triangleStream := ReadBinaryTriangleStream("binary.stl")
	counter := 0
	for _ = range triangleStream {
		counter += 1
	}
	if counter != 12 {
		t.Error("Unable to read the triangles properly from a binary file.")
	}
}

func TestREAMDE(t *testing.T) {
	var maxX float32
	var maxY float32
	var maxZ float32
	for triangle := range ReadBinaryTriangleStream("binary.stl") {
		for _, vertex := range triangle.Vertices {
			if vertex.X > maxX {
				maxX = vertex.X
			}
			if vertex.Y > maxY {
				maxY = vertex.Y
			}
			if vertex.Z > maxZ {
				maxZ = vertex.Z
			}
		}
	}
	fmt.Println("max X value: ", maxX)
	fmt.Println("max Y value: ", maxY)
	fmt.Println("max Z value: ", maxZ)
}

func TestReadBinaryAllTriangles(t *testing.T) {
	triangles := ReadBinaryAllTriangles("binary.stl")
	if len(triangles) != 12 {
		t.Error("Unable to properly read all the triangles from a binary file into memory.")
	}
}
