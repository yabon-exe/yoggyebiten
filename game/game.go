package game

import (
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/system"
)

type GameDeviceType int

const (
	PC GameDeviceType = iota
	MOBILE_PHONE_PORTRAIT
	MOBILE_PHONE_LANDSCAPE
	MOBILE_TABLET_PORTRAIT
	MOBILE_TABLET_LANDSCAPE
)

const (
	// 開発環境の大きさ次第
	MOBILE_WIDTH  = 720
	MOBILE_HEIGHT = 1280
)

type GameOption struct {
	DeviceType  GameDeviceType
	WindowTitle string
	WindowSize  model.Size[int]
	LayoutSize  model.Size[int]
	LogLevel    system.LogLevel
}

type Game interface {
	Init() error
	Update() error
	Draw(screen *ebiten.Image)
	GetGameOption() (option GameOption)
}

func GetDefaulDeviceSize(deviceType GameDeviceType) model.Size[int] {

	switch deviceType {
	case MOBILE_PHONE_PORTRAIT:
		return model.Size[int]{W: MOBILE_WIDTH, H: MOBILE_HEIGHT}
	case MOBILE_PHONE_LANDSCAPE:
		return model.Size[int]{W: MOBILE_HEIGHT, H: MOBILE_WIDTH}
	case MOBILE_TABLET_PORTRAIT:
		return model.Size[int]{W: MOBILE_WIDTH, H: MOBILE_HEIGHT}
	case MOBILE_TABLET_LANDSCAPE:
		return model.Size[int]{W: MOBILE_HEIGHT, H: MOBILE_WIDTH}
	default:
		return model.Size[int]{W: 880, H: 495}
	}

}

func IsMobile() bool {
	return ebiten.Monitor().DeviceScaleFactor() > 1.0 || runtime.GOOS == "android" || runtime.GOOS == "ios"
}
