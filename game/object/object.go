package object

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Init() error
	Update() error
	Draw(screen *ebiten.Image)
}
