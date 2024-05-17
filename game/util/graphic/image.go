package graphic

import (
	"image"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func ReadImageFile(imgFile fs.File) *ebiten.Image {
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func DrawImage(screen *ebiten.Image, img *ebiten.Image) {
	imageOptions := &ebiten.DrawImageOptions{}
	imageOptions.Filter = ebiten.FilterLinear
	screen.DrawImage(img, imageOptions)
}

func DrawBackImage(screen *ebiten.Image, img *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	scaleX := float64(screenWidth) / float64(imgWidth)
	scaleY := float64(screenHeight) / float64(imgHeight)

	op.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(img, op)
}
