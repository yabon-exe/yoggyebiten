package picturewipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionUp struct {
}

func (motion *MotionUp) runSlide(pictureLeftTop *model.Vertex[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speed := float64(maxHeight) * speedRate
	pictureLeftTop.Y -= int(speed)
	if pictureLeftTop.Y <= -maxHeight {
		pictureLeftTop.Y = -maxHeight
		return true
	}
	return false
}
