package ai

import (
	"board"
	"fmt"
	"math/rand"
	"time"
)

//SimpleAi structure
type SimpleAi struct {
	Name      string
	Sign      string
	WinCond   []string
	gameBoard *board.GameBoard
}

//NewSimpleAi creates simple SimpleAi
func NewSimpleAi(pName string, pSign string, pGameBoard *board.GameBoard) *SimpleAi {
	winCond := []string{}
	for i := 0; i < pGameBoard.Dx; i++ {
		winCond = append(winCond, pSign)
	}
	return &SimpleAi{Name: pName, Sign: pSign, WinCond: winCond, gameBoard: pGameBoard}
}

//MakeMove function
func (ai *SimpleAi) MakeMove() {
	rField := ai.gameBoard.FreeFields[0]
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	maxIdx := len(ai.gameBoard.FreeFields) - 1
	if maxIdx > 0 {
		rField = ai.gameBoard.FreeFields[random.Intn(maxIdx)]
	}
	ai.gameBoard.FillField(rField.Dx, rField.Dy, ai.Sign)
}

//Pretty print Ai
func (ai SimpleAi) String() string {
	return fmt.Sprintf("SimpleAi.Name: %s, SimpleAi.WinCond: %s", ai.Name, ai.WinCond)
}
