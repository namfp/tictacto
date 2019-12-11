package minimax

import (
	"fmt"
	"math"
	"math/rand"
	. "tictactoe/game"
)

type UltimateState struct {
	Result        float64
	bestMove      *UltimateState
	NbEvaluations int
	GameData      *DataGame
}

func StartUltimateState() UltimateState {
	return  UltimateState{0.0, nil, 0,
		&DataGame{Self: true, UBoard: EmptyUltimateBoard(),
			BoardResult: EmptyBoard(), LastMove: MoveCoordinate{BoardCoordinate: -1, Coordinate: -1}}}
}

func setResult(ultimateState *UltimateState, result float64, nbEvaluation int) {
	ultimateState.Result = result
	ultimateState.bestMove = nil
	ultimateState.NbEvaluations = nbEvaluation
}

func createState(ultimateState *UltimateState, game *DataGame) UltimateState {
	cloned := *ultimateState
	cloned.GameData = game
	cloned.bestMove = nil
	cloned.Result = 0
	cloned.NbEvaluations = 0
	return cloned
}


func alphaBeta(ultimateState *UltimateState, depth int, alpha float64, beta float64, scoreTable *ScoreTable) {


	winner := FindWinnerUltimate(&ultimateState.GameData.UBoard)
	if winner == SELF {
		setResult(ultimateState, 10000, 1)
	} else if winner == OPPONENT {
		setResult(ultimateState, -10000, 1)
	} else if depth == 0 {
		setResult(ultimateState, scoreGameState(&ultimateState.GameData.UBoard, scoreTable), 1)
	} else {
		nextPossibilities := FindNextPossibilities(ultimateState.GameData)
		if len(nextPossibilities) == 0 {
			setResult(ultimateState, 0, 1)
		} else if ultimateState.GameData.Self {
			value := math.Inf(-1)
			var bestMove *UltimateState
			nbEvaluations := 1
			for _, s := range nextPossibilities {
				state := createState(ultimateState, s)
				alphaBeta(&state, depth - 1, alpha, beta, scoreTable)
				nbEvaluations += state.NbEvaluations
				if state.Result > value {
					value = state.Result
					bestMove = &state
				}

				alpha = math.Max(alpha, value)
				if alpha >= beta {
					ultimateState.Result = value
					ultimateState.bestMove = &state
					ultimateState.NbEvaluations = nbEvaluations
					break
				}
			}
			ultimateState.Result = value
			ultimateState.bestMove = bestMove
			ultimateState.NbEvaluations = nbEvaluations

		} else {
			value := math.Inf(1)
			var bestMove *UltimateState
			nbEvaluations := 1
			for _, s := range nextPossibilities {
				state := createState(ultimateState, s)
				alphaBeta(&state, depth - 1, alpha, beta, scoreTable)
				nbEvaluations += state.NbEvaluations

				if state.Result < value {
					value = state.Result
					bestMove = &state
				}

				beta = math.Min(beta, value)

				if alpha >= beta {
					ultimateState.Result = value
					ultimateState.bestMove = &state
					ultimateState.NbEvaluations = nbEvaluations
					break
				}
			}
			ultimateState.Result = value
			ultimateState.bestMove = bestMove
			ultimateState.NbEvaluations = nbEvaluations
		}
	}
}


func Play(ultimateState *UltimateState, depth int, scoreTable *ScoreTable) *UltimateState {
	alphaBeta(ultimateState, depth, math.Inf(-1), math.Inf(1), scoreTable)
	if ultimateState.bestMove == nil {
		// create a random move
		possibleMoves := FindNextPossibilities(ultimateState.GameData)
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


func runUltimate() {
	state := UltimateState{0.0, nil, 0,
		&DataGame{Self: true, UBoard: EmptyUltimateBoard(),
			BoardResult: EmptyBoard(), LastMove: MoveCoordinate{BoardCoordinate: -1, Coordinate: -1}}}
	scoreTable := ScoreTable{
		PlayedCenterBoard: 3,
		ConsWinning:       0,
		ConsWinningBoard:  0,
		WinCenterBoard:    10,
		WindEdge:          3,
		WinBoard:          5,
	}
	for {
		var opponentRow, opponentCol int
		_, _ = fmt.Scan(&opponentRow, &opponentCol)

		var validActionCount int
		_, _ = fmt.Scan(&validActionCount)

		for i := 0; i < validActionCount; i++ {
			var row, col int
			_, _ = fmt.Scan(&row, &col)
		}

		if opponentRow == -1 && opponentCol == -1 {
			state.GameData.Self = true
		} else {
			state.GameData.Self = false
			Move(state.GameData, opponentCol, opponentRow)
		}
		next := Play(&state, 4, scoreTable)
		if next != nil {
			x, y := ComputeMove(next.GameData.LastMove)
			Move(state.GameData, x, y)
			fmt.Println(y, x)// Write action to stdout
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")


	}
}