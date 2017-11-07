package main

import (
	"board"
	"fmt"
	"game"
)

func main() {
	gameBoard := board.NewGameBoard(3)
	game := game.NewGame(gameBoard)
	gameResult := game.Run()
	fmt.Println(gameResult)
}
