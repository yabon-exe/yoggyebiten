package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/scene"
)

type MulitScene struct {
	gameOption GameOption
	nowScene   *scene.Scene
	sceneList  []*scene.Scene
}

func NewMulitSceneGame(gameOption GameOption, sceneList []*scene.Scene, nowSceneIdx int) Game {

	return &MulitScene{
		gameOption: gameOption,
		sceneList:  sceneList,
		nowScene:   sceneList[nowSceneIdx],
	}
}

func (game *MulitScene) Init() error {

	// for _, scene := range game.sceneList {
	// 	scene->Init()
	// }

	return nil
}

func (game *MulitScene) Update() error {
	return nil
}

func (game *MulitScene) Draw(screen *ebiten.Image) {

}

func (game *MulitScene) GetGameOption() (option GameOption) {
	return game.gameOption
}

func (game *MulitScene) ChangeScene(sceneIdx int) {

}
