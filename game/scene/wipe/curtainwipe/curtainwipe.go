package curtainwipe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

const (
	MAX_RGB = 255
)

type Direct int

const (
	MOTION_LEFT Direct = iota
	MOTION_RIGHT
	MOTION_UP
	MOTION_DOWN
)

type CurtainWipeMotion interface {
	reset(rect *model.Rect[int], width int, height int)
	runClose(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool
	runOpen(rect *model.Rect[int], speedRate float64, maxWidth int, maxHeight int) bool
}

type CurtainWipe struct {
	isClosing   bool
	curtainRect *model.Rect[int]
	maxWidth    int
	maxHeight   int
	motion      CurtainWipeMotion
	SpeedRate   float64
	Direct      Direct
}

func (w *CurtainWipe) Init() error {

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

	w.curtainRect = &model.Rect[int]{}

	return nil
}
func (w *CurtainWipe) Reset(width int, height int) error {
	w.isClosing = true
	w.maxWidth = width
	w.maxHeight = height
	w.motion.reset(w.curtainRect, width, height)

	return nil
}
func (w *CurtainWipe) Update() (bool, error) {

	wipeEnd := false

	if w.isClosing {
		closeEnd := w.motion.runClose(w.curtainRect, w.SpeedRate, w.maxWidth, w.maxHeight)
		if closeEnd {
			w.isClosing = false
		}
	} else {
		oepnEnd := w.motion.runOpen(w.curtainRect, w.SpeedRate, w.maxWidth, w.maxHeight)
		if oepnEnd {
			wipeEnd = true
		}
	}

	return wipeEnd, nil
}
func (w *CurtainWipe) Draw(screen *ebiten.Image, screenCapture *ebiten.Image) {
	if w.isClosing {
		screen.DrawImage(screenCapture, nil)
	}

	curtain := ebiten.NewImage(w.curtainRect.GetHW())
	curtain.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(w.curtainRect.Left), float64(w.curtainRect.Top))
	screen.DrawImage(curtain, op)
}
