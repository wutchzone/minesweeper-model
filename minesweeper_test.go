package minesweeper

import (
	"testing"
)

func TestGenerateGame(t *testing.T) {
	var x uint = 5
	var y uint = 5
	board, err := GenerateGame(x, y, 5)

	if err == nil {
		if uint(len(board)) != x*y {
			t.Errorf("Wront array length. Have %v expected %v", len(board), x*y)
		}
	} else {
		t.Error(err)
	}
}
