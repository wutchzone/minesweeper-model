package minesweeper

var (
	gameField []Tile
)

type Tile struct {
	X    uint
	Y    uint
	Type string
}

func GenerateGame(xSize, ySize, numberOfMines uint) []Tile {
	gameField := []Tile{}
	if xSize*ySize > numberOfMines {
		// ERROR
	}

	for x := uint(0); x < xSize; x++ {
		for y := uint(0); y < ySize; y++ {
			gameField = append(gameField, Tile{X: x, Y: y, Type: "mine"})
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
	return Tile{}
}

func RevealItem(x, y uint) Tile {

}
