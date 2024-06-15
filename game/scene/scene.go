package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Init() error
	Reset() error
	Update(bool) (int, int, error)
	Draw(screen *ebiten.Image)
}
