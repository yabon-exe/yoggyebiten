package chainfire

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
	"github.com/yabon-exe/yoggyebiten/game/util/physics"
)

const lifespan = 200
const decline = 50

type Fire struct {
	time        int
	color       color.RGBA
	declineTime int
	ignition    bool
	pos         model.Vertex
	vel0        *model.Velocity2d
	vel         *model.Velocity2d
	g           float64
	orbit       []model.Vertex
}

func NewFire(degree float64, speed float64, g float64, color color.RGBA) *Fire {

	velocity := model.NewVelocity2dFromDegree(degree)
	velocity.Scale(speed)
	velocity0 := model.NewVelocity2dFromDegree(degree)
	velocity0.Scale(speed)

	return &Fire{
		time:        0,
		color:       color,
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

		if fire.time > lifespan {
			fire.ignition = false
		} else if fire.time > decline {
			rate := 1.0 - (float64(fire.declineTime) / float64(lifespan-decline))
			if math.Abs(fire.vel.GetX()) > 0 {
				fire.vel.SetX(fire.vel.GetX() * rate)
			}
			if math.Abs(fire.vel.GetY()) > 0 {
				fire.vel.SetY(fire.vel.GetY() * rate)
			}
			fire.declineTime++
		} else {
			// Yは座標系が逆のためマイナス
			v0 := fire.vel0.GetY() // Y軸の値
			fire.vel.SetY(-physics.MoveFall(v0, fire.g, fire.time))

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

	color := color.RGBA{}
	if fire.time < decline {
		rate := 1.0 - float64(fire.time)/float64(decline)
		maxR := 255 - fire.color.R
		maxG := 255 - fire.color.G
		maxB := 255 - fire.color.B
		color.R = fire.color.R + uint8((float64(maxR) * rate))
		color.G = fire.color.G + uint8((float64(maxG) * rate))
		color.B = fire.color.B + uint8((float64(maxB) * rate))
	} else if fire.time < lifespan {
		rate := 1.0 - float64(fire.declineTime)/float64(lifespan-decline)
		color.R = uint8(float64(fire.color.R) * rate)
		color.G = uint8(float64(fire.color.G) * rate)
		color.B = uint8(float64(fire.color.B) * rate)
	}
	graphic.DrawLineArray(screen, fire.orbit, color, 1)
}

func (fire *Fire) Reset() {
	fire.time = 0
	fire.declineTime = 0
	fire.ignition = false
	fire.vel.SetX(fire.vel0.GetX())
	fire.vel.SetY(fire.vel0.GetY())
	fire.orbit = []model.Vertex{}
}
