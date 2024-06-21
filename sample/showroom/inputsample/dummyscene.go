package inputsample

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/object/message"
	"github.com/yabon-exe/yoggyebiten/game/scene"
	"github.com/yabon-exe/yoggyebiten/game/system/input"
)

type DummyScene struct {
	scene.Scene
	keyboard          *input.Keyboard
	mouse             *input.Mouse
	labelKeyBoard     *message.SimpleMessage
	infoKeyBoard      *message.SimpleMessage
	labelMouse        *message.SimpleMessage
	infoMouseCursol   *message.SimpleMessage
	infoMouseOn       *message.SimpleMessage
	infoMousePressed  *message.SimpleMessage
	infoMouseReleased *message.SimpleMessage
	infoMouseDragInfo *message.SimpleMessage
}

func (scene *DummyScene) Init() error {

	scene.keyboard = input.GetKeyboard()
	scene.mouse = input.GetMouse()

	scene.labelKeyBoard = createMsg(20, 0)
	scene.labelKeyBoard.SetText("KeyBoard:")
	scene.infoKeyBoard = createMsg(30, 20)

	scene.labelMouse = createMsg(20, 50)
	scene.labelMouse.SetText("Mouse:")
	scene.infoMouseCursol = createMsg(30, 70)
	scene.infoMouseOn = createMsg(30, 90)
	scene.infoMousePressed = createMsg(30, 110)
	scene.infoMouseReleased = createMsg(30, 130)
	scene.infoMouseDragInfo = createMsg(30, 150)

	return nil
}

func (scene *DummyScene) Reset() error {

	return nil
}

func (scene *DummyScene) Update(isWiping bool) (int, int, error) {

	scene.keyboard.Listen()
	scene.mouse.Listen()

	onKeys := scene.keyboard.GetPressedKeys()
	onKeyStrs := []string{}
	for _, key := range onKeys {
		onKeyStrs = append(onKeyStrs, key.String())
	}

	mousePoint := scene.mouse.GetMousePoint()
	mouseOnL, mouseOnR := scene.mouse.GetOnButtonInfo()
	pressedL, pressedR := scene.mouse.GetPressedInfo()
	releasedL, releasedR := scene.mouse.GetReleasedInfo()
	isDragging, dragStart, dragEnd := scene.mouse.GetDragInfo()

	scene.infoKeyBoard.SetText(fmt.Sprintf("key=>%s", strings.Join(onKeyStrs, ",")))
	scene.infoMouseCursol.SetText(fmt.Sprintf("mouse cursol=> %d, %d", mousePoint.X, mousePoint.Y))
	scene.infoMouseOn.SetText(fmt.Sprintf("mouse on=> L:%5t, R:%5t", mouseOnL, mouseOnR))
	scene.infoMousePressed.SetText(fmt.Sprintf("mouse pressed=> L:%5t, R:%5t", pressedL, pressedR))
	scene.infoMouseReleased.SetText(fmt.Sprintf("mouse released=> L:%5t, R:%5t", releasedL, releasedR))
	scene.infoMouseDragInfo.SetText(fmt.Sprintf("mouse drag=> %t (%d,%d)->(%d,%d)",
		isDragging, dragStart.X, dragStart.Y, dragEnd.X, dragEnd.Y))

	return -1, -1, nil
}

func (scene *DummyScene) Draw(screen *ebiten.Image) {
	scene.labelKeyBoard.Draw(screen)
	scene.infoKeyBoard.Draw(screen)
	scene.labelMouse.Draw(screen)
	scene.infoMouseCursol.Draw(screen)
	scene.infoMouseOn.Draw(screen)
	scene.infoMousePressed.Draw(screen)
	scene.infoMouseReleased.Draw(screen)
	scene.infoMouseDragInfo.Draw(screen)
}

func createMsg(x int, y int) *message.SimpleMessage {
	msg := &message.SimpleMessage{}
	msg.Init()
	msg.SetColor(0, 255, 0, 255)
	msg.SetSize(20)
	msg.SetPosition(model.Vertex[float64]{X: float64(x), Y: float64(y)})
	return msg
}
