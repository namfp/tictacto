package minimax

import (
	"fmt"
	. "tictactoe/game"
	"time"
)

func Bench() {
	state := StartUltimateState()
	for {
		start := time.Now()
		next := Play(&state, 5)
		t := time.Now()
		elapsed := t.Sub(start)


		if next != nil {
			x, y := ComputeMove(next.GameData.LastMove)
			Move(state.GameData, x, y)
			fmt.Println(y, x, state.NbEvaluations, elapsed) // Write action to stdout
		} else {
			fmt.Println(state.Result)
			fmt.Println("Finished")
			break
		}
	}
}