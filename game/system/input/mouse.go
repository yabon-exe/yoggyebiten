package input

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type Mouse struct {
	mousePoint model.Vertex[int]
	leftClick  bool
	rightClick bool
	dragging   bool
	dragStart  model.Vertex[int]
	dragEnd    model.Vertex[int]
}

var mouse *Mouse
var onceMouse sync.Once

func GetMouse() *Mouse {
	onceMouse.Do(func() {
		mouse = &Mouse{}
	})
	return mouse
}

func (m *Mouse) Listen() {
	m.mousePoint.Set(ebiten.CursorPosition())

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !m.leftClick {
			m.leftClick = true
			m.dragStart.Set(m.mousePoint.Get())
		}
	} else {
		if m.leftClick {
			m.leftClick = false
			m.dragging = false
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		m.rightClick = true
	} else {
		m.rightClick = false
	}

	if m.leftClick {
		m.dragging = true
		m.dragEnd.Set(m.mousePoint.Get())
	}
}

func (m *Mouse) GetMousePoint() model.Vertex[int] {
	return m.mousePoint
}

func (m *Mouse) GetClickInfo() (bool, bool) {
	return m.leftClick, m.rightClick
}

func (m *Mouse) GetDragInfo() (bool, model.Vertex[int], model.Vertex[int]) {
	return m.dragging, m.dragStart, m.dragEnd
}
