package chainfire

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
	"github.com/yabon-exe/yoggyebiten/game/util/physics"
)

type Fire struct {
	time     int
	loopRate float64
	pos      model.Vertex
	vel0     *model.Velocity2d
	vel      *model.Velocity2d
	g        float64
	orbit    []model.Vertex
}

func NewFire(start model.Vertex, degree float64, speed float64, g float64, loopRate float64) *Fire {

	velocity := model.NewVelocity2dFromDegree(degree)
	velocity.Scale(speed)
	velocity0 := model.NewVelocity2dFromDegree(degree)
	velocity0.Scale(speed)

	return &Fire{
		time:     0,
		loopRate: loopRate,
		pos:      start,
		vel:      velocity,
		vel0:     velocity0,
		g:        g,
		orbit:    []model.Vertex{start},
	}
}

func (fire *Fire) Update() {

	fire.time++
	reciprocal := 1 / fire.loopRate
	if fire.time%int(reciprocal) == 0 {
		// Yは座標系が逆のためマイナス
		v0 := fire.vel0.GetY() // Y軸の値
		fire.vel.SetY(-physics.MoveFall(v0, fire.g, fire.time))

		fire.pos.X += fire.vel.GetX()
		fire.pos.Y += fire.vel.GetY()

		// 負荷すごいなら考える
		fire.orbit = append(fire.orbit, model.NewVertex(fire.pos.X, fire.pos.Y))
	}

}

func (fire *Fire) Draw(screen *ebiten.Image) {

	graphic.DrawLineArray(screen, fire.orbit, color.RGBA{255, 0, 0, 100})
}
