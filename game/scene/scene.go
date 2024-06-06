package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Init() error
	Reset() error
	Update() error
	Draw(screen *ebiten.Image)
}
