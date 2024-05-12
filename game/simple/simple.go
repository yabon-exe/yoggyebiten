package simple

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game"
)

type SimpleGame struct {
	v string
}

func (g *SimpleGame) Init() error {
	return nil
}

func (g *SimpleGame) Update() error {
	return nil
}

func (g *SimpleGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!o0o0o0o0o0o0o0o0o0")
}

func (g *SimpleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func (g *SimpleGame) GetWindowOption() (windowTitle string, windowWidth, windowHeight int) {
	return "*** Simple Game ***", 640, 480
}

func New() game.Game {
	return &SimpleGame{
		v: "taro",
	}
}
