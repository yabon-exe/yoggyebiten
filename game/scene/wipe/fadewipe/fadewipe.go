package fadewipe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const MAX_RGB = 255

type FadeInOutWipe struct {
	rgb       int
	SpeedRate float64
	isClosing bool
	width     int
	height    int
}

func (w *FadeInOutWipe) Init() error {

	return nil
}
func (w *FadeInOutWipe) Reset(width int, height int) error {
	w.rgb = 0
	w.isClosing = false
	w.width = width
	w.height = height
	return nil
}
func (w *FadeInOutWipe) Update() (bool, error) {

	speed := MAX_RGB * w.SpeedRate
	wipeEnd := false
	if !w.isClosing {
		w.rgb += int(speed)
		if w.rgb > MAX_RGB {
			w.isClosing = true
		}
	} else {
		w.rgb -= int(speed)
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
	if w.rgb > MAX_RGB {
		rgb = MAX_RGB
	}

	rect := ebiten.NewImage(w.width, w.height)
	rect.Fill(color.RGBA{0, 0, 0, rgb})
	screen.DrawImage(rect, nil)

}
