package sample

import (
	"bytes"
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

func (g *SampleGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	// 非常に縦長のデバイスの場合（スマートフォンなど）、9:16のアスペクト比を使用
	if float64(outsideHeight) > float64(outsideWidth)*1.8 {
		aspectRatio := 9.0 / 16.0
		return int(float64(outsideHeight) * aspectRatio), outsideHeight
	}
	// 縦長のデバイス（タブレット等）の場合は3:4のアスペクト比を使用
	if float64(outsideHeight) > float64(outsideWidth)*1.3 {
		aspectRatio := 3.0 / 4.0
		return int(float64(outsideHeight) * aspectRatio), outsideHeight
	}
	// 横長のデバイス（デスクトップ等）の場合は16:9のアスペクト比を使用
	aspectRatio := 16.0 / 9.0
	screenWidth := int(float64(outsideHeight) * aspectRatio)
	if screenWidth > outsideWidth {
		screenWidth = outsideWidth
		screenHeight := int(float64(screenWidth) / aspectRatio)
		return screenWidth, screenHeight
	}
	return screenWidth, outsideHeight
}

func (g *SampleGame) GetWindowOption() (windowTitle string, windowWidth, windowHeight int) {
	return "*** Sample01 Game ***", 640 * 2, 480 * 2
}

func New() game.Game {
	return &SampleGame{
		count: 0,
	}
}
