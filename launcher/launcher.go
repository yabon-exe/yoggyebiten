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
	return option.WindowWidth, option.WindowHeight
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
