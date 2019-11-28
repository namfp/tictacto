package main

const (
	SELF = iota
	OPPONENT = iota
	EMPTY = iota
)

type gameState struct {
	self bool
	board Board
}

type Board = [9]int

func emptyBoard() *Board {
	var board [9]int
	for i := 0; i < 9; i++ {
		board[i] = EMPTY
	}
	return &board
}

func checkRow(player int, row int, board *Board) bool {
	return false
}

func findWinner(player int, board *Board) *int {
	return nil
}

func get(i int, j int, board *Board) int {
	return board[3 * i + j]
}