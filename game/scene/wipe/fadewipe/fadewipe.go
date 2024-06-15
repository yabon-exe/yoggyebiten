package fadewipe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type FadeInOutWipe struct {
	rgb       int
	Speed     int
	isClosing bool
}

func (w *FadeInOutWipe) Init() error {

	return nil
}
func (w *FadeInOutWipe) Reset(width int, height int) error {
	w.rgb = 0
	w.isClosing = false
	return nil
}
func (w *FadeInOutWipe) Update() (bool, error) {

	wipeEnd := false
	if !w.isClosing {
		w.rgb += w.Speed
	} else {
		w.rgb -= w.Speed
		if w.rgb <= 0 {
			wipeEnd = true
		}
	}
	return wipeEnd, nil
}
func (w *FadeInOutWipe) Draw(screen *ebiten.Image, screenCapture *ebiten.Image) {
	if !w.isClosing {
		screen.DrawImage(screenCapture, nil)
	}
	rgb := uint8(w.rgb)
	screen.Fill(color.RGBA{rgb, rgb, rgb, 0})
}
