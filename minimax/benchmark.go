package minimax

import (
	"fmt"
	. "tictactoe/game"
	"time"
)

func Bench() {
	state := UltimateState{0.0, nil, 0,
		&DataGame{true, EmptyUltimateBoard(),
			EmptyBoard(), MoveCoordinate{-1, -1}}}
	for {
		start := time.Now()
		next := play(&state, 5)
		t := time.Now()
		elapsed := t.Sub(start)


		if next != nil {
			x, y := computeMove(next.gameData.LastMove)
			Move(state.gameData, x, y)
			fmt.Println(y, x, state.nbEvaluations, elapsed)// Write action to stdout
		} else {
			fmt.Println(state.result)
			fmt.Println("Finished")
			break
		}
	}
}