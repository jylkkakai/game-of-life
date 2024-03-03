package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
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
	gridSqrSize        = 5
	gameGridWidth      = int((gameAreaRightEdge - gameAreaLeftEdge) / gridSqrSize)
	gameGridHeight     = int((gameAreaBottomEdge - gameAreaTopEdge) / gridSqrSize)
)

var (
	gameFont     font.Face
	startButton  button
	clearButton  button
	randomButton button
	tickCounter  int = 0
	delay        int = 20
)

type Game struct {
	gameGrid [gameGridHeight][gameGridWidth]bool
	gameIsOn bool
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	// log.Println(startButton)
	if startButton.isHovered(x, y) {
		startButton.setHovered()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			startButton.setClicked()
			if g.gameIsOn {
				startButton.label = "Start"
				g.gameIsOn = false
				randomButton.toggleDisabled()
				clearButton.toggleDisabled()
			} else {
				startButton.label = "Stop"
				g.gameIsOn = true
				randomButton.toggleDisabled()
				clearButton.toggleDisabled()
			}
		}
	} else {
		startButton.reSet()
	}
	if clearButton.isHovered(x, y) {
		clearButton.setHovered()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !clearButton.isDisabled {
			clearButton.setClicked()
			var newGrid [gameGridHeight][gameGridWidth]bool
			g.gameGrid = newGrid
		}
	} else {
		clearButton.reSet()
	}
	if randomButton.isHovered(x, y) {
		randomButton.setHovered()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !randomButton.isDisabled {
			randomButton.setClicked()
			rndCounter := int(gameGridHeight * gameGridWidth / 10)
			for rndCounter > 0 {
				rndX := rand.Intn(gameGridHeight)
				rndY := rand.Intn(gameGridWidth)
				if !g.gameGrid[rndX][rndY] {
					g.gameGrid[rndX][rndY] = true
					rndCounter--
				}
			}
		}
	} else {
		randomButton.reSet()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		startButton.reSet()
		clearButton.reSet()
		randomButton.reSet()
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if x < gameAreaRightEdge && x > gameAreaLeftEdge && y < gameAreaBottomEdge && y > gameAreaTopEdge {
			g.gameGrid[int((y-topBorder)/gridSqrSize)][int((x-leftBorder)/gridSqrSize)] = true

		}
	}

	if g.gameIsOn {
		if tickCounter == delay {
			var newGrid [gameGridHeight][gameGridWidth]bool
			for i := 0; i < gameGridHeight; i++ {
				for j := 0; j < gameGridWidth; j++ {

					numOfNeighbours := getNumOfNeighbours(g.gameGrid, j, i)

					if numOfNeighbours == 3 || numOfNeighbours == 2 && g.gameGrid[i][j] {
						newGrid[i][j] = true
					}
				}
			}
			g.gameGrid = newGrid
			tickCounter = 0
		}
		tickCounter++
	}

	return nil
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func getNumOfNeighbours(arr [gameGridHeight][gameGridWidth]bool, x, y int) int {

	var num int
	for i := max(0, y-1); i <= min(y+1, gameGridHeight-1); i++ {
		for j := max(0, x-1); j <= min(x+1, gameGridWidth-1); j++ {

			if arr[i][j] && !(x == j && y == i) {
				num++
			}
		}
	}
	return num
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.Gray{200})
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
	for i := 0; i < gameGridHeight; i++ {
		for j := 0; j < gameGridWidth; j++ {
			if g.gameGrid[i][j] {
				vector.DrawFilledRect(screen, float32(gameAreaLeftEdge+j*gridSqrSize), float32(gameAreaTopEdge+i*gridSqrSize), gridSqrSize, gridSqrSize, color.Gray{50}, false)
			}
		}
	}
	startButton.draw(screen)
	clearButton.draw(screen)
	randomButton.draw(screen)

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
