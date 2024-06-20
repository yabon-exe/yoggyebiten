package input

import (
	"math"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type TouchPad struct {
	pinchDistance  float64
	isPinching     bool
	touchPoint     model.Vertex[int]
	isDragging     bool
	dragStart      model.Vertex[int]
	dragEnd        model.Vertex[int]
	flickDirection model.Vector2d
}

var touchPad *TouchPad
var onceTouchPad sync.Once

func GetTouchPad() *TouchPad {
	onceTouchPad.Do(func() {
		touchPad = &TouchPad{}
	})
	return touchPad
}

func (t *TouchPad) Listen() {
	touchIDs := ebiten.TouchIDs()
	// ui.TouchIDが使えない、とかいう意味不明な状況になっているため
	// ids := ebiten.AppendTouchIDs(t.touchId)

	// タップ箇所の取得
	for _, id := range touchIDs {
		t.touchPoint.Set(ebiten.TouchPosition(id))
	}

	// ピンチイン・ピンチアウトの取得
	if len(touchIDs) == 2 {
		x1, y1 := ebiten.TouchPosition(touchIDs[0])
		x2, y2 := ebiten.TouchPosition(touchIDs[1])
		dx := float64(x2 - x1)
		dy := float64(y2 - y1)
		dist := math.Sqrt(dx*dx + dy*dy)

		t.pinchDistance = dist
		t.isPinching = true
	} else {
		t.isPinching = false
	}

	// ドラッグ入力のベクトルとフリック入力の方向
	if len(touchIDs) > 0 {
		id := touchIDs[0]
		x, y := ebiten.TouchPosition(id)

		if !t.isDragging {
			t.dragStart.Set(x, y)
			t.isDragging = true
		} else {
			t.dragEnd.Set(x, y)
			t.flickDirection.Set(float64(t.dragEnd.X-t.dragStart.X), float64(t.dragEnd.Y-t.dragStart.Y))
		}
	} else {
		t.isDragging = false
	}
}
