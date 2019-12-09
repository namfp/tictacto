package main

import (
	"fmt"
	"tictactoe/game"
	"tictactoe/minimax"
	"tictactoe/montecarlo"
	"time"
)

func isGameFinished(board *game.UltimateBoard) bool {
	winner := game.FindWinnerUltimate(board)
	if winner == game.SELF {
		fmt.Println("MonteCarlo won")
		return true
	} else if winner == game.OPPONENT {
		fmt.Println("Minimax won")
		return true
	} else {
		gameOver := true
		for _, b := range board {
			for _, c := range b {
				if c == game.EMPTY {
					gameOver = false
				}
			}
		}
		if gameOver {
			fmt.Println("equal!")
			return true
		}
	}
	return false
}

func match() {
	minimaxState := minimax.StartUltimateState()
	s := montecarlo.StartStateMCTS()
	monteCarloState := &s
	monteCarloState.GameData.Self = false
	durationLimit, _ := time.ParseDuration("1s")
	turn := 0
	for {
		turn ++
		fmt.Println("turn", turn)
		next := minimax.Play(&minimaxState, 5)
		col, row := game.ComputeMove(next.GameData.LastMove)
		game.Move(minimaxState.GameData, col, row)
		monteCarloState = montecarlo.AfterMove(monteCarloState, col, row)
		fmt.Println("Minmax move", row, col) // Write action to stdout
		if isGameFinished(&monteCarloState.GameData.UBoard) {
			break
		}
		montecarlo.RunMCTS(monteCarloState, time.Now().Add(durationLimit))
		col2, row2 := montecarlo.ChooseBestMove(monteCarloState)
		monteCarloState = montecarlo.AfterMove(monteCarloState, col2, row2)
		game.Move(minimaxState.GameData, col2, row2)
		fmt.Println("MontCarlo move", row2, col2)
		if isGameFinished(&monteCarloState.GameData.UBoard) {
			break
		}



	}
}


func main() {
	match()
}