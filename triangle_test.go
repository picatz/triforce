package triforce

import "testing"
import "os"
import "fmt"

func TestTrianlgeFromBinaryChunk(t *testing.T) {
	f, err := os.Open("binary.stl")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data := make([]byte, 50)
	count, err := f.ReadAt(data, 84)
	if err != nil {
		panic(err)
	}
	if count != 50 {
		panic("Incomplete binary chunk (missing data).")
	}
	triangle := TrianlgeFromBinaryChunk(data)
	if triangle.Normal.I != 0 && triangle.Normal.J != 0 && triangle.Normal.K != 1 {
		t.Error("Unable to create a triangle with proper normals parsed from chunk")
	}
}

func TestTriangleArea(t *testing.T) {
	f, err := os.Open("binary.stl")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data := make([]byte, 50)
	count, err := f.ReadAt(data, 84)
	if err != nil {
		panic(err)
	}
	if count != 50 {
		panic("Incomplete binary chunk (missing data).")
	}
	triangle := TrianlgeFromBinaryChunk(data)
	fmt.Println(triangle.Area())
}
