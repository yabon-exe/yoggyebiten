package curtainwipe

import "github.com/yabon-exe/yoggyebiten/game/model"

type MotionUp struct {
}

func (motion *MotionUp) reset(rect *model.Rect[int], width int, height int) {
	rect.Left = 0
	rect.Right = width
	rect.Top = height
	rect.Bottom = height
}

func (motion *MotionUp) runClose(rect *model.Rect[int], speed int, maxWidth int, maxHeight int) bool {
	rect.Top -= speed
	if rect.Top <= 0 {
		rect.Top = 0
		return true
	}
	return false
}
func (motion *MotionUp) runOpen(rect *model.Rect[int], speed int, maxWidth int, maxHeight int) bool {
	rect.Bottom -= speed
	if rect.Bottom <= 0 {
		rect.Bottom = 0
		return true
	}
	return false
}
