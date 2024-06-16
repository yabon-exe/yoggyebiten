package curtainwipe

import "github.com/yabon-exe/yoggyebiten/game/model"

type MotionRight struct {
}

func (motion *MotionRight) reset(rect *model.Rect[int], width int, height int) {
	rect.Left = width
	rect.Right = width
	rect.Top = 0
	rect.Bottom = height
}

func (motion *MotionRight) runClose(rect *model.Rect[int], speed int, maxWidth int, maxHeight int) bool {
	rect.Left -= speed
	if rect.Left <= 0 {
		rect.Left = 0
		return true
	}
	return false
}
func (motion *MotionRight) runOpen(rect *model.Rect[int], speed int, maxWidth int, maxHeight int) bool {
	rect.Right -= speed
	if rect.Right <= 0 {
		rect.Right = 0
		return true
	}
	return false
}
