package model

type Vertex struct {
	X float64
	Y float64
}

func NewVertex(x float64, y float64) Vertex {
	return Vertex{X: x, Y: y}
}
