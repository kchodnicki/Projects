package game

import (
	"ai"
	"board"
	"fmt"
	"reflect"
)

//Game structure for actual gameplay
type Game struct {
	gameBoard *board.GameBoard
	players   []*ai.SimpleAi
}

//NewGame creates new game structure
func NewGame(pGameBoard *board.GameBoard) *Game {
	playerX := ai.NewSimpleAi("Player X", "X", pGameBoard)
	playerO := ai.NewSimpleAi("Player O", "O", pGameBoard)
	players := []*ai.SimpleAi{playerX, playerO}
	return &Game{gameBoard: pGameBoard, players: players}
}

//Run the game
func (game *Game) Run() string {
	var endGame bool
	var result string
	var currPlayer *ai.SimpleAi
	turn := 0
	for true {
		currPlayer = game.players[turn%2]
		currPlayer.MakeMove()
		fmt.Println(game.gameBoard)
		result, endGame = game.checkWin(currPlayer)
		if endGame {
			break
		}
		turn++
	}
	if result == "Draw" {
		return result
	}
	return fmt.Sprintf("%s won!", result)
}

//checkWin checks if after last move game is won
func (game *Game) checkWin(ai *ai.SimpleAi) (string, bool) {
	diagOneWin := []string{}
	diagTwoWin := []string{}
	for i, rowWin := range game.gameBoard.Board {
		//if line on row
		if reflect.DeepEqual(rowWin, ai.WinCond) {
			return ai.Name, true
		}
		colWin := []string{}
		for j, col := range rowWin {
			colWin = append(colWin, col)
			if i == j {
				diagOneWin = append(diagOneWin, col)
			}
			if i == game.gameBoard.Dx-1-j {
				diagTwoWin = append(diagTwoWin, col)
			}
		}
		//if line on column
		if reflect.DeepEqual(colWin, ai.WinCond) {
			return ai.Name, true
		}
	}
	//if line on diagonal
	if reflect.DeepEqual(diagOneWin, ai.WinCond) || reflect.DeepEqual(diagTwoWin, ai.WinCond) {
		return ai.Name, true
	}
	//if board if full then noone won
	if len(game.gameBoard.FreeFields) == 0 {
		return "Draw", true
	}
	return "", false
}
