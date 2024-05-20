package chainfire

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type FireWork struct {
	fireList []*Fire
}

func NewFireWork(start model.Vertex, fireListNum int, speed float64) *FireWork {

	list := []*Fire{}

	degree := model.PI_FULL_CIRCLE / float64(fireListNum)
	for i := 0; i < fireListNum; i++ {
		list = append(list, NewFire(start, degree*float64(i), speed, 0.002, 1.0))
	}

	return &FireWork{
		fireList: list,
	}
}

func (fireWork *FireWork) Update() {

	for _, fire := range fireWork.fireList {
		fire.Update()
	}
}

func (fireWork *FireWork) Draw(screen *ebiten.Image) {
	for _, fire := range fireWork.fireList {
		fire.Draw(screen)
	}
}
