package physics

import (
	"math"

	"github.com/yabon-exe/yoggyebiten/game/model"
)

func CheckCollisionVertexAndCircle(vertex model.Vertex, circle model.Circle) bool {
	distance := math.Sqrt((vertex.X-circle.Vertex.X)*(vertex.X-circle.Vertex.X) + (vertex.Y-circle.Vertex.Y)*(vertex.Y-circle.Vertex.Y))
	return distance <= float64(circle.Rad)
}
