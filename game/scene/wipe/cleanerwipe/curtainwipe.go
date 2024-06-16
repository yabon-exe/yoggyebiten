package cleanerwipe

import (
	"image"

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

type CleanerWipeMotion interface {
	reset(cutoffRect *model.Rect[int], width int, height int)
	runClean(cutoffRect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool
}

type CleanerWipe struct {
	cutoffRect *model.Rect[int]
	maxWidth   int
	maxHeight  int
	motion     CleanerWipeMotion
	SpeedRate  float64
	Direct     Direct
}

func (w *CleanerWipe) Init() error {

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

	w.cutoffRect = &model.Rect[int]{}

	return nil
}
func (w *CleanerWipe) Reset(width int, height int) error {
	w.maxWidth = width
	w.maxHeight = height
	w.motion.reset(w.cutoffRect, width, height)

	return nil
}
func (w *CleanerWipe) Update() (bool, error) {

	wipeEnd := false
	closeEnd := w.motion.runClean(w.cutoffRect, w.SpeedRate, w.maxWidth, w.maxHeight)
	if closeEnd {
		wipeEnd = true
	}
	return wipeEnd, nil
}
func (w *CleanerWipe) Draw(screen *ebiten.Image, screenCapture *ebiten.Image) {
	rect := image.Rect(w.cutoffRect.Left, w.cutoffRect.Top, w.cutoffRect.Right, w.cutoffRect.Bottom)
	subCapture := screenCapture.SubImage(rect).(*ebiten.Image)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(w.cutoffRect.Left), float64(w.cutoffRect.Top))
	screen.DrawImage(subCapture, op)
}
