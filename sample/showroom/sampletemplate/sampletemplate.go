package sampletemplate

import (
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe"
)

type UiSample struct {
	game.MulitSceneGame
}

func NewGame() game.Game {

	option := game.GameOption{
		DeviceType:  game.PC,
		WindowTitle: "*** Sample ***",
		WindowSize:  game.GetDefaulDeviceSize(game.PC),
		LayoutSize:  game.GetDefaulDeviceSize(game.PC),
	}

	return &UiSample{
		MulitSceneGame: game.MulitSceneGame{
			Option: option,
			SceneList: []scene.Scene{
				&DummyScene{},
			},
			WipeList: []wipe.Wipe{},
		},
	}
}

func (game *UiSample) CrateSceneList() []*scene.Scene {
	return nil
}

func (game *UiSample) CrateWipeList() []*wipe.Wipe {
	return nil
}
