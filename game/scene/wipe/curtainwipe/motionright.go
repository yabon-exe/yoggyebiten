package curtainwipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionRight struct {
}

func (motion *MotionRight) reset(rect *model.Rect[int], width int, height int) {
	rect.Left = 0
	rect.Right = 0
	rect.Top = 0
	rect.Bottom = height
}

func (motion *MotionRight) runClose(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxWidth) * speedRate
	rect.Right += int(speed)
	if rect.Right >= maxWidth {
		rect.Right = maxWidth
		return true
	}
	return false
}
func (motion *MotionRight) runOpen(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxWidth) * speedRate
	rect.Left += int(speed)
	if rect.Left >= maxWidth {
		rect.Left = maxWidth
		return true
	}
	return false
}
