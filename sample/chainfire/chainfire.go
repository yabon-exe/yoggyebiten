package chainfire

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/system"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
)

//go:embed assets/*
var assets embed.FS

const GameWidth = 880.0
const GameHeight = 495.0
const FireWorkNum = 16
const ShotInterval = 10

type ChainFire struct {
	time    int
	backImg *ebiten.Image
	testFW  *FireWork
	fwList  []*FireWork
	random  *system.Random
}

func NewGame() game.Game {
	return &ChainFire{
		time:   0,
		random: system.NewRandom(),
	}
}

func (chainFire *ChainFire) Init() error {

	imgBackFile, err := assets.Open("assets/back.png")
	if err != nil {
		return err
	}
	chainFire.backImg = graphic.ReadImageFile(imgBackFile)

	chainFire.testFW = NewFireWork(model.NewVertex(250, 150), 64, 2)

	fws := []*FireWork{}
	fwBownds := model.Bounds(GameWidth/2.0-20, GameHeight+200.0, GameWidth*0.8, 200)
	for i := 0; i < FireWorkNum; i++ {
		x, y := chainFire.random.GetRandFromRect(fwBownds)
		fws = append(fws, NewFireWork(model.NewVertex(x, y), 32, 2))
	}
	chainFire.fwList = fws

	return nil
}

func (chainFire *ChainFire) Update() error {

	if chainFire.time%ShotInterval == 0 && chainFire.time/ShotInterval < FireWorkNum {
		chainFire.fwList[chainFire.time/ShotInterval].Shot()
	}

	chainFire.testFW.Update()

	x, y := ebiten.CursorPosition()
	chainFire.testFW.Move(x, y)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		chainFire.testFW.Explode()
	}

	for i := 0; i < FireWorkNum; i++ {
		chainFire.fwList[i].Update()
	}

	chainFire.time++

	return nil
}

func (chainFire *ChainFire) Draw(screen *ebiten.Image) {

	// ？？これがないと、画像読み込みで「image: unknown format」となる？？
	ebitenutil.DebugPrint(screen, "")

	graphic.DrawBackImage(screen, chainFire.backImg)

	chainFire.testFW.Draw(screen)
	for i := 0; i < FireWorkNum; i++ {
		chainFire.fwList[i].Draw(screen)
	}
}

func (chainFire *ChainFire) GetGameOption() game.GameOption {
	option := game.GameOption{
		DeviceType:   game.PC,
		WindowTitle:  "*** Yoggy ChainFire ***",
		WindowWidth:  GameWidth,
		WindowHeight: GameHeight,
	}
	return option
}
