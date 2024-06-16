package cleanerwipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionLeft struct {
}

func (motion *MotionLeft) reset(cutoffRect *model.Rect[int], width int, height int) {
	cutoffRect.Left = 0
	cutoffRect.Right = width
	cutoffRect.Top = 0
	cutoffRect.Bottom = height
}

func (motion *MotionLeft) runClean(cutoffRect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxWidth) * speedRate
	cutoffRect.Right -= int(speed)
	if cutoffRect.Right <= 0 {
		cutoffRect.Right = 0
		return true
	}
	return false
}
