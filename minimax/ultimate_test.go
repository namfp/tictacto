
package minimax

import (
	"fmt"
	"math"
	"testing"
)


func gameStateTest() UltimateState {
	gameState := UltimateState{true, 0.0, emptyUltimateBoard(), emptyBoard(),
		MoveCoordinate{-1, -1}, nil, 1}
	return gameState
}

func TestFindNextPossibilities(t *testing.T) {

	gameState := gameStateTest()
	allPossibilites := findNextPossibilitiesUltimate(&gameState)
	if len(allPossibilites) != 81 {
		t.Errorf("allPossibilites should be 81, got %v", len(allPossibilites))
	}
}


func TestAlphaBeta1(t *testing.T) {
	gameState := gameStateTest()
	alphaBeta(&gameState, 6, math.Inf(-1), math.Inf(1))
}

func TestPlayFunc(t *testing.T) {
	gameState := gameStateTest()
	play(&gameState, 5)
}

func TestMoveUltimate(t *testing.T) {
	gameState := gameStateTest()
	moveUltimate(&gameState, 8, 8)
	compare := MoveCoordinate{8, 8}
	if gameState.lastMove != compare {
		t.Errorf("lastMove should be %v, got %v", compare, gameState.lastMove)
	}

	if gameState.self {
		t.Errorf("self should be false")
	}
}


func TestMoveUltimate2(t *testing.T) {
	gameState := gameStateTest()
	moveUltimate(&gameState, 4, 6)
	compare := MoveCoordinate{7, 1}
	if gameState.lastMove != compare {
		t.Errorf("lastMove should be %v, got %v", compare, gameState.lastMove)
	}

	if gameState.self {
		t.Errorf("self should be false")
	}

}

type Cd struct {
	x int
	y int
}

func TestFindNextPossibilitiesUltimate(t *testing.T) {
	played := [...]Cd{
		Cd{4,5 },
		Cd{5, 8},
		Cd{8, 6},
		Cd{8, 0},
		Cd{8, 2},
		Cd{8, 8},
		Cd{6, 6},
		Cd{0, 2},
		Cd{1, 7},
		Cd{5, 5},
		Cd{6, 7},
		Cd{1, 3},
		Cd{5, 1},
		Cd{8, 3},
		Cd{7, 2},
		Cd{5, 6},
		Cd{7, 1},
		Cd{3, 5},
		Cd{1, 8},
		Cd{5, 7}}

	gameState := gameStateTest()

	for _, c := range played {
		moveUltimate(&gameState, c.y, c.x)
	}
	tests := findNextPossibilitiesUltimate(&gameState)
	if len(tests) != 8 {
		t.Errorf("lenfth must be 8, got %v", len(tests))
	}
}

func TestConversionCoordinate(t *testing.T) {
	c := toBoardCoordinate(5, 7)
	if c !=( MoveCoordinate{7, 5}) {
		t.Errorf("%v not correct", c)
	}
}

func TestComputeMove(t *testing.T) {
	x, y := computeMove(MoveCoordinate{4, 2})
	if x != 5 && y != 3 {
		t.Errorf("x should be 5 and y should be 3, got %v and %v instead", x, y)
	}
}

func TestScoreBoard(t *testing.T) {

	played := [...]Cd{Cd{4,  5}, Cd{4,  8}, Cd{5,  8}, Cd{7,  8}, Cd{3,  7}, Cd{0,  4},
		Cd{1,  4}, Cd{4,  4}, Cd{5,  5}, Cd{8,  8}, Cd{7,  7}, Cd{4,  3},
		Cd{5,  1}, Cd{8,  3}, Cd{8,  2}, Cd{6,  8}, Cd{1,  7}, Cd{3,  4},
		Cd{0,  5}, Cd{0,  6}, Cd{1,  2}, Cd{4,  7}, Cd{5,  4}, Cd{6,  3},
		Cd{0,  2}, Cd{1,  6}, Cd{3,  1}, Cd{0,  3}, Cd{0,  1}, Cd{1,  3},
		Cd{4,  2}, Cd{4,  6}, Cd{4,  1}, Cd{3,  3}, Cd{2,  2}, Cd{2,  3},
		Cd{7,  2}, Cd{2,  6}}
	gameState := gameStateTest()

	for _, c := range played {
		moveUltimate(&gameState, c.y, c.x)
	}

	score := scoreGameState(&gameState.ultimateBoard)
	fmt.Println(score)

}