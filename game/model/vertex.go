package model

type Vertex[T int | uint | float32 | float64] struct {
	X T
	Y T
}

func NewVertex[T int | uint | float32 | float64](x T, y T) Vertex[T] {
	return Vertex[T]{X: x, Y: y}
}

func (v *Vertex[T]) Set(x T, y T) {
	v.X = x
	v.Y = y
}

func (v *Vertex[T]) Get() (T, T) {
	return v.X, v.Y
}

const PI_FULL_CIRCLE = 360.0

type Circle struct {
	Vertex Vertex[float64]
	Rad    int
}
