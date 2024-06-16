package curtainwipe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

const MAX_RGB = 255

type CurtainWipe struct {
	isClosing   bool
	curtainRect model.Rect[int]
	Speed       int
	maxWidth    int
	maxHeight   int
}

func (w *CurtainWipe) Init() error {

	return nil
}
func (w *CurtainWipe) Reset(width int, height int) error {
	w.isClosing = false
	w.maxWidth = width
	w.maxHeight = height

	w.curtainRect.Left = 0
	w.curtainRect.Right = 0
	w.curtainRect.Top = 0
	w.curtainRect.Bottom = height

	return nil
}
func (w *CurtainWipe) Update() (bool, error) {

	wipeEnd := false

	if !w.isClosing {
		w.curtainRect.Right += w.Speed
		if w.curtainRect.Right >= w.maxWidth {
			w.curtainRect.Right = w.maxWidth
			w.isClosing = true
		}
	} else {
		w.curtainRect.Left += w.Speed
		if w.curtainRect.Left >= w.maxWidth {
			w.curtainRect.Left = w.maxWidth
			wipeEnd = true
		}
	}

	return wipeEnd, nil
}
func (w *CurtainWipe) Draw(screen *ebiten.Image, screenCapture *ebiten.Image) {
	if !w.isClosing {
		screen.DrawImage(screenCapture, nil)
	}

	curtain := ebiten.NewImage(w.curtainRect.GetHW())
	curtain.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(w.curtainRect.Left), float64(w.curtainRect.Top))
	screen.DrawImage(curtain, op)
}
