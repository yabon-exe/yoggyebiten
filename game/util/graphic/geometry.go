package graphic

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

func DrawCircle(screen *ebiten.Image, circle model.Circle, color color.Color) {
	diameter := circle.Rad * 2
	img := image.NewRGBA(image.Rect(0, 0, diameter, diameter))
	for y := 0; y < diameter; y++ {
		for x := 0; x < diameter; x++ {
			dx := float64(x - circle.Rad)
			dy := float64(y - circle.Rad)
			if dx*dx+dy*dy <= float64(circle.Rad*circle.Rad) {
				img.Set(x, y, color)
			}
		}
	}
	circleImage := ebiten.NewImageFromImage(img)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(circle.X-float64(circle.Rad), circle.Y-float64(circle.Rad))
	screen.DrawImage(circleImage, op)
}

func DrawLineArray(screen *ebiten.Image, vertices []model.Vertex, color color.RGBA, width float32) {

	if len(vertices) < 2 {
		return
	}

	for i := 0; i < len(vertices)-1; i++ {
		vector.StrokeLine(
			screen,
			float32(vertices[i].X),
			float32(vertices[i].Y),
			float32(vertices[i+1].X),
			float32(vertices[i+1].Y),
			1,
			color,
			true,
		)
	}
}
