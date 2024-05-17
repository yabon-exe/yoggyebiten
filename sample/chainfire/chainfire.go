package chainfire

import (
	"embed"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/util"
)

//go:embed assets/*
var assets embed.FS

type ChainFire struct {
	backImg *ebiten.Image
}

func NewGame() game.Game {
	return &ChainFire{}
}

func (chainFire *ChainFire) Init() error {

	imgBackFile, err := assets.Open("assets/back.png")
	if err != nil {
		return err
	}
	chainFire.backImg = util.ReadImageFile(imgBackFile)
	return nil
}

func (chainFire *ChainFire) Update() error {
	return nil
}

func (chainFire *ChainFire) Draw(screen *ebiten.Image) {

	util.DrawBackImage(screen, chainFire.backImg)

	circle := util.Circle{
		X:      0,
		Y:      0,
		Radius: 495,
		Color:  color.RGBA{0, 255, 255, 100},
	}

	util.DrawCircle(screen, circle)
	x, y := ebiten.CursorPosition()
	msg := fmt.Sprintf("Cursor Position: (%d, %d)", x, y)
	ebitenutil.DebugPrint(screen, msg)

	v1 := util.Vertex{
		X: 30,
		Y: 30,
	}
	v2 := util.Vertex{
		X: 50,
		Y: 90,
	}
	v3 := util.Vertex{
		X: 100,
		Y: 100,
	}
	util.DrawLineArray(screen, []util.Vertex{v1, v2, v3}, color.RGBA{255, 0, 255, 0})
}

func (chainFire *ChainFire) GetGameOption() game.GameOption {
	option := game.GameOption{
		DeviceType:   game.PC,
		WindowTitle:  "*** Yoggy ChainFire ***",
		WindowWidth:  880,
		WindowHeight: 495,
	}
	return option
}
