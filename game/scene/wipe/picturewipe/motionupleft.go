package picturewipe

import (
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type MotionUpLeft struct {
}

func (motion *MotionUpLeft) runSlide(pictureLeftTop *model.Vertex[int], speedRate float64, maxWidth int, maxHeight int) bool {

	speedX := float64(maxWidth) * speedRate
	pictureLeftTop.X -= int(speedX)
	if pictureLeftTop.X <= -maxWidth {
		pictureLeftTop.X = -maxWidth
		return true
	}

	speedY := float64(maxHeight) * speedRate
	pictureLeftTop.Y -= int(speedY)
	if pictureLeftTop.Y <= -maxHeight {
		pictureLeftTop.Y = -maxHeight
		return true
	}
	return false
}
