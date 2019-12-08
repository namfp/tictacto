package minimax

import (
	"fmt"
	. "tictactoe/game"
	"time"
)

func Bench() {
	state := UltimateState{0.0, nil, 0,
		&DataGame{Self: true, UBoard: EmptyUltimateBoard(),
			BoardResult: EmptyBoard(), LastMove: MoveCoordinate{BoardCoordinate: -1, Coordinate: -1}}}
	for {
		start := time.Now()
		next := play(&state, 5)
		t := time.Now()
		elapsed := t.Sub(start)


		if next != nil {
			x, y := ComputeMove(next.gameData.LastMove)
			Move(state.gameData, x, y)
			fmt.Println(y, x, state.nbEvaluations, elapsed)// Write action to stdout
		} else {
			fmt.Println(state.result)
			fmt.Println("Finished")
			break
		}
	}
}