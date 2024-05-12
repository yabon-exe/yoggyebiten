package sample

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game"
)

var (
	runnerImage *ebiten.Image
)

type SampleGame struct {
	count int
}

func (g *SampleGame) Init() error {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)
	return nil
}

func (g *SampleGame) Update() error {
	return nil
}

func (g *SampleGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!o0o0o0o0o0o0o0o0o0")
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(runnerImage.SubImage(image.Rect(0, 0, 128, 128)).(*ebiten.Image), op)
}

func (g *SampleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	fmt.Println(outsideWidth)
	fmt.Println(outsideHeight)
	return 320, 240
}

func (g *SampleGame) GetWindowOption() (windowTitle string, windowWidth, windowHeight int) {
	return "*** Sample01 Game ***", 640 * 2, 480 * 2
}

func New() game.Game {
	return &SampleGame{
		count: 0,
	}
}
