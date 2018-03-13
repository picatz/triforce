package triforce

import "math"

type Vertex struct {
	X float32
	Y float32
	Z float32
}

func (v1 *Vertex) DistanceBetween(v2 *Vertex) float64 {
	x := math.Pow((float64(v1.X) - float64(v2.X)), 2)
	y := math.Pow((float64(v1.Y) - float64(v2.Y)), 2)
	z := math.Pow((float64(v1.Z) - float64(v2.Z)), 2)
	a := x + y + z
	return math.Pow(a, float64(0.5))
}
