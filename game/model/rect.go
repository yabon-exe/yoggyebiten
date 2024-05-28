package model

type Rect struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (rect *Rect) GetHW() (float64, float64) {
	return rect.Right - rect.Left, rect.Bottom - rect.Top
}

func Bounds(CenterX float64, CenterY float64, Width float64, Height float64) Rect {

	harfW := Width / 2.0
	harfH := Height / 2.0

	return Rect{
		Left:   CenterX - harfW,
		Right:  CenterX + harfW,
		Top:    CenterY - harfH,
		Bottom: CenterY + harfH,
	}
}
