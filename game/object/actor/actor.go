package actor

import "github.com/yabon-exe/yoggyebiten/game/object"

type Actor interface {
	object.Object
	CheckCollision(opponent Actor) bool
	OnCollision(opponents []Actor)
}
