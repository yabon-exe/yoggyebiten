package chainfire

import (
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/system"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
	"github.com/yabon-exe/yoggyebiten/game/util/physics"
)

//go:embed assets/*
var assets embed.FS

const GameWidth = 880.0
const GameHeight = 495.0
const ShotInterval = 8

type FWParam struct {
	fireListNum int
	power       float64
	color       color.RGBA
}

type ChainFire struct {
	time     int
	backImg  *ebiten.Image
	playerFw *FireWork
	fwList   []*FireWork
	random   *system.Random
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

	chainFire.playerFw = NewFireWork(model.NewVertex(250, 150), 64, 2, color.RGBA{255, 255, 255, 0})

	params := []FWParam{}
	colors := []color.RGBA{
		{R: 255, G: 0, B: 0, A: 0},
		{R: 0, G: 255, B: 0, A: 0},
		{R: 0, G: 0, B: 255, A: 0},
		{R: 255, G: 255, B: 0, A: 0},
		{R: 255, G: 0, B: 255, A: 0},
		{R: 0, G: 255, B: 255, A: 0},
	}

	// 花火の種類
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for _, color := range colors {
				params = append(params, FWParam{fireListNum: i * 16, power: float64(j), color: color})
			}
		}
	}

	fws := []*FireWork{}
	fwBownds := model.Bounds(GameWidth/2.0-20, GameHeight+200.0, GameWidth*0.8, 200)
	for _, param := range params {
		x, y := chainFire.random.GetRandFromRect(fwBownds)
		fws = append(fws, NewFireWork(model.NewVertex(x, y), param.fireListNum, param.power, param.color))
	}
	chainFire.fwList = fws

	return nil
}

func (chainFire *ChainFire) Update() error {

	if chainFire.time%ShotInterval == 0 && chainFire.time/ShotInterval < len(chainFire.fwList) {
		chainFire.fwList[chainFire.time/ShotInterval].Shot()
	}

	chainFire.playerFw.Update()

	x, y := ebiten.CursorPosition()
	chainFire.playerFw.Move(x, y)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		chainFire.playerFw.Explode()
	}

	// 衝突判定
	fList := []*Fire{}
	if !chainFire.playerFw.seedMode {
		fList = append(fList, chainFire.playerFw.fireList...)
	}
	for _, fw := range chainFire.fwList {
		fw.Update()
		if !fw.seedMode {
			fList = append(fList, fw.fireList...)
		}
	}
	// for i := 0; i < FireWorkNum; i++ {
	// 	fw := chainFire.fwList[i]
	// 	fw.Update()

	// 	if !fw.seedMode {
	// 		fList = append(fList, fw.fireList...)
	// 	}
	// }
	for _, fw := range chainFire.fwList {
		if fw.seedMode {
			for _, f := range fList {
				CollisionFire(fw, f)
			}
		}
	}
	// for i := 0; i < FireWorkNum; i++ {
	// 	fw := chainFire.fwList[i]
	// 	if fw.seedMode {
	// 		for _, f := range fList {
	// 			CollisionFire(fw, f)
	// 		}
	// 	}
	// }

	chainFire.time++

	return nil
}

func (chainFire *ChainFire) Draw(screen *ebiten.Image) {

	// ？？これがないと、画像読み込みで「image: unknown format」となる？？
	ebitenutil.DebugPrint(screen, "")

	graphic.DrawBackImage(screen, chainFire.backImg)

	chainFire.playerFw.Draw(screen)
	for _, fw := range chainFire.fwList {
		fw.Draw(screen)
	}
	// for i := 0; i < FireWorkNum; i++ {
	// 	chainFire.fwList[i].Draw(screen)
	// }
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

func CollisionFire(fw *FireWork, f *Fire) {
	if physics.CheckCollisionVertexAndCircle(f.pos, fw.seedBody) {
		fw.Explode()
	}
}
