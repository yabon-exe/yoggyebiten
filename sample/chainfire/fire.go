package chainfire

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
	"github.com/yabon-exe/yoggyebiten/game/util/physics"
)

const lifespan = 120
const decline = 80

type Fire struct {
	time        int
	declineTime int
	ignition    bool
	pos         model.Vertex
	vel0        *model.Velocity2d
	vel         *model.Velocity2d
	g           float64
	orbit       []model.Vertex
}

func NewFire(degree float64, speed float64, g float64) *Fire {

	velocity := model.NewVelocity2dFromDegree(degree)
	velocity.Scale(speed)
	velocity0 := model.NewVelocity2dFromDegree(degree)
	velocity0.Scale(speed)

	return &Fire{
		time:        0,
		declineTime: 0,
		ignition:    false,
		vel:         velocity,
		vel0:        velocity0,
		g:           g,
	}
}

func (fire *Fire) Update() {

	if fire.ignition {
		fire.time++

		// Yは座標系が逆のためマイナス
		v0 := fire.vel0.GetY() // Y軸の値
		fire.vel.SetY(-physics.MoveFall(v0, fire.g, fire.time))

		if fire.time > lifespan {
			fire.ignition = false
		} else if fire.time > decline {
			rate := 1.0 - (float64(fire.time) / float64(lifespan))
			if math.Abs(fire.vel.GetX()) > 0 {
				fire.vel.SetX(fire.vel.GetX() * rate)
			}
			if math.Abs(fire.vel.GetY()) > 0 {
				fire.vel.SetY(fire.vel.GetY() * rate)
			}
			fire.declineTime++
		}

		fire.pos.X += fire.vel.GetX()
		fire.pos.Y += fire.vel.GetY()

		// 負荷すごいなら考える
		fire.orbit = append(fire.orbit, model.NewVertex(fire.pos.X, fire.pos.Y))
	}
}

func (fire *Fire) Ignit(start model.Vertex) {
	fire.pos = start
	fire.orbit = []model.Vertex{start}
	fire.ignition = true
}

func (fire *Fire) Draw(screen *ebiten.Image) {

	graphic.DrawLineArray(screen, fire.orbit, color.RGBA{R: 200, G: 0, B: 0, A: 0}, 1)
}
