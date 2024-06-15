package wipesample

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yabon-exe/yoggyebiten/game"
)

type DummyScene struct {
	game.Scene
}

func (scene *DummyScene) Init() error {
	return nil
}

func (scene *DummyScene) Reset(*ebiten.Image) error {
	return nil
}

func (scene *DummyScene) Update(bool) (int, int, error) {
	return -1, -1, nil
}

func (scene *DummyScene) Draw(screen *ebiten.Image) {

}
