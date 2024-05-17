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

func NewSampleGame() game.Game {
	return &SampleGame{
		count: 0,
	}
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
	// ebitenutil.DebugPrint(screen, "Hello, World!o0o0o0o0o0o0o0o0o0")
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(0, 0)
	// screen.DrawImage(runnerImage.SubImage(image.Rect(0, 0, 128, 128)).(*ebiten.Image), op)

	x, y := ebiten.CursorPosition()
	msg := fmt.Sprintf("Cursor Position: (%d, %d)", x, y)

	// 左マウスボタンが押されているか確認
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		msg += " - Left Button Pressed"
	}

	ebitenutil.DebugPrint(screen, msg)
}

func (g *SampleGame) GetGameOption() game.GameOption {
	option := game.GameOption{
		DeviceType:   game.PC,
		WindowTitle:  "*** Sample0001 Game ***",
		WindowWidth:  880,
		WindowHeight: 495,
	}
	return option
}
