package minesweeper

import (
	"errors"
	"math/rand"
	"time"
)

var (
	isGameOver bool
	gameField  GameField
)

// Tile for GameField
type Tile struct {
	X      uint
	Y      uint
	Type   string
	Hidden bool
	Around uint
}

// GameField is type of game
type GameField []Tile

// GenerateGame creates new GameField
func GenerateGame(xSize, ySize, numberOfMines uint) (GameField, error) {
	minesLeft := numberOfMines

	if xSize*ySize < numberOfMines {
		return nil, errors.New("Too many mines for game")
	}

	for x := uint(0); x < xSize; x++ {
		for y := uint(0); y < ySize; y++ {
			t := "normal"
			if minesLeft != 0 {
				t = "mine"
				minesLeft--
			}
			gameField = append(gameField, Tile{X: x, Y: y, Type: t, Hidden: true, Around: 0})
		}
	}

	gameField.shuffle()
	gameField.checkForCorrectIndex(xSize, ySize)
	gameField.calculateAround()

	// Make deep copy
	output := make(GameField, len(gameField))
	copy(output, gameField)
	return GameField(output), nil
}

// RevealTile Reveals tile in game
func RevealTile(x, y uint) (Tile, error) {
	for _, i := range gameField {
		if i.X == x && i.Y == y {
			if i.Type == "mine" {
				isGameOver = true
			} else {
				isGameOver = false
			}
			return i, nil
		}
	}

	// ERROR
	return Tile{}, errors.New("Not found")
}

func (gf GameField) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range gf {
		num := r.Intn(len(gf) - 1)

		gf[i], gf[num] = gf[num], gf[i]
	}
}

func (gf GameField) checkForCorrectIndex(xSize uint, ySize uint) {
	x := uint(0)
	y := uint(0)

	for i := range gf {
		gf[i].X = x
		gf[i].Y = y
		if y == ySize {
			y = 0
			x++
		} else {
			y++
		}
	}
}

func (gf GameField) calculateAround() {
	for i := range gf {
		var carry uint
		element := gf[i]

		tilesAround := make([]Tile, 8)

		tilesAround[0], _ = gf.check(element.X-1, element.Y-1)
		tilesAround[1], _ = gf.check(element.X-1, element.Y)
		tilesAround[2], _ = gf.check(element.X-1, element.Y+1)
		tilesAround[3], _ = gf.check(element.X, element.Y-1)
		tilesAround[4], _ = gf.check(element.X, element.Y+1)
		tilesAround[5], _ = gf.check(element.X+1, element.Y-1)
		tilesAround[6], _ = gf.check(element.X+1, element.Y)
		tilesAround[7], _ = gf.check(element.X+1, element.Y+1)

		for _, item := range tilesAround {
			if item.Type == "mine" {
				carry++
			}
		}

		gameField[i].Around = carry
	}
}

func (gf GameField) check(x, y uint) (Tile, error) {
	for _, item := range gameField {
		if item.X == x && item.Y == y {
			return item, nil
		}
	}

	// ERROR
	return Tile{}, errors.New("Not found")
}
