package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/scene/wipe"
)

type MulitSceneGame struct {
	Option           GameOption
	SceneList        []scene.Scene
	WipeList         []wipe.Wipe
	nowSceneIdx      int
	nowWipeIdx       int
	isWiping         bool
	screenCaptureImg *ebiten.Image
}

func (game *MulitSceneGame) Init() error {

	for _, scene := range game.SceneList {
		scene.Init()
	}
	for _, wipe := range game.WipeList {
		wipe.Init()
	}

	game.isWiping = false
	game.screenCaptureImg = nil
	game.nowSceneIdx = 0
	game.SceneList[game.nowSceneIdx].Reset()
	game.nowWipeIdx = -1

	return nil
}

func (game *MulitSceneGame) Update() error {

	sceneIdx, wipeIdx, sErr := game.SceneList[game.nowSceneIdx].Update(game.isWiping)
	if sErr != nil {
		return sErr
	}

	if sceneIdx >= 0 {
		// 次のシーン開始
		game.nowSceneIdx = sceneIdx
		game.SceneList[game.nowSceneIdx].Reset()
	}

	if wipeIdx >= 0 {
		// ワイプ開始
		game.isWiping = true
		game.nowWipeIdx = wipeIdx
		layoutSize := game.Option.LayoutSize
		game.WipeList[game.nowWipeIdx].Reset(layoutSize.ToWH())
	}

	if game.isWiping {
		endWipe, wErr := game.WipeList[game.nowWipeIdx].Update()
		if wErr != nil {
			return sErr
		}
		if endWipe {
			game.isWiping = false
			game.nowWipeIdx = -1
			// ワイプ開始時ゲームキャプチャ情報を消す
			game.screenCaptureImg = nil
		}
	}

	return nil
}

func (game *MulitSceneGame) Draw(screen *ebiten.Image) {

	// ？？これがないと、画像読み込みで「image: unknown format」となる？？
	ebitenutil.DebugPrint(screen, "")

	game.SceneList[game.nowSceneIdx].Draw(screen)

	if !game.isWiping {
		// ワイプ起動時はキャプチャ情報なし
		// この時のゲーム画面を取得
		layoutSize := game.Option.LayoutSize
		game.screenCaptureImg = ebiten.NewImage(layoutSize.ToWH())
		game.screenCaptureImg.DrawImage(screen, nil)
	} else {
		game.WipeList[game.nowWipeIdx].Draw(screen, game.screenCaptureImg)
	}
}

func (g *MulitSceneGame) GetGameOption() GameOption {
	return g.Option
}
