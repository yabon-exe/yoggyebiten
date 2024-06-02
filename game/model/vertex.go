package model

type Vertex struct {
	X float64
	Y float64
}

func NewVertex(x float64, y float64) Vertex {
	return Vertex{X: x, Y: y}
}

func (v *Vertex) Set(x float64, y float64) {
	v.X = x
	v.Y = y
}

const PI_FULL_CIRCLE = 360.0

type Circle struct {
	Vertex Vertex
	Rad    int
}
