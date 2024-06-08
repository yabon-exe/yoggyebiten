package chainfire

import (
	"embed"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/object/message"
	"github.com/yabon-exe/yoggyebiten/game/system"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
	"github.com/yabon-exe/yoggyebiten/game/util/physics"
)

//go:embed assets/*
var assetsMobile embed.FS

const GameWidthMobile = game.MOBILE_WIDTH
const GameHeightMobile = game.MOBILE_HEIGHT
const ShotIntervalMobile = 8
const GameEndTimeMobile = 640

type FWParamMobile struct {
	fireListNum int
	power       float64
	color       color.RGBA
}

type ChainFireMobile struct {
	time          int
	backImg       *ebiten.Image
	playerFw      *FireWork
	fwList        []*FireWork
	random        *system.Random
	limitedRandom *system.LimitedRandom[int]
	msg           *message.SimpleMessage
	msgR          *message.SimpleMessage
	explodeCount  int
}

func NewGameMobile() game.Game {
	return &ChainFireMobile{}
}

func (chainFire *ChainFireMobile) Init() error {

	chainFire.time = 0

	imgBackFile, err := assetsMobile.Open("assets/back.png")
	if err != nil {
		return err
	}
	chainFire.backImg = graphic.ReadImageFile(imgBackFile)

	chainFire.playerFw = NewFireWork(model.NewVertex(250, 150), 16, 1, color.RGBA{255, 255, 255, 0}, 0.06)

	params := []FWParamMobile{}
	colors := []color.RGBA{
		{R: 255, G: 0, B: 0, A: 0},
		{R: 0, G: 255, B: 0, A: 0},
		{R: 0, G: 0, B: 255, A: 0},
		{R: 255, G: 255, B: 0, A: 0},
		{R: 255, G: 0, B: 255, A: 0},
		{R: 0, G: 255, B: 255, A: 0},
	}

	// 花火の種類
	idxs := []int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k, color := range colors {
				params = append(params, FWParamMobile{fireListNum: (i + 1) * 12, power: float64(j + 1), color: color})
				idxs = append(idxs, i*3*len(colors)+j*len(colors)+k)
			}
		}
	}

	chainFire.limitedRandom = system.NewLimitedRandom[int](idxs)

	chainFire.random = system.NewRandom()
	fws := []*FireWork{}
	fwBownds := model.Bounds(GameWidthMobile/2.0, GameHeightMobile+200, GameWidthMobile, 200)
	for _, param := range params {
		x, y := chainFire.random.GetRandFromRect(fwBownds)
		fws = append(fws, NewFireWork(model.NewVertex(x, y), param.fireListNum, param.power, param.color, 0.026))
	}
	chainFire.fwList = fws

	// メッセージ
	chainFire.msg = &message.SimpleMessage{}
	chainFire.msg.Init()
	chainFire.msg.SetColor(0, 255, 0, 255)
	chainFire.msg.SetSize(40)
	chainFire.msg.SetPosition(model.Vertex{X: 300, Y: 150})
	chainFire.explodeCount = 0

	chainFire.msgR = &message.SimpleMessage{}
	chainFire.msgR.Init()
	chainFire.msgR.SetColor(255, 255, 255, 255)
	chainFire.msgR.SetSize(20)
	chainFire.msgR.SetPosition(model.Vertex{X: 300, Y: 220})
	chainFire.msgR.SetText("[R]キーでリトライ")

	return nil
}

func (chainFire *ChainFireMobile) Update() error {

	chainFire.explodeCount = 0

	// リセット
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		chainFire.reset()
	}

	// 自分花火
	chainFire.playerFw.Update()
	x, y := ebiten.CursorPosition()
	chainFire.playerFw.Move(x, y)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		chainFire.playerFw.Explode()
	}
	// スマホのタップ入力を取得
	touchIDs := []ebiten.TouchID{}
	touchIDs = ebiten.AppendTouchIDs(touchIDs[:0])
	for _, t := range touchIDs {
		x, y := ebiten.TouchPosition(t)
		chainFire.playerFw.Move(x, y)
		chainFire.playerFw.Explode()
	}

	// 打ち上げ
	if chainFire.time%ShotInterval == 0 && chainFire.time/ShotInterval < len(chainFire.fwList) {
		idx := chainFire.limitedRandom.PopRandValue()
		chainFire.fwList[idx].Shot()
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

	for _, fw := range chainFire.fwList {
		if fw.seedMode {
			for _, f := range fList {
				collisionFireMobile(fw, f)
			}
		} else {
			chainFire.explodeCount++
		}
	}

	chainFire.time++

	return nil
}

func (chainFire *ChainFireMobile) Draw(screen *ebiten.Image) {

	// ？？これがないと、画像読み込みで「image: unknown format」となる？？
	ebitenutil.DebugPrint(screen, "")

	graphic.DrawBackImage(screen, chainFire.backImg)

	chainFire.playerFw.Draw(screen)
	for _, fw := range chainFire.fwList {
		fw.Draw(screen)
	}

	// メッセージ
	if chainFire.time > GameEndTime {
		chainFire.msg.SetText(fmt.Sprintf("%d発(全%d発)", chainFire.explodeCount, len(chainFire.fwList)))
		chainFire.msg.Draw(screen)
		chainFire.msgR.Draw(screen)
	}
}

func (chainFire *ChainFireMobile) GetGameOption() game.GameOption {
	option := game.GameOption{
		DeviceType:   game.MOBILE_PHONE_PORTRAIT,
		WindowTitle:  "*** Yoggy ChainFire Mobile ***",
		WindowWidth:  GameWidthMobile / 2,
		WindowHeight: GameHeightMobile / 2,
	}
	return option
}

func (chainFire *ChainFireMobile) reset() {

	chainFire.explodeCount = 0
	chainFire.time = 0
	chainFire.limitedRandom.Reset()
	fwBownds := model.Bounds(GameWidthMobile/2.0, GameHeightMobile+200, GameWidthMobile, 200)
	for _, fw := range chainFire.fwList {
		x, y := chainFire.random.GetRandFromRect(fwBownds)
		fw.Move(int(x), int(y))
		fw.Reset()
	}
	chainFire.playerFw.Reset()
}

func collisionFireMobile(fw *FireWork, f *Fire) {
	if physics.CheckCollisionVertexAndCircle(f.pos, fw.seedBody) {
		fw.Explode()
	}
}
