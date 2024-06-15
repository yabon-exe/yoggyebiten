package system

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct {
	keys []ebiten.Key
}

var keyboard *Keyboard
var onceKeyboard sync.Once

func GetKeyboard() *Keyboard {
	onceKeyboard.Do(func() {
		keyboard = &Keyboard{}
	})
	return keyboard
}

func (k *Keyboard) Listen() {
	k.keys = inpututil.AppendPressedKeys(k.keys[:0])
}

func (k *Keyboard) GetPressedKeys() []ebiten.Key {
	return k.keys
}
