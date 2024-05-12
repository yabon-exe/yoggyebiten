package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	Init() error
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	GetWindowOption() (windowTitle string, windowWidth, windowHeight int)
}
