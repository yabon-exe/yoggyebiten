package model

type Rect[T int | uint | float32 | float64] struct {
	Left   T
	Right  T
	Top    T
	Bottom T
}

func (rect *Rect[T]) GetHW() (T, T) {
	return rect.Right - rect.Left, rect.Bottom - rect.Top
}

func Bounds[T int | uint | float32 | float64](CenterX T, CenterY T, Width T, Height T) Rect[T] {

	harfW := Width / 2.0
	harfH := Height / 2.0

	return Rect[T]{
		Left:   CenterX - harfW,
		Right:  CenterX + harfW,
		Top:    CenterY - harfH,
		Bottom: CenterY + harfH,
	}
}
