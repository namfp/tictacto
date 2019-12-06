package main

import (
	"fmt"
	"math"
	"math/rand"
)

type MoveCoordinate struct {
	boardCoordinate int
	coordinate int
}


type UltimateState struct {
	self bool
	result float64
	ultimateBoard UltimateBoard
	boardResult Board
	lastMove MoveCoordinate // value equal -1 when there is no lastMove
	bestMove *UltimateState
	nbEvaluations int
}


type UltimateBoard = [9]Board

func emptyUltimateBoard() UltimateBoard {
	var board UltimateBoard
	for i := range board {
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}
	return board
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
	for i, v := range board {
		if v == EMPTY {
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
		cloned.boardResult[nextBoard] = getWinner(&cloned.ultimateBoard[nextBoard])
		cloned.self = !cloned.self
		cloned.lastMove = MoveCoordinate{nextBoard, coordinates[j]}
		cloned.bestMove = nil
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
	} else if state.boardResult[state.lastMove.coordinate] != EMPTY {
		for i:= range state.ultimateBoard {
			if state.boardResult[i] == EMPTY {
				for _, s:= range possibleUltimateState(state, i) {
					result = append(result, s)
				}
			}
		}
	} else {
		nextBoard := state.lastMove.coordinate
		for _, s:= range possibleUltimateState(state, nextBoard) {
			result = append(result, s)
		}
	}
	return result
}




func setResult(ultimateState *UltimateState, result float64, nbEvaluation int) {
	ultimateState.result = result
	ultimateState.bestMove = nil
	ultimateState.nbEvaluations = nbEvaluation
}

func alphaBeta(ultimateState *UltimateState, depth int, alpha float64, beta float64) {
	winner := findWinnerUltimate(&ultimateState.ultimateBoard)
	if winner == SELF {
		setResult(ultimateState, 10000, 1)
	} else if winner == OPPONENT {
		setResult(ultimateState, -10000, 1)
	} else if depth == 0 {
		setResult(ultimateState, scoreGameState(&ultimateState.ultimateBoard), 1)
	} else {
		nextPossibilities := findNextPossibilitiesUltimate(ultimateState)
		if len(nextPossibilities) == 0 {
			setResult(ultimateState, 0, 1)
		} else if ultimateState.self {
			value := math.Inf(-1)
			var bestMove *UltimateState
			nbEvaluations := 1
			for i, s := range nextPossibilities {
				alphaBeta(&s, depth - 1, alpha, beta)
				nbEvaluations += s.nbEvaluations
				if s.result > value {
					value = s.result
					bestMove = &nextPossibilities[i]
				}

				alpha = math.Max(alpha, value)
				if alpha >= beta {
					ultimateState.result = value
					ultimateState.bestMove = &nextPossibilities[i]
					ultimateState.nbEvaluations = nbEvaluations
					break
				}
			}
			ultimateState.result = value
			ultimateState.bestMove = bestMove
			ultimateState.nbEvaluations = nbEvaluations

		} else {
			value := math.Inf(1)
			var bestMove *UltimateState
			nbEvaluations := 1
			for i, s := range nextPossibilities {
				alphaBeta(&s, depth - 1, alpha, beta)
				nbEvaluations += s.nbEvaluations
				value = math.Min(value, s.result)
				beta = math.Min(beta, value)

				if s.result < value {
					value = s.result
					bestMove = &nextPossibilities[i]
				}

				if alpha >= beta {
					ultimateState.result = value
					ultimateState.bestMove = &nextPossibilities[i]
					ultimateState.nbEvaluations = nbEvaluations
					break
				}
			}
			ultimateState.result = value
			ultimateState.bestMove = bestMove
			ultimateState.nbEvaluations = nbEvaluations
		}
	}
}

func play(ultimateState *UltimateState, depth int) *UltimateState {
	alphaBeta(ultimateState, depth, math.Inf(-1), math.Inf(1))
	if ultimateState.bestMove == nil {
		// create a random move
		possibleMoves := findNextPossibilitiesUltimate(ultimateState)
		nb := len(possibleMoves)
		var chosen UltimateState
		if nb != 0 {
			chosen = possibleMoves[rand.Int31n(int32(nb))]
			return &chosen
		} else {
			return nil
		}
	} else {
		return ultimateState.bestMove
	}
}

func fullBoard(board *Board)  bool {
	for i:= range board {
		if board[i] == EMPTY {
			return false
		}
	}
	return true
}

func getWinner(board *Board) int {
	if checkWinner(SELF, board) {
		return SELF
	} else if checkWinner(OPPONENT, board) {
		return OPPONENT
	} else if fullBoard(board) {
		return FULL
	} else {
		return EMPTY
	}
}



func moveUltimate(state *UltimateState, i int, j int) {
	lastMove := toBoardCoordinate(i, j)
	state.ultimateBoard[lastMove.boardCoordinate][lastMove.coordinate] = player(state.self)
	state.self = !state.self
	state.lastMove = lastMove
	state.bestMove = nil
	// Check winner to update resultBoard
	state.boardResult[lastMove.boardCoordinate] = getWinner(&state.ultimateBoard[lastMove.boardCoordinate])
}

func convertCoordinate(oneDim int) (int, int) {
	return oneDim % 3, oneDim / 3
}

func computeMove(move MoveCoordinate) (int, int) {
	boardX, boardY := convertCoordinate(move.boardCoordinate)
	i, j := convertCoordinate(move.coordinate)
	return boardX * 3 + i, boardY * 3 + j
}

func toBoardCoordinate(i int, j int) MoveCoordinate {
	boardi := i / 3
	boardj := j / 3
	boardCoordinate := boardj* 3 + boardi
	cellCoordinate := (j % 3) * 3 + i % 3
	return MoveCoordinate{boardCoordinate, cellCoordinate}
}

func runUltimate() {
	state := UltimateState{true, 0.0,emptyUltimateBoard(), emptyBoard(),
		MoveCoordinate{-1, -1}, nil, 0}
	turn := 0
	for {
		turn++
		var opponentRow, opponentCol int
		_, _ = fmt.Scan(&opponentRow, &opponentCol)

		var validActionCount int
		_, _ = fmt.Scan(&validActionCount)

		for i := 0; i < validActionCount; i++ {
			var row, col int
			_, _ = fmt.Scan(&row, &col)
		}

		if opponentRow == -1 && opponentCol == -1 {
			state.self = true
		} else {
			state.self = false
			moveUltimate(&state, opponentCol, opponentRow)
		}
		var depth int
		if turn < 10 {
			depth = 5
		} else if turn < 20 {
			depth = 6
		} else {
			depth = 7
		}
		next := play(&state, depth)
		if next != nil {
			x, y := computeMove(next.lastMove)
			moveUltimate(&state, x, y)
			fmt.Println(y, x)// Write action to stdout
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")


	}
}