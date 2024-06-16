package cleanerwipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionRight struct {
}

func (motion *MotionRight) reset(cutoffRect *model.Rect[int], width int, height int) {
	cutoffRect.Left = 0
	cutoffRect.Right = width
	cutoffRect.Top = 0
	cutoffRect.Bottom = height
}

func (motion *MotionRight) runClean(cutoffRect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxWidth) * speedRate
	cutoffRect.Left += int(speed)
	if cutoffRect.Left >= maxWidth {
		cutoffRect.Left = maxWidth
		return true
	}
	return false
}
