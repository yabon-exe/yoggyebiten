package launcher

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game"
	"github.com/yabon-exe/yoggyebiten/game/system"
)

type EbitenGame struct {
	yoggyGame game.Game
}

func (game *EbitenGame) Update() error {
	return game.yoggyGame.Update()
}

func (game *EbitenGame) Draw(screen *ebiten.Image) {
	game.yoggyGame.Draw(screen)
}

func (ebitenGame *EbitenGame) Layout(outsideWidth int, outsideHeight int) (int, int) {

	logger := system.GetLogger(system.INFO)
	option := ebitenGame.yoggyGame.GetGameOption()

	switch option.DeviceType {
	case game.PC:
		return option.WindowWidth, option.WindowHeight
	case game.MOBILE_PHONE_PORTRAIT:
		return game.MOBILE_WIDTH, game.MOBILE_HEIGHT
	case game.MOBILE_PHONE_LANDSCAPE:
		return game.MOBILE_HEIGHT, game.MOBILE_WIDTH
	case game.MOBILE_TABLET_PORTRAIT:
		return game.MOBILE_WIDTH, game.MOBILE_HEIGHT
	case game.MOBILE_TABLET_LANDSCAPE:
		return game.MOBILE_HEIGHT, game.MOBILE_WIDTH
	default:
		logger.Fatal("unexpected DeviceType.")
	}

	// // 非常に縦長のデバイスの場合（スマートフォンなど）、9:16のアスペクト比を使用
	// if float64(outsideHeight) > float64(outsideWidth)*1.8 {
	// 	aspectRatio := 9.0 / 16.0
	// 	logger.Info("window aspect ratio 9:16 (SmartPhone)")
	// 	return int(float64(outsideHeight) * aspectRatio), outsideHeight
	// }
	// // 縦長のデバイス（タブレット等）の場合は3:4のアスペクト比を使用
	// if float64(outsideHeight) > float64(outsideWidth)*1.3 {
	// 	aspectRatio := 3.0 / 4.0
	// 	logger.Info("window aspect ratio 3:4 (Tablet)")
	// 	return int(float64(outsideHeight) * aspectRatio), outsideHeight
	// }
	// // 横長のデバイス（デスクトップ等）の場合は16:9のアスペクト比を使用
	// aspectRatio := 16.0 / 9.0
	// screenWidth := int(float64(outsideHeight) * aspectRatio)
	// if screenWidth > outsideWidth {
	// 	screenWidth = outsideWidth
	// 	screenHeight := int(float64(screenWidth) / aspectRatio)
	// 	logger.Info("window aspect ratio 16:9 (PC Adjusted)")
	// 	return screenWidth, screenHeight
	// }

	// logger.Info("window aspect ratio 16:9 (PC)")
	return option.WindowWidth, option.WindowHeight
	// return option.WindowWidth * 2, option.WindowHeight * 2
}

func RunGame(yoggyGame game.Game) {

	logger := system.GetLogger(system.INFO)
	option := yoggyGame.GetGameOption()

	ebiten.SetWindowTitle(option.WindowTitle)

	// スマホ、タブレットは大きさ固定
	switch option.DeviceType {
	case game.PC:
		logger.Info("device type: PC")
	case game.MOBILE_PHONE_PORTRAIT:
		logger.Info("device type: MOBILE_PHONE_PORTRAIT")
	case game.MOBILE_PHONE_LANDSCAPE:
		logger.Info("device type: MOBILE_PHONE_LANDSCAPE")
	case game.MOBILE_TABLET_PORTRAIT:
		logger.Info("device type: MOBILE_TABLET_PORTRAIT")
	case game.MOBILE_TABLET_LANDSCAPE:
		logger.Info("device type: MOBILE_TABLET_LANDSCAPE")
	default:
		logger.Fatal("unexpected DeviceType.")
	}
	ebiten.SetWindowSize(option.WindowWidth, option.WindowHeight)

	game := &EbitenGame{
		yoggyGame: yoggyGame,
	}

	yoggyGame.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
