package triforce

import (
	"encoding/binary"
	"math"
)

// A STL triangle ( also knows as a "facet" ) is made
// of three vertices, a normal and an attribute.
type Triangle struct {
	Normal    *Normal
	Vertices  *[3]Vertex
	Attribute uint16
}

// Calculate the area of a given STL triangle.
func (t *Triangle) Area() float32 {
	a := t.Vertices[0].DistanceBetween(&t.Vertices[1])
	b := t.Vertices[1].DistanceBetween(&t.Vertices[2])
	c := t.Vertices[2].DistanceBetween(&t.Vertices[0])
	s := (a + b + c) / 2
	A := math.Pow((s * (s - a) * (s - b) * (s - c)), float64(0.5))
	return float32(A)
}

// From a given slice of bytes, this will produce
// a reference to a triangle made from that data.
func TrianlgeFromBinaryChunk(chunk []byte) *Triangle {
	n := &Normal{}
	v := &[3]Vertex{}
	n.I = math.Float32frombits(binary.LittleEndian.Uint32(chunk[0:4]))
	n.J = math.Float32frombits(binary.LittleEndian.Uint32(chunk[4:8]))
	n.K = math.Float32frombits(binary.LittleEndian.Uint32(chunk[8:12]))
	v[0].X = math.Float32frombits(binary.LittleEndian.Uint32(chunk[12:16]))
	v[0].Y = math.Float32frombits(binary.LittleEndian.Uint32(chunk[16:20]))
	v[0].Z = math.Float32frombits(binary.LittleEndian.Uint32(chunk[20:24]))
	v[1].X = math.Float32frombits(binary.LittleEndian.Uint32(chunk[24:28]))
	v[1].Y = math.Float32frombits(binary.LittleEndian.Uint32(chunk[28:32]))
	v[1].Z = math.Float32frombits(binary.LittleEndian.Uint32(chunk[32:36]))
	v[2].X = math.Float32frombits(binary.LittleEndian.Uint32(chunk[36:40]))
	v[2].Y = math.Float32frombits(binary.LittleEndian.Uint32(chunk[40:44]))
	v[2].Z = math.Float32frombits(binary.LittleEndian.Uint32(chunk[44:48]))
	return &Triangle{Normal: n, Vertices: v, Attribute: binary.LittleEndian.Uint16(chunk[48:50])}
}
