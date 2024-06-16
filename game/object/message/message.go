package message

import (
	"bytes"
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yabon-exe/yoggyebiten/game/model"
	"github.com/yabon-exe/yoggyebiten/game/object"
)

type SimpleMessage struct {
	object.Object
	tf       *text.GoTextFace
	tsf      *text.GoTextFaceSource
	op       *text.DrawOptions
	text     string
	size     float64
	position *model.Vertex[float64]
	color    *color.RGBA
}

func (message *SimpleMessage) Init() error {

	tsf, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	message.tsf = tsf
	if err != nil {
		return err
	}
	message.op = &text.DrawOptions{}
	message.position = &model.Vertex[float64]{}
	message.tf = &text.GoTextFace{
		Source: message.tsf,
		Size:   message.size,
	}
	message.color = &color.RGBA{}
	return nil
}

func (message *SimpleMessage) Update() error {

	return nil
}

func (message *SimpleMessage) Draw(screen *ebiten.Image) {
	message.op.GeoM.Reset()
	message.op.GeoM.Translate(message.position.X, message.position.Y)
	message.op.ColorScale.ScaleWithColor(message.color)
	message.tf.Size = message.size
	text.Draw(screen, message.text, message.tf, message.op)
}

func (message *SimpleMessage) SetPosition(pos model.Vertex[float64]) {
	message.position.Set(pos.X, pos.Y)
}

func (message *SimpleMessage) SetText(text string) {
	message.text = text
}

func (message *SimpleMessage) SetColor(r uint8, g uint8, b uint8, a uint8) {
	message.color.R = r
	message.color.G = g
	message.color.B = b
	message.color.A = a
}

func (message *SimpleMessage) SetSize(size float64) {
	message.size = size
}
