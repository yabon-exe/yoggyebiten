package input

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
)

type Mouse struct {
	mousePoint    model.Vertex[int]
	onLeft        bool
	onRight       bool
	pressedLeft   bool
	pressedRight  bool
	releasedLeft  bool
	releasedRight bool
	dragging      bool
	dragStart     model.Vertex[int]
	dragEnd       model.Vertex[int]
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

	m.pressedLeft = false
	m.pressedRight = false
	m.releasedLeft = false
	m.releasedRight = false

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !m.onLeft {
			m.onLeft = true
			m.pressedLeft = true
			m.dragStart.Set(m.mousePoint.Get())
		}
	} else {
		if m.onLeft {
			m.onLeft = false
			m.releasedLeft = true
			m.dragging = false
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if !m.onRight {
			m.pressedRight = true
		}
		m.onRight = true
	} else {
		if m.onRight {
			m.releasedRight = true
		}
		m.onRight = false
	}

	if m.onLeft {
		m.dragging = true
		m.dragEnd.Set(m.mousePoint.Get())
	}
}

func (m *Mouse) GetMousePoint() model.Vertex[int] {
	return m.mousePoint
}

func (m *Mouse) GetOnButtonInfo() (bool, bool) {
	return m.onLeft, m.onRight
}

func (m *Mouse) GetPressedInfo() (bool, bool) {
	return m.pressedLeft, m.pressedRight
}

func (m *Mouse) GetReleasedInfo() (bool, bool) {
	return m.releasedLeft, m.releasedRight
}

func (m *Mouse) GetDragInfo() (bool, model.Vertex[int], model.Vertex[int]) {
	return m.dragging, m.dragStart, m.dragEnd
}
