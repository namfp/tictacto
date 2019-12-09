package montecarlo

import (
	"fmt"
	"testing"
	"tictactoe/game"
	"time"
)

func TestStateCopy(t *testing.T) {
	state := MCTS{0.0, 0, &game.DataGame{Self: true, UBoard: game.EmptyUltimateBoard(),
		BoardResult: game.EmptyBoard(), LastMove: game.MoveCoordinate{BoardCoordinate: -1, Coordinate: -1}},
		make([]MCTS, 0), nil}
	durationLimit, _ := time.ParseDuration("100ms")
	RunMCTS(&state, time.Now().Add(durationLimit))
	x, y := ChooseBestMove(&state)
	fmt.Println(x, y)

}