package actor

import "github.com/hajimehoshi/ebiten"

type Actor interface {
	Init() error
	Update() error
	CheckCollision(opponent Actor) bool
	OnCollision(opponent Actor)
	Draw(screen *ebiten.Image)
}
