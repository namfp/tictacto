package minimax

import (
	"fmt"
	"math"
	"math/rand"
	. "tictactoe/game"
)

type UltimateState struct {
	result        float64
	bestMove      *UltimateState
	nbEvaluations int
	gameData      *DataGame
}


func setResult(ultimateState *UltimateState, result float64, nbEvaluation int) {
	ultimateState.result = result
	ultimateState.bestMove = nil
	ultimateState.nbEvaluations = nbEvaluation
}

func createState(ultimateState *UltimateState, game *DataGame) UltimateState {
	cloned := *ultimateState
	cloned.gameData = game
	cloned.bestMove = nil
	cloned.result = 0
	cloned.nbEvaluations = 0
	return cloned
}


func alphaBeta(ultimateState *UltimateState, depth int, alpha float64, beta float64) {


	winner := FindWinnerUltimate(&ultimateState.gameData.UBoard)
	if winner == SELF {
		setResult(ultimateState, 10000, 1)
	} else if winner == OPPONENT {
		setResult(ultimateState, -10000, 1)
	} else if depth == 0 {
		setResult(ultimateState, scoreGameState(&ultimateState.gameData.UBoard), 1)
	} else {
		nextPossibilities := FindNextPossibilities(ultimateState.gameData)
		if len(nextPossibilities) == 0 {
			setResult(ultimateState, 0, 1)
		} else if ultimateState.gameData.Self {
			value := math.Inf(-1)
			var bestMove *UltimateState
			nbEvaluations := 1
			for _, s := range nextPossibilities {
				state := createState(ultimateState, s)
				alphaBeta(&state, depth - 1, alpha, beta)
				nbEvaluations += state.nbEvaluations
				if state.result > value {
					value = state.result
					bestMove = &state
				}

				alpha = math.Max(alpha, value)
				if alpha >= beta {
					ultimateState.result = value
					ultimateState.bestMove = &state
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
			for _, s := range nextPossibilities {
				state := createState(ultimateState, s)
				alphaBeta(&state, depth - 1, alpha, beta)
				nbEvaluations += state.nbEvaluations

				if state.result < value {
					value = state.result
					bestMove = &state
				}

				beta = math.Min(beta, value)

				if alpha >= beta {
					ultimateState.result = value
					ultimateState.bestMove = &state
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
		possibleMoves := FindNextPossibilities(ultimateState.gameData)
		nb := len(possibleMoves)
		var chosen UltimateState
		if nb != 0 {
			game := possibleMoves[rand.Int31n(int32(nb))]
			chosen = createState(ultimateState, game)
			return &chosen
		} else {
			return nil
		}
	} else {
		return ultimateState.bestMove
	}
}


func convertCoordinate(oneDim int) (int, int) {
	return oneDim % 3, oneDim / 3
}

func computeMove(move MoveCoordinate) (int, int) {
	boardX, boardY := convertCoordinate(move.BoardCoordinate)
	i, j := convertCoordinate(move.Coordinate)
	return boardX * 3 + i, boardY * 3 + j
}


func runUltimate() {
	state := UltimateState{0.0, nil, 0,
		&DataGame{true, EmptyUltimateBoard(),
			EmptyBoard(), MoveCoordinate{-1, -1}}}
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
			state.gameData.Self = true
		} else {
			state.gameData.Self = false
			Move(state.gameData, opponentCol, opponentRow)
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
			x, y := computeMove(next.gameData.LastMove)
			Move(state.gameData, x, y)
			fmt.Println(y, x)// Write action to stdout
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")


	}
}