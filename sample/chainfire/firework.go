package chainfire

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
	"github.com/yabon-exe/yoggyebiten/game/util/physics"
)

const upV0 = 8.0

type FireWork struct {
	time        int
	explodeTime int
	enable      bool
	seedMode    bool
	seedBody    model.Circle
	seedVelY    float64
	fireList    []*Fire
	g           float64
}

func NewFireWork(start model.Vertex[float64], fireListNum int, power float64, color color.RGBA, g float64) *FireWork {

	list := []*Fire{}

	degree := model.PI_FULL_CIRCLE / float64(fireListNum)
	for i := 0; i < fireListNum; i++ {
		list = append(list, NewFire(degree*float64(i), power, 0.015, color))
	}

	return &FireWork{
		time:        0,
		explodeTime: 0,
		enable:      false,
		seedMode:    true,
		seedBody: model.Circle{
			Vertex: start,
			Rad:    3,
		},
		seedVelY: 0.0,
		fireList: list,
		g:        g,
	}
}

func (fireWork *FireWork) Update() {

	if fireWork.enable {
		fireWork.time++
		if fireWork.seedMode {
			fireWork.seedVelY = -physics.MoveFall(upV0, fireWork.g, fireWork.time)
			fireWork.seedBody.Vertex.Y += fireWork.seedVelY
		} else {
			fireWork.explodeTime++
			for _, fire := range fireWork.fireList {
				fire.Update()
			}
		}

		if fireWork.explodeTime > lifespan {
			fireWork.enable = false
		}
	}

}

func (fireWork *FireWork) Draw(screen *ebiten.Image) {

	if fireWork.enable {
		if fireWork.seedMode {
			graphic.DrawCircle(screen, fireWork.seedBody, color.RGBA{255, 255, 255, 0})
		} else {
			for _, fire := range fireWork.fireList {
				fire.Draw(screen)
			}
		}
	}
}

func (fireWork *FireWork) Shot() {
	fireWork.enable = true
}

func (fireWork *FireWork) Explode() {

	if fireWork.seedMode {
		for _, fire := range fireWork.fireList {
			fire.Ignit(fireWork.seedBody.Vertex)
		}
		fireWork.seedMode = false
	}
}

func (fireWork *FireWork) Move(x int, y int) {
	fireWork.enable = true
	fireWork.seedBody.Vertex.X = float64(x)
	fireWork.seedBody.Vertex.Y = float64(y)
}

func (fireWork *FireWork) Reset() {
	fireWork.time = 0
	fireWork.explodeTime = 0
	fireWork.enable = false
	fireWork.seedMode = true
	fireWork.seedVelY = 0.0
	for _, fire := range fireWork.fireList {
		fire.Reset()
	}
}
