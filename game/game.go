package game

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	DeviceType   GameDeviceType
	WindowTitle  string
	WindowWidth  int
	WindowHeight int
	LogLevel     system.LogLevel
}

type Game interface {
	Init() error
	Update() error
	Draw(screen *ebiten.Image)
	GetGameOption() (option GameOption)
}
