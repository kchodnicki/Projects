package board

import (
	"bytes"
	"fmt"
)

//Field single field reporesentation
type Field struct {
	Dx, Dy int
}

//GameBoard Structure representing tictactoe gameboard
type GameBoard struct {
	Dx         int
	Board      [][]string
	FreeFields []*Field
}

//NewGameBoard is a constructor for GameBoard
func NewGameBoard(x int) *GameBoard {
	newboard := make([][]string, x)
	freefields := []*Field{}
	for i := range newboard {
		newboard[i] = make([]string, x)
		for j := range newboard[i] {
			newboard[i][j] = "-"
			newfield := &Field{Dx: i, Dy: j}
			freefields = append(freefields, newfield)
		}
	}
	return &GameBoard{Dx: x, Board: newboard, FreeFields: freefields}
}

//FillField enters sign under (x,y) and removes field from list of free fields
func (gameboard *GameBoard) FillField(x int, y int, pSign string) {
	gameboard.Board[x][y] = pSign
	gameboard.removeFreeField(x, y)
	return
}

//removeFreeField removes free field from list of avaliable fields
func (gameboard *GameBoard) removeFreeField(x int, y int) {
	for i := range gameboard.FreeFields {
		if gameboard.FreeFields[i].Dx == x && gameboard.FreeFields[i].Dy == y {
			gameboard.FreeFields[i] = gameboard.FreeFields[len(gameboard.FreeFields)-1]
			gameboard.FreeFields = gameboard.FreeFields[:len(gameboard.FreeFields)-1]
			//Field found and removed, we can break the loop
			break
		}
	}
	return
}

//Pretty print GameBoard
func (gameboard GameBoard) String() string {
	var b bytes.Buffer
	for i := 0; i < gameboard.Dx; i++ {
		b.WriteString("\n")
		for j := 0; j < gameboard.Dx; j++ {
			b.WriteString(fmt.Sprintf("%s\t", gameboard.Board[i][j]))
		}
	}
	b.WriteString("\n\n")
	return b.String()
}
