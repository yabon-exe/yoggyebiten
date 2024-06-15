package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init() error
	Reset(*ebiten.Image) error
	Update(bool) (int, int, error)
	Draw(screen *ebiten.Image)
}

type Wipe interface {
	Init() error
	Reset() error
	Update() (bool, error)
	Draw(screen *ebiten.Image, screenCapture *ebiten.Image)
}

type IMulitSceneGame interface {
	Game
	crateSceneList() []*Scene
	crateWipeList() []*Wipe
}

type MulitSceneGame struct {
	IMulitSceneGame
	sceneList        []*Scene
	wipeList         []*Wipe
	nowSceneIdx      int
	nowWipeIdx       int
	isWiping         bool
	screenCaptureImg *ebiten.Image
}

func (game *MulitSceneGame) Init() error {

	game.sceneList = game.crateSceneList()
	game.wipeList = game.crateWipeList()

	game.nowSceneIdx = 0
	game.nowWipeIdx = -1

	for _, scene := range game.sceneList {
		(*scene).Init()
	}
	for _, wipe := range game.wipeList {
		(*wipe).Init()
	}

	game.isWiping = false

	game.screenCaptureImg = nil

	return nil
}

func (game *MulitSceneGame) Update() error {

	sceneIdx, wipeIdx, sErr := (*game.sceneList[game.nowSceneIdx]).Update(game.isWiping)
	if sErr != nil {
		return sErr
	}

	if sceneIdx >= 0 {
		game.nowSceneIdx = sceneIdx
	}

	if wipeIdx >= 0 {

		game.isWiping = true
		game.nowWipeIdx = wipeIdx
	}

	if game.isWiping {
		endWipe, wErr := (*game.wipeList[game.nowWipeIdx]).Update()
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
	if game.isWiping {
		if game.screenCaptureImg == nil {
			// ワイプ起動時はキャプチャ情報なし
			// この時のゲーム画面を取得
			game.screenCaptureImg = ebiten.NewImage(ebiten.WindowSize())
			game.screenCaptureImg.DrawImage(screen, nil)
		}
		(*game.wipeList[game.nowWipeIdx]).Draw(screen, game.screenCaptureImg)
	}

	(*game.sceneList[game.nowSceneIdx]).Draw(screen)
}
