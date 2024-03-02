package main

import (
	"log"

	"golang.org/x/image/font/opentype"

	"golang.org/x/image/font/gofont/goregular"
)

func init() {
	f, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size: 32,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}

	gameFont = face

	face, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size: 22,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}
	startButton = button{
		posX:            40,
		posY:            300,
		w:               100,
		h:               40,
		label:           "Start",
		labelFont:       face,
		colorLabel:      colorBtnLabel,
		colorBackground: colorBtnBG,
		colorEdgeline:   colorBtnLine,
	}
	clearButton = button{
		posX:            40,
		posY:            500,
		w:               100,
		h:               40,
		label:           "Clear",
		labelFont:       face,
		colorLabel:      colorBtnLabel,
		colorBackground: colorBtnBG,
		colorEdgeline:   colorBtnLine,
	}
	randomButton = button{
		posX:            40,
		posY:            400,
		w:               100,
		h:               40,
		label:           "Random",
		labelFont:       face,
		colorLabel:      colorBtnLabel,
		colorBackground: colorBtnBG,
		colorEdgeline:   colorBtnLine,
	}
}
