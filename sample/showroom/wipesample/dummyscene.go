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

const backImgNum = 14

type DummyScene struct {
	scene.Scene
	nextSceneIdx int
	backImgList  []*ebiten.Image
	rand         *system.Random
	imgIdx       int
	keyboard     *system.Keyboard
	startIdx     int
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

	scene.imgIdx = scene.startIdx
	scene.rand = system.NewRandom()

	scene.keyboard = system.GetKeyboard()

	return nil
}

func (scene *DummyScene) Reset() error {

	nextIdx := (scene.imgIdx + 2)
	if nextIdx >= backImgNum {
		nextIdx = nextIdx - backImgNum
	}
	scene.imgIdx = nextIdx
	return nil
}

func (scene *DummyScene) Update(isWiping bool) (int, int, error) {

	nextScene := -1
	wipeIdx := -1

	if !isWiping {
		scene.keyboard.Listen()
		keys := scene.keyboard.GetPressedKeys()
		if slices.Contains(keys, ebiten.KeyQ) {
			nextScene = scene.nextSceneIdx
			wipeIdx = WIPE_FADEINOUT_IDX
		} else if slices.Contains(keys, ebiten.KeyW) {
			nextScene = scene.nextSceneIdx
			wipeIdx = WIPE_CURTAIN_U_IDX
			if slices.Contains(keys, ebiten.KeyArrowUp) {
				wipeIdx = WIPE_CURTAIN_U_IDX
			} else if slices.Contains(keys, ebiten.KeyArrowDown) {
				wipeIdx = WIPE_CURTAIN_D_IDX
			} else if slices.Contains(keys, ebiten.KeyArrowLeft) {
				wipeIdx = WIPE_CURTAIN_L_IDX
			} else if slices.Contains(keys, ebiten.KeyArrowRight) {
				wipeIdx = WIPE_CURTAIN_R_IDX
			}
		} else if slices.Contains(keys, ebiten.KeyE) {
			nextScene = scene.nextSceneIdx
			wipeIdx = WIPE_CLEANER_U_IDX
			if slices.Contains(keys, ebiten.KeyArrowUp) {
				wipeIdx = WIPE_CLEANER_U_IDX
			} else if slices.Contains(keys, ebiten.KeyArrowDown) {
				wipeIdx = WIPE_CLEANER_D_IDX
			} else if slices.Contains(keys, ebiten.KeyArrowLeft) {
				wipeIdx = WIPE_CLEANER_L_IDX
			} else if slices.Contains(keys, ebiten.KeyArrowRight) {
				wipeIdx = WIPE_CLEANER_R_IDX
			}
		} else if slices.Contains(keys, ebiten.KeyR) {
			nextScene = scene.nextSceneIdx
			wipeIdx = WIPE_PICTURE_UL_IDX
			if slices.Contains(keys, ebiten.KeyArrowUp) {
				if slices.Contains(keys, ebiten.KeyArrowLeft) {
					wipeIdx = WIPE_PICTURE_UL_IDX
				} else if slices.Contains(keys, ebiten.KeyArrowRight) {
					wipeIdx = WIPE_PICTURE_UR_IDX
				} else {
					wipeIdx = WIPE_PICTURE_U_IDX
				}
			} else if slices.Contains(keys, ebiten.KeyArrowDown) {
				if slices.Contains(keys, ebiten.KeyArrowLeft) {
					wipeIdx = WIPE_PICTURE_DL_IDX
				} else if slices.Contains(keys, ebiten.KeyArrowRight) {
					wipeIdx = WIPE_PICTURE_DR_IDX
				} else {
					wipeIdx = WIPE_PICTURE_D_IDX
				}
			} else if slices.Contains(keys, ebiten.KeyArrowLeft) {
				wipeIdx = WIPE_PICTURE_L_IDX
			} else if slices.Contains(keys, ebiten.KeyArrowRight) {
				wipeIdx = WIPE_PICTURE_R_IDX
			}
		}
	}

	return nextScene, wipeIdx, nil
}

func (scene *DummyScene) Draw(screen *ebiten.Image) {
	graphic.DrawBackImage(screen, scene.backImgList[scene.imgIdx])
}
