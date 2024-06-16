package cleanerwipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionUp struct {
}

func (motion *MotionUp) reset(cutoffRect *model.Rect[int], width int, height int) {
	cutoffRect.Left = 0
	cutoffRect.Right = width
	cutoffRect.Top = 0
	cutoffRect.Bottom = height
}

func (motion *MotionUp) runClean(cutoffRect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxHeight) * speedRate
	cutoffRect.Bottom -= int(speed)
	if cutoffRect.Bottom <= 0 {
		cutoffRect.Bottom = 0
		return true
	}
	return false
}
