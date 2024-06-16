package picturewipe

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type Direct int

const (
	MOTION_LEFT Direct = iota
	MOTION_RIGHT
	MOTION_UP
	MOTION_DOWN
)

type PictureWipeMotion interface {
	runSlide(pictureLeftTop *model.Vertex[int], speedRate float64, maxWidth int, maxHeight int) bool
}

type PictureWipe struct {
	pictureLeftTop *model.Vertex[int]
	maxWidth       int
	maxHeight      int
	motion         PictureWipeMotion
	SpeedRate      float64
	Direct         Direct
}

func (w *PictureWipe) Init() error {

	switch w.Direct {
	case MOTION_LEFT:
		w.motion = &MotionLeft{}
	case MOTION_RIGHT:
		w.motion = &MotionRight{}
	case MOTION_UP:
		w.motion = &MotionUp{}
	case MOTION_DOWN:
		w.motion = &MotionDown{}
	default:
		w.motion = &MotionLeft{}
	}

	w.pictureLeftTop = &model.Vertex[int]{}

	return nil
}
func (w *PictureWipe) Reset(width int, height int) error {
	w.maxWidth = width
	w.maxHeight = height
	w.pictureLeftTop.X = 0
	w.pictureLeftTop.Y = 0

	return nil
}
func (w *PictureWipe) Update() (bool, error) {

	wipeEnd := false
	closeEnd := w.motion.runSlide(w.pictureLeftTop, w.SpeedRate, w.maxWidth, w.maxHeight)
	if closeEnd {
		wipeEnd = true
	}
	return wipeEnd, nil
}
func (w *PictureWipe) Draw(screen *ebiten.Image, screenCapture *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(w.pictureLeftTop.X), float64(w.pictureLeftTop.Y))
	screen.DrawImage(screenCapture, op)
}
