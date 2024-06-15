package uisample

import "github.com/yabon-exe/yoggyebiten/game"

type UiSample struct {
	game.MulitSceneGame
}

func NewGame() game.Game {
	return &UiSample{}
}

func (game *UiSample) CrateSceneList() []*game.Scene {
	return nil
}

func (game *UiSample) CrateWipeList() []*game.Wipe {
	return nil
}

func (g *UiSample) GetGameOption() game.GameOption {
	option := game.GameOption{
		DeviceType:   game.PC,
		WindowTitle:  "*** Ui Sample ***",
		WindowWidth:  880.0,
		WindowHeight: 495.0,
	}
	return option
}
