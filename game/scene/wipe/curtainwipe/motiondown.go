package curtainwipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionDown struct {
}

func (motion *MotionDown) reset(rect *model.Rect[int], width int, height int) {
	rect.Left = 0
	rect.Right = width
	rect.Top = 0
	rect.Bottom = 0
}

func (motion *MotionDown) runClose(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxHeight) * speedRate
	rect.Bottom += int(speed)
	if rect.Bottom >= maxHeight {
		rect.Bottom = maxHeight
		return true
	}
	return false
}
func (motion *MotionDown) runOpen(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxHeight) * speedRate
	rect.Top += int(speed)
	if rect.Top >= maxHeight {
		rect.Top = maxHeight
		return true
	}
	return false
}
