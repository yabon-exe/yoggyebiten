package launcher

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game"
)

func RunGame(game game.Game) {
	wTitle, wWidth, wHeight := game.GetWindowOption()
	ebiten.SetWindowSize(wWidth, wHeight)
	ebiten.SetWindowTitle(wTitle)

	game.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
