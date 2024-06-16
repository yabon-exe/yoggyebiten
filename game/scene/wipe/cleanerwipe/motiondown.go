package cleanerwipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionDown struct {
}

func (motion *MotionDown) reset(cutoffRect *model.Rect[int], width int, height int) {
	cutoffRect.Left = 0
	cutoffRect.Right = width
	cutoffRect.Top = 0
	cutoffRect.Bottom = height
}

func (motion *MotionDown) runClean(cutoffRect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxHeight) * speedRate
	cutoffRect.Top += int(speed)
	if cutoffRect.Top >= maxHeight {
		cutoffRect.Top = maxHeight
		return true
	}
	return false
}
