package wipesample

import (
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe/curtainwipe"
)

type WipeSample struct {
	game.MulitSceneGame
}

const (
	DUMMY_SCENE_ID_1 = 0
	DUMMY_SCENE_ID_2 = 1

	WIPE_FADEINOUT_IDX = 0
)

func NewGame() *WipeSample {
	return &WipeSample{
		MulitSceneGame: game.MulitSceneGame{
			SceneList: []scene.Scene{
				&DummyScene{nextSceneIdx: DUMMY_SCENE_ID_2},
				&DummyScene{nextSceneIdx: DUMMY_SCENE_ID_1},
			},
			WipeList: []wipe.Wipe{
				// &fadewipe.FadeInOutWipe{Speed: 5},
				// &curtainwipe.CurtainWipe{Speed: 50, Direct: curtainwipe.MOTION_LEFT},
				// &curtainwipe.CurtainWipe{Speed: 50, Direct: curtainwipe.MOTION_RIGHT},
				// &curtainwipe.CurtainWipe{Speed: 20, Direct: curtainwipe.MOTION_UP},
				&curtainwipe.CurtainWipe{Speed: 20, Direct: curtainwipe.MOTION_DOWN},
			},
		},
	}
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
