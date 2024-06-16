package wipesample

import (
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe/picturewipe"
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
				&DummyScene{nextSceneIdx: DUMMY_SCENE_ID_2, startIdx: 0},
				&DummyScene{nextSceneIdx: DUMMY_SCENE_ID_1, startIdx: 1},
			},
			WipeList: []wipe.Wipe{
				// &fadewipe.FadeInOutWipe{SpeedRate: 0.05},
				// &curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_LEFT},
				// &curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_RIGHT},
				// &curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_UP},
				// &curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_DOWN},
				// &cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_LEFT},
				// &cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_RIGHT},
				// &cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_UP},
				// &cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_DOWN},
				// &picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_LEFT},
				// &picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_RIGHT},
				// &picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_UP},
				// &picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_DOWN},
				// &picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_UP_LEFT},
				// &picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_UP_RIGHT},
				// &picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_DOWN_LEFT},
				&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_DOWN_RIGHT},
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
