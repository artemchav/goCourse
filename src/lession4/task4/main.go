package main

import (
	"fmt"
	"structs"
)

func main() {

	board := &structs.ChessBoard{}
	chessKnight := board.NewChessKnight(0, 0)
	fmt.Println(chessKnight.GetAvailableTurns()) //[[1 2] [2 1]]

	chessKnight.SetPosition(1, 2)
	fmt.Println(chessKnight.GetAvailableTurns()) //[[0 4] [0 0] [2 0] [2 4] [3 3] [3 1]]

	chessKnight.SetPosition(2, 4)
	fmt.Println(chessKnight.GetAvailableTurns()) //[[1 6] [1 2] [3 2] [3 6] [0 5] [0 3] [4 5] [4 3]]

	chessKnight.SetPosition(0, 5)
	fmt.Println(chessKnight.GetAvailableTurns()) //[[1 3] [1 7] [2 6] [2 4]]

}
