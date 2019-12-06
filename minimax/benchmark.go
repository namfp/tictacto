package minimax

import (
	"fmt"
	"tictactoe/game"
	"time"
)

func Bench() {
	state := UltimateState{true, 0.0, game.EmptyUltimateBoard(), game.EmptyBoard(),
		game.MoveCoordinate{-1, -1}, nil, 1}
	for {
		start := time.Now()
		next := play(&state, 5)
		t := time.Now()
		elapsed := t.Sub(start)


		if next != nil {
			x, y := computeMove(next.lastMove)
			moveUltimate(&state, x, y)
			fmt.Println(y, x, state.nbEvaluations, elapsed)// Write action to stdout
		} else {
			fmt.Println(state.result)
			fmt.Println("Finished")
			break
		}
	}
}