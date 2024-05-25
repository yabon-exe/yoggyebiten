package chainfire

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
)

//go:embed assets/*
var assets embed.FS

type ChainFire struct {
	backImg *ebiten.Image
	testFW  *FireWork
}

func NewGame() game.Game {
	return &ChainFire{}
}

func (chainFire *ChainFire) Init() error {

	imgBackFile, err := assets.Open("assets/back.png")
	if err != nil {
		return err
	}
	chainFire.backImg = graphic.ReadImageFile(imgBackFile)

	chainFire.testFW = NewFireWork(model.NewVertex(250, 150), 64, 2)

	return nil
}

func (chainFire *ChainFire) Update() error {
	chainFire.testFW.Update()

	x, y := ebiten.CursorPosition()
	chainFire.testFW.Move(x, y)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		chainFire.testFW.Explode()
	}

	return nil
}

func (chainFire *ChainFire) Draw(screen *ebiten.Image) {

	// ？？これがないと、画像読み込みで「image: unknown format」となる？？
	ebitenutil.DebugPrint(screen, "")

	graphic.DrawBackImage(screen, chainFire.backImg)

	chainFire.testFW.Draw(screen)
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
