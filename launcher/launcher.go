package launcher

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"
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

	// logger := system.GetLogger(system.INFO)
	option := ebitenGame.yoggyGame.GetGameOption()
	return option.LayoutSize.ToWH()
}

func RunGame(yoggyGame game.Game) {

	logger := system.GetLogger(system.INFO)
	option := yoggyGame.GetGameOption()

	ebiten.SetWindowTitle(option.WindowTitle)

	mobileMode := false
	switch option.DeviceType {
	case game.PC:
		logger.Info("device type: PC")
	case game.MOBILE_PHONE_PORTRAIT:
		mobileMode = true
		logger.Info("device type: MOBILE_PHONE_PORTRAIT")
	case game.MOBILE_PHONE_LANDSCAPE:
		mobileMode = true
		logger.Info("device type: MOBILE_PHONE_LANDSCAPE")
	case game.MOBILE_TABLET_PORTRAIT:
		mobileMode = true
		logger.Info("device type: MOBILE_TABLET_PORTRAIT")
	case game.MOBILE_TABLET_LANDSCAPE:
		mobileMode = true
		logger.Info("device type: MOBILE_TABLET_LANDSCAPE")
	default:
		logger.Fatal("unexpected DeviceType.")
	}
	ebiten.SetWindowSize(option.WindowSize.ToWH())

	game := &EbitenGame{
		yoggyGame: yoggyGame,
	}

	yoggyGame.Init()

	if mobileMode {
		mobile.SetGame(game)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
