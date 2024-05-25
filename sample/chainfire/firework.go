package chainfire

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
	"github.com/yabon-exe/yoggyebiten/game/util/physics"
)

const upV0 = 8.0
const g = 0.1

type FireWork struct {
	time     int
	seedMode bool
	seedRad  int
	seedVelY float64
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
		time:     0,
		seedMode: true,
		seedRad:  int(power * 2),
		seedVelY: 0.0,
		startPos: start,
		fireList: list,
	}
}

func (fireWork *FireWork) Update() {
	fireWork.time++
	if fireWork.seedMode {
		fireWork.seedVelY = -physics.MoveFall(upV0, g, fireWork.time)
		fireWork.startPos.Y += fireWork.seedVelY
	} else {
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
