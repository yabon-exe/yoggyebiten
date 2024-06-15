package wipesample

import "github.com/yabon-exe/yoggyebiten/game"

type WipeSample struct {
	game.MulitSceneGame
}

func NewGame() game.Game {
	return &WipeSample{}
}

func (game *WipeSample) CrateSceneList() []*game.Scene {
	return nil
}

func (game *WipeSample) CrateWipeList() []*game.Wipe {
	return nil
}

func (g *WipeSample) GetGameOption() game.GameOption {
	option := game.GameOption{
		DeviceType:   game.PC,
		WindowTitle:  "*** Wipe Sample ***",
		WindowWidth:  880.0,
		WindowHeight: 495.0,
	}
	return option
}
