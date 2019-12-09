
package minimax

import (
	"fmt"
	"math"
	"testing"
	. "tictactoe/game"
)


func gameStateTest() UltimateState {
	gameState := UltimateState{0.0, nil, 0,
		&DataGame{true, EmptyUltimateBoard(),
			EmptyBoard(), MoveCoordinate{-1, -1}}}
	return gameState
}



func TestAlphaBeta1(t *testing.T) {
	gameState := gameStateTest()
	alphaBeta(&gameState, 6, math.Inf(-1), math.Inf(1))
}

func TestPlayFunc(t *testing.T) {
	gameState := gameStateTest()
	Play(&gameState, 5)
}

func TestMoveUltimate(t *testing.T) {
	gameData := gameStateTest().GameData
	Move(gameData, 8, 8)
	compare := MoveCoordinate{8, 8}
	if gameData.LastMove != compare {
		t.Errorf("lastMove should be %v, got %v", compare, gameData)
	}

	if gameData.Self {
		t.Errorf("self should be false")
	}
}


type Cd struct {
	x int
	y int
}

func TestScoreBoard(t *testing.T) {

	played := [...]Cd{Cd{4,  5}, Cd{4,  8}, Cd{5,  8}, Cd{7,  8}, Cd{3,  7}, Cd{0,  4},
		Cd{1,  4}, Cd{4,  4}, Cd{5,  5}, Cd{8,  8}, Cd{7,  7}, Cd{4,  3},
		Cd{5,  1}, Cd{8,  3}, Cd{8,  2}, Cd{6,  8}, Cd{1,  7}, Cd{3,  4},
		Cd{0,  5}, Cd{0,  6}, Cd{1,  2}, Cd{4,  7}, Cd{5,  4}, Cd{6,  3},
		Cd{0,  2}, Cd{1,  6}, Cd{3,  1}, Cd{0,  3}, Cd{0,  1}, Cd{1,  3},
		Cd{4,  2}, Cd{4,  6}, Cd{4,  1}, Cd{3,  3}, Cd{2,  2}, Cd{2,  3},
		Cd{7,  2}, Cd{2,  6}}
	gameData := gameStateTest().GameData

	for _, c := range played {
		Move(gameData, c.y, c.x)
	}

	score := scoreGameState(&gameData.UBoard)
	fmt.Println(score)

}