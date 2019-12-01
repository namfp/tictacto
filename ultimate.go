package main

import (
	"fmt"
)

type MoveCoordinate struct {
	boardCoordinate int
	coordinate int
}


type UltimateState struct {
	self bool
	result float64
	ultimateBoard UltimateBoard
	lastMove MoveCoordinate // value equal -1 when there is no lastMove
}

type UltimateBoard = [9]Board

func otherPlayer(player int) int {
	if player == SELF {
		return OPPONENT
	} else if player == OPPONENT {
		return SELF
	} else {
		panic(fmt.Sprintf("Player should only be SELF or OPPONENT, got %v", player))
	}
}

func findWinner(board *Board) int {
	if checkWinner(SELF, board) {
		return SELF
	} else if checkWinner(OPPONENT, board) {
		return OPPONENT
	} else {
		return EMPTY
	}
}

/*
Return SELF, OPPONENT or EMPTY means no winner
 */
func findWinnerUltimate(ultimateBoard *UltimateBoard) int {
	board := emptyBoard()
	for i, b:= range ultimateBoard {
		board[i] = findWinner(&b)
	}
	return findWinner(&board)
}

func nextBoards(player int, board *Board) ([]Board, []int) {
	var boards []Board
	var coordinates []int
	for i:= range board {
		if i == EMPTY {
			newBoard := *board
			newBoard[i] = player
			boards = append(boards, newBoard)
			coordinates = append(coordinates, i)
		}
	}
	return boards, coordinates
}

func player(self bool) int {
	if self {
		return SELF
	} else {
		return OPPONENT
	}
}

func possibleUltimateState(state *UltimateState, nextBoard int) []UltimateState {
	var result []UltimateState
	possibleBoards, coordinates := nextBoards(player(state.self), &state.ultimateBoard[nextBoard])
	for j, p := range possibleBoards {
		cloned := *state
		cloned.ultimateBoard[nextBoard] = p
		cloned.self = !cloned.self
		cloned.lastMove = MoveCoordinate{nextBoard, coordinates[j]}
		result = append(result, cloned)
	}
	return result
}

func findNextPossibilitiesUltimate(state *UltimateState) []UltimateState {
	var result []UltimateState
	if state.lastMove.boardCoordinate == -1 {
		for i:= range state.ultimateBoard {
			for _, s:= range possibleUltimateState(state, i) {
				result = append(result, s)
			}
		}
	} else {
		nextBoard := state.lastMove.coordinate
		for _, s:= range possibleUltimateState(state, nextBoard) {
			result = append(result, s)
		}
	}
	return []UltimateState{}
}

func scoreGameState(ultimateBoard *UltimateBoard) float32 {
	b := emptyBoard()
	for i, b:= range ultimateBoard {
		b[i] = findWinner(&b)
	}
	score := float32(0.0)
	for _, v := range b {
		if v == SELF {
			score += 1
		} else if v == OPPONENT {
			score -= 1
		}
	}
	return score
}


//func minimax(gameState UltimateState) UltimateState {
//
//}

