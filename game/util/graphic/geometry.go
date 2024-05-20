package graphic

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

func DrawCircle(screen *ebiten.Image, circle model.Circle, color color.Color) {
	diameter := circle.Radius * 2
	img := image.NewRGBA(image.Rect(0, 0, diameter, diameter))
	for y := 0; y < diameter; y++ {
		for x := 0; x < diameter; x++ {
			dx := float64(x - circle.Radius)
			dy := float64(y - circle.Radius)
			if dx*dx+dy*dy <= float64(circle.Radius*circle.Radius) {
				img.Set(x, y, color)
			}
		}
	}
	circleImage := ebiten.NewImageFromImage(img)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(circle.X-float64(circle.Radius), circle.Y-float64(circle.Radius))
	screen.DrawImage(circleImage, op)
}

func DrawLineArray(screen *ebiten.Image, vertices []model.Vertex, color color.Color) {

	if len(vertices) < 2 {
		return
	}

	first := vertices[0]
	restVertices := vertices[1:]

	var path vector.Path
	path.MoveTo(float32(first.X), float32(first.Y))
	for _, v := range restVertices {
		path.LineTo(float32(v.X), float32(v.Y))
	}

	// 白い線で描画
	op := &vector.StrokeOptions{}
	op.Width = 2
	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, op)

	whiteImage := ebiten.NewImage(1, 1)
	whiteImage.Fill(color)

	screen.DrawTriangles(vs, is, whiteImage, &ebiten.DrawTrianglesOptions{})
}
