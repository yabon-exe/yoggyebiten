package wipesample

import (
	"embed"
	"fmt"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/system"
	"github.com/yabon-exe/yoggyebiten/game/util/graphic"
)

//go:embed assets/*
var assets embed.FS

const backImgNum = 10

type DummyScene struct {
	scene.Scene
	nextSceneIdx int
	backImgList  []*ebiten.Image
	rand         *system.Random
	imgIdx       int
	keyboard     *system.Keyboard
}

func (scene *DummyScene) Init() error {

	backList := []*ebiten.Image{}

	for i := 0; i < backImgNum; i++ {
		imgBackFile, err := assets.Open(fmt.Sprintf("assets/back%d.png", i))
		if err != nil {
			return err
		}
		backList = append(backList, graphic.ReadImageFile(imgBackFile))
	}
	scene.backImgList = backList

	scene.imgIdx = 0
	scene.rand = system.NewRandom()

	scene.keyboard = system.GetKeyboard()

	return nil
}

func (scene *DummyScene) Reset() error {

	scene.imgIdx = scene.rand.GetRangeInt(backImgNum)
	return nil
}

func (scene *DummyScene) Update(isWiping bool) (int, int, error) {

	nextScene := -1
	wipeIdx := -1

	if !isWiping {
		scene.keyboard.Listen()
		keys := scene.keyboard.GetPressedKeys()
		if slices.Contains(keys, ebiten.KeySpace) {
			nextScene = scene.nextSceneIdx
			wipeIdx = WIPE_FADEINOUT_IDX
		}
	}

	return nextScene, wipeIdx, nil
}

func (scene *DummyScene) Draw(screen *ebiten.Image) {
	graphic.DrawBackImage(screen, scene.backImgList[scene.imgIdx])
}
