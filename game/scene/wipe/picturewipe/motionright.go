package picturewipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionRight struct {
}

func (motion *MotionRight) runSlide(pictureLeftTop *model.Vertex[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxWidth) * speedRate
	pictureLeftTop.X += int(speed)
	if pictureLeftTop.X >= maxWidth {
		pictureLeftTop.X = maxWidth
		return true
	}
	return false
}
