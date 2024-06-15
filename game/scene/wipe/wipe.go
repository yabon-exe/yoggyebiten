package wipe

import "github.com/hajimehoshi/ebiten/v2"

type Wipe interface {
	Init() error
	Reset(width int, height int) error
	Update() (bool, error)
	Draw(screen *ebiten.Image, screenCapture *ebiten.Image)
}
