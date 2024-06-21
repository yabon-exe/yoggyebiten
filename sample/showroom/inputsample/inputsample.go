package inputsample

import (
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe"
)

type InputSample struct {
	game.MulitSceneGame
}

func NewGame() game.Game {

	option := game.GameOption{
		DeviceType:  game.PC,
		WindowTitle: "*** UI Sample ***",
		WindowSize:  game.GetDefaulDeviceSize(game.PC),
		LayoutSize:  game.GetDefaulDeviceSize(game.PC),
	}

	return &InputSample{
		MulitSceneGame: game.MulitSceneGame{
			Option: option,
			SceneList: []scene.Scene{
				&DummyScene{},
			},
			WipeList: []wipe.Wipe{},
		},
	}
}

func (game *InputSample) CrateSceneList() []*scene.Scene {
	return nil
}

func (game *InputSample) CrateWipeList() []*wipe.Wipe {
	return nil
}
