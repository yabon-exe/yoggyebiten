package chainfire

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
)

type FireWork struct {
	seedMode bool
	seedRad  int
	startPos model.Vertex
	fireList []*Fire
}

func NewFireWork(start model.Vertex, fireListNum int, power float64) *FireWork {

	list := []*Fire{}

	degree := model.PI_FULL_CIRCLE / float64(fireListNum)
	for i := 0; i < fireListNum; i++ {
		list = append(list, NewFire(degree*float64(i), power, 0.015))
	}

	return &FireWork{
		seedMode: true,
		seedRad:  int(power * 2),
		startPos: start,
		fireList: list,
	}
}

func (fireWork *FireWork) Update() {

	if !fireWork.seedMode {
		for _, fire := range fireWork.fireList {
			fire.Update()
		}
	}
}

func (fireWork *FireWork) Draw(screen *ebiten.Image) {

	if fireWork.seedMode {
		circle := model.Circle{
			X:   fireWork.startPos.X,
			Y:   fireWork.startPos.Y,
			Rad: fireWork.seedRad,
		}

		graphic.DrawCircle(screen, circle, color.RGBA{255, 255, 255, 100})
	} else {
		for _, fire := range fireWork.fireList {
			fire.Draw(screen)
		}
	}

}

func (fireWork *FireWork) Explode() {

	if fireWork.seedMode {
		for _, fire := range fireWork.fireList {
			fire.Ignit(fireWork.startPos)
		}
		fireWork.seedMode = false
	}
}

func (fireWork *FireWork) Move(x int, y int) {

	fireWork.startPos.X = float64(x)
	fireWork.startPos.Y = float64(y)

}
