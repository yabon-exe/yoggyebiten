package curtainwipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionLeft struct {
}

func (motion *MotionLeft) reset(rect *model.Rect[int], width int, height int) {
	rect.Left = width
	rect.Right = width
	rect.Top = 0
	rect.Bottom = height
}

func (motion *MotionLeft) runClose(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxWidth) * speedRate
	rect.Left -= int(speed)
	if rect.Left <= 0 {
		rect.Left = 0
		return true
	}
	return false
}
func (motion *MotionLeft) runOpen(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxWidth) * speedRate
	rect.Right -= int(speed)
	if rect.Right <= 0 {
		rect.Right = 0
		return true
	}
	return false
}
