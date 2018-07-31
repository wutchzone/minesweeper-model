package minesweeper

import (
	"math/rand"
	"time"
)

var (
	gameField  []Tile
	isGameOver bool
)

type Tile struct {
	X      uint
	Y      uint
	Type   string
	Hidden bool
	Around uint
}

func GenerateGame(xSize, ySize, numberOfMines uint) []Tile {
	gameField := []Tile{}
	minesLeft := numberOfMines

	if xSize*ySize > numberOfMines {
		// ERROR
	}

	for x := uint(0); x < xSize; x++ {
		for y := uint(0); y < ySize; y++ {
			t := "normal"
			if minesLeft != 0 {
				t = "mine"
			}
			gameField = append(gameField, Tile{X: x, Y: y, Type: t, Hidden: true, Around: 0})
		}
	}

	return gameField
}

func GetTile(x, y uint) Tile {
	for _, item := range gameField {
		if item.X == x && item.Y == y {
			return item
		}
	}
	// ERROR
	return Tile{}
}

func RevealItem(x, y uint) Tile {
	for _, i := range gameField {
		if i.X == x && i.Y == y {
			if i.Type == "mine" {
				isGameOver = true
			} else {
				isGameOver = false
			}
			return i
		}
	}

	// ERROR
	return Tile{}
}

func GetBoard() []Tile {
	return gameField
}

func GetGameStatus() bool {
	return isGameOver
}

func generateMines() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for {

	}
}
