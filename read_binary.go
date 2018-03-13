package triforce

import "os"
import "io"

func ReadBinaryHeader(path string) Header {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data := make([]byte, 84)
	count, err := f.Read(data)
	if err != nil {
		panic(err)
	}
	if count != 84 {
		panic("Incomplete binary header found for file (missing data).")
	}
	return Header{Data: data}
}

func ReadBinaryTriangleStream(path string) <-chan *Triangle {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	_, err = f.Seek(84, 0)
	if err != nil {
		panic(err)
	}
	stream := make(chan *Triangle)
	go func() {
		defer f.Close()
		defer close(stream)
		chunk := make([]byte, 50)
		for {
			_, err = f.Read(chunk)
			if err != nil {
				if err == io.EOF {
					break
				}
				// clean up
				f.Close()
				close(stream)
				panic(err)
			}
			stream <- TrianlgeFromBinaryChunk(chunk)
		}
	}()
	return stream
}

func ReadBinaryAllTriangles(path string) []*Triangle {
	triangles := []*Triangle{}
	for triangle := range ReadBinaryTriangleStream(path) {
		triangles = append(triangles, triangle)
	}
	return triangles
}

func MeshFromBinary(path string) Mesh {
	return Mesh{Name: path, Triangles: ReadBinaryAllTriangles(path)}
}
