package curtainwipe

import "github.com/yabon-exe/yoggyebiten/game/model"

type MotionLeft struct {
}

func (motion *MotionLeft) reset(rect *model.Rect[int], width int, height int) {
	rect.Left = 0
	rect.Right = 0
	rect.Top = 0
	rect.Bottom = height
}

func (motion *MotionLeft) runClose(rect *model.Rect[int], speed int, maxWidth int, maxHeight int) bool {
	rect.Right += speed
	if rect.Right >= maxWidth {
		rect.Right = maxWidth
		return true
	}
	return false
}
func (motion *MotionLeft) runOpen(rect *model.Rect[int], speed int, maxWidth int, maxHeight int) bool {
	rect.Left += speed
	if rect.Left >= maxWidth {
		rect.Left = maxWidth
		return true
	}
	return false
}
