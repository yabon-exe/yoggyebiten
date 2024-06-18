package sampletemplate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/system"
)

type DummyScene struct {
	scene.Scene
	keyboard *system.Keyboard
}

func (scene *DummyScene) Init() error {

	scene.keyboard = system.GetKeyboard()
	return nil
}

func (scene *DummyScene) Reset() error {

	return nil
}

func (scene *DummyScene) Update(isWiping bool) (int, int, error) {

	return -1, -1, nil
}

func (scene *DummyScene) Draw(screen *ebiten.Image) {

}
