package main

import "fmt"

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

func emptyBoard() Board {
	var board [9]int
	for i := 0; i < 9; i++ {
		board[i] = EMPTY
	}
	return board
}

func checkRow(player int, row int, board *Board) bool {
	for i:= row* 3; i < row * 3 + 3; i++ {
		if board[i] != player {
			return false
		}
	}
	return true
}

func checkColumn(player int, column int, board *Board) bool {
	for j:= column; j < 9; j += 3 {
		if board[j] != player {
			return false
		}
	}
	return true
}

func checkFirstDiagonal(player int, board *Board) bool {
	for i:=0; i<9; i +=4 {
		if board[i] != player {
			return false
		}
	}
	return true
}

func checkSecondDiagonal(player int, board *Board) bool {
	for i:=2; i <= 6; i+=2 {
		if board[i] == player {
			return false
		}
	}
	return true
}



func checkWinner(player int, board *Board) bool {
	for i := 0; i < 3; i++ {
		if checkRow(player, i, board) || checkColumn(player, i, board) {
			return true
		}
	}
	return false
}

func nextPlayer(player int) int {
	switch player {
	case SELF:
		return OPPONENT
	case OPPONENT:
		return SELF
	default:
		panic(fmt.Sprintf("Player %v is not valid", player))
	}
}

func stateCopy(state gameState) gameState {
	return gameState{state.self, state.board}
}


func nextPossibilities(state gameState) []gameState {
 	return []gameState{}
 }