package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"golang.org/x/image/font/gofont/goregular"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
	// "github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	scrWidth           = 1280
	scrHeight          = 720
	leftBorder         = 200
	rightBorder        = 20
	topBorder          = 20
	bottomBorder       = 20
	gameAreaLeftEdge   = leftBorder
	gameAreaRightEdge  = scrWidth - rightBorder
	gameAreaTopEdge    = topBorder
	gameAreaBottomEdge = scrHeight - bottomBorder
	gridSqrSize        = 20
	gameGridWidth      = int((gameAreaRightEdge - gameAreaLeftEdge) / gridSqrSize)
	gameGridHeight     = int((gameAreaBottomEdge - gameAreaTopEdge) / gridSqrSize)
)

var gameFont font.Face

type Game struct {
	gameGrid [gameGridHeight][gameGridHeight]bool
}

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
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")

	screen.Fill(color.Gray{200})
	// vector.DrawFilledRect(screen, 0, 0, scrWidth, scrHeight, color.Gray{200}, false)
	text.Draw(screen, "Game", gameFont, 50, 70, color.Black)
	text.Draw(screen, "of", gameFont, 80, 110, color.Black)
	text.Draw(screen, "Life", gameFont, 70, 150, color.Black)

	vector.DrawFilledRect(screen, gameAreaLeftEdge, gameAreaTopEdge, gameAreaRightEdge-gameAreaLeftEdge, gameAreaBottomEdge-gameAreaTopEdge, color.White, false)
	vector.StrokeRect(screen, gameAreaLeftEdge-1, gameAreaTopEdge-1, gameAreaRightEdge-gameAreaLeftEdge+1, gameAreaBottomEdge-gameAreaTopEdge+1, 2, color.Gray{100}, false)
	for i := 1; i < gameGridWidth; i++ {
		vector.StrokeLine(screen, float32(gameAreaLeftEdge+i*gridSqrSize), gameAreaTopEdge, float32(gameAreaLeftEdge+i*gridSqrSize), gameAreaBottomEdge-1, 1, color.Gray{220}, false)
	}
	for i := 1; i < gameGridHeight; i++ {
		vector.StrokeLine(screen, gameAreaLeftEdge, float32(gameAreaTopEdge+i*gridSqrSize), gameAreaRightEdge-1, float32(gameAreaTopEdge+i*gridSqrSize), 1, color.Gray{220}, false)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scrWidth, scrHeight
}

func main() {
	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle("Game of life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
