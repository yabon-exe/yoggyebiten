package wipesample

import (
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe/cleanerwipe"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe/curtainwipe"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe/fadewipe"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe/picturewipe"
)

type WipeSample struct {
	game.MulitSceneGame
}

const (
	DUMMY_SCENE_ID_1 = 0
	DUMMY_SCENE_ID_2 = 1

	WIPE_FADEINOUT_IDX  = 0
	WIPE_CURTAIN_L_IDX  = 1
	WIPE_CURTAIN_R_IDX  = 2
	WIPE_CURTAIN_U_IDX  = 3
	WIPE_CURTAIN_D_IDX  = 4
	WIPE_CLEANER_L_IDX  = 5
	WIPE_CLEANER_R_IDX  = 6
	WIPE_CLEANER_U_IDX  = 7
	WIPE_CLEANER_D_IDX  = 8
	WIPE_PICTURE_L_IDX  = 9
	WIPE_PICTURE_R_IDX  = 10
	WIPE_PICTURE_U_IDX  = 11
	WIPE_PICTURE_D_IDX  = 12
	WIPE_PICTURE_UL_IDX = 13
	WIPE_PICTURE_UR_IDX = 14
	WIPE_PICTURE_DL_IDX = 15
	WIPE_PICTURE_DR_IDX = 16
)

func NewGame() *WipeSample {

	wList := []wipe.Wipe{
		&fadewipe.FadeInOutWipe{SpeedRate: 0.05},
		&curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_LEFT},
		&curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_RIGHT},
		&curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_UP},
		&curtainwipe.CurtainWipe{SpeedRate: 0.05, Direct: curtainwipe.MOTION_DOWN},
		&cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_LEFT},
		&cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_RIGHT},
		&cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_UP},
		&cleanerwipe.CleanerWipe{SpeedRate: 0.01, Direct: cleanerwipe.MOTION_DOWN},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_LEFT},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_RIGHT},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_UP},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_DOWN},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_UP_LEFT},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_UP_RIGHT},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_DOWN_LEFT},
		&picturewipe.PictureWipe{SpeedRate: 0.01, Direct: picturewipe.MOTION_DOWN_RIGHT},
	}

	// option := game.GameOption{
	// 	DeviceType:   game.PC,
	// 	WindowTitle:  "*** Wipe Sample ***",
	// 	WindowWidth:  880.0,
	// 	WindowHeight: 495.0,
	// }
	option := game.GameOption{
		DeviceType:   game.MOBILE_PHONE_PORTRAIT,
		WindowTitle:  "*** Wipe Sample ***",
		WindowWidth:  game.MOBILE_WIDTH / 2,
		WindowHeight: game.MOBILE_HEIGHT / 2,
		LayoutWidth:  game.MOBILE_WIDTH,
		LayoutHeight: game.MOBILE_HEIGHT,
	}

	return &WipeSample{
		MulitSceneGame: game.MulitSceneGame{
			Option: option,
			SceneList: []scene.Scene{
				&DummyScene{nextSceneIdx: DUMMY_SCENE_ID_2, startIdx: 0},
				&DummyScene{nextSceneIdx: DUMMY_SCENE_ID_1, startIdx: 1},
			},
			WipeList: wList,
		},
	}
}

// func (g *WipeSample) GetGameOption() game.GameOption {
// option := game.GameOption{
// 	DeviceType:   game.PC,
// 	WindowTitle:  "*** Wipe Sample ***",
// 	WindowWidth:  880.0,
// 	WindowHeight: 495.0,
// }
// 	option := game.GameOption{
// 		DeviceType:   game.MOBILE_PHONE_PORTRAIT,
// 		WindowTitle:  "*** Wipe Sample ***",
// 		WindowWidth:  game.MOBILE_WIDTH / 2,
// 		WindowHeight: game.MOBILE_HEIGHT / 2,
// 		LayoutWidth:  game.MOBILE_WIDTH,
// 		LayoutHeight: game.MOBILE_HEIGHT,
// 	}
// 	return option
// }
