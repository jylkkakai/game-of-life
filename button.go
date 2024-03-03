package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type button struct {
	posX float32
	posY float32
	w    float32
	h    float32

	label     string
	labelFont font.Face

	colorLabel      color.Gray16
	colorBackground color.Gray
	colorEdgeline   color.Gray16

	isClicked  bool
	isDisabled bool
}

var (
	colorBtnBG       = color.Gray{170}
	colorBtnHoverBG  = color.Gray{200}
	colorBtnLine     = color.Black
	colorBtnLabel    = color.Black
	colorBtnDisabled = color.Gray16{40000}
)

func (b *button) isHovered(x, y int) bool {
	return float32(x) > b.posX && float32(x) < b.posX+b.w && float32(y) > b.posY && float32(y) < b.posY+b.h

}
func (b *button) setHovered() {
	b.colorBackground = colorBtnHoverBG
}
func (b *button) setClicked() {
	if !b.isClicked {
		b.isClicked = true
		b.posX = b.posX + 1
		b.posY = b.posY + 1
		b.w = b.w - 2
		b.h = b.h - 2
	}
}
func (b *button) toggleDisabled() {
	if b.isDisabled {
		b.isDisabled = false
		b.colorLabel = colorBtnLabel
		b.colorEdgeline = colorBtnLine
		b.colorBackground = colorBtnBG
	} else {
		b.isDisabled = true
		b.colorLabel = colorBtnDisabled
		b.colorEdgeline = colorBtnDisabled
		b.colorBackground = colorBtnHoverBG

	}
}
func (b *button) reSet() {
	if !b.isDisabled {
		b.colorBackground = colorBtnBG
	}
	if b.isClicked {
		b.isClicked = false
		b.posX = b.posX - 1
		b.posY = b.posY - 1
		b.w = b.w + 2
		b.h = b.h + 2
	}
}
func (b *button) draw(screen *ebiten.Image) {

	vector.DrawFilledRect(screen, b.posX, b.posY, b.w, b.h, b.colorBackground, false)
	// vector.StrokeRect(screen, b.posX, b.posY, b.w, b.h, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX+2, b.posY, b.posX+b.w-2, b.posY, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX+b.w-2, b.posY, b.posX+b.w, b.posY+2, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX+b.w, b.posY+2, b.posX+b.w, b.posY+b.h-2, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX+b.w, b.posY+b.h-2, b.posX+b.w-2, b.posY+b.h, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX+b.w-2, b.posY+b.h, b.posX+2, b.posY+b.h, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX+2, b.posY+b.h, b.posX, b.posY+b.h-2, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX, b.posY+b.h-2, b.posX, b.posY+2, 2, b.colorEdgeline, false)
	vector.StrokeLine(screen, b.posX, b.posY+2, b.posX+2, b.posY, 2, b.colorEdgeline, false)

	labelOffset := (b.w - float32(font.MeasureString(b.labelFont, b.label).Round())) / 2
	text.Draw(screen, b.label, b.labelFont, int(b.posX)+int(labelOffset), int(b.posY)+27, b.colorLabel)
}
