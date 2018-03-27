package triforce

type Mesh struct {
	Name      string
	Triangles []*Triangle
}

func (m *Mesh) Area() float32 {
	var area float32
	for _, triangle := range m.Triangles {
		area += triangle.Area()
	}
	return area
}

func (m *Mesh) MaximumX() float32 {
	var max float32
	for _, triangle := range m.Triangles {
		for _, vertex := range triangle.Vertices {
			if vertex.X > max {
				max = vertex.X
			}
		}
	}
	return max
}

func (m *Mesh) MinimumX() float32 {
	var min float32
	for _, triangle := range m.Triangles {
		for _, vertex := range triangle.Vertices {
			if vertex.X < min {
				min = vertex.X
			}
		}
	}
	return min
}

func (m *Mesh) MaximumY() float32 {
	var max float32
	for _, triangle := range m.Triangles {
		for _, vertex := range triangle.Vertices {
			if vertex.Y > max {
				max = vertex.Y
			}
		}
	}
	return max
}

func (m *Mesh) MinimumY() float32 {
	var min float32
	for _, triangle := range m.Triangles {
		for _, vertex := range triangle.Vertices {
			if vertex.Y < min {
				min = vertex.Y
			}
		}
	}
	return min
}

func (m *Mesh) MaximumZ() float32 {
	var max float32
	for _, triangle := range m.Triangles {
		for _, vertex := range triangle.Vertices {
			if vertex.Z > max {
				max = vertex.Z
			}
		}
	}
	return max
}

func (m *Mesh) MinimumZ() float32 {
	var min float32
	for _, triangle := range m.Triangles {
		for _, vertex := range triangle.Vertices {
			if vertex.Z < min {
				min = vertex.Z
			}
		}
	}
	return min
}
