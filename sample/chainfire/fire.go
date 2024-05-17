package chainfire

import "github.com/hajimehoshi/ebiten/v2"

type Point struct {
	X int
	Y int
}

type Fire struct {
	time   int
	StartX int
	StartY int
	orbit  []Point
}

func NewFire() Fire {
	return Fire{
		time: 0,
	}
}

func (fire *Fire) Update() error {

	fire.time++
	return nil
}

func (fire *Fire) Draw(screen *ebiten.Image) {

}
