
package main

import (
	"math"
	"testing"
)

func TestFindNextPossibilities(t *testing.T) {
	gameState := UltimateState{true, 0.0,emptyUltimateBoard(),
		MoveCoordinate{-1, -1}, nil}
	allPossibilites := findNextPossibilitiesUltimate(&gameState)
	if len(allPossibilites) != 81 {
		t.Errorf("allPossibilites should be 81, got %v", len(allPossibilites))
	}
}


func TestAlphaBeta1(t *testing.T) {
	gameState := UltimateState{true, 0.0,emptyUltimateBoard(),
		MoveCoordinate{-1, -1}, nil}
	alphaBeta(&gameState, 6, math.Inf(-1), math.Inf(1))
}

func TestPlayFunc(t *testing.T) {
	gameState := UltimateState{true, 0.0,emptyUltimateBoard(),
		MoveCoordinate{-1, -1}, nil}
	play(&gameState, 5)
}

func TestMoveUltimate(t *testing.T) {
	gameState := UltimateState{true, 0.0,emptyUltimateBoard(),
		MoveCoordinate{-1, -1}, nil}
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
	gameState := UltimateState{true, 0.0,emptyUltimateBoard(),
		MoveCoordinate{-1, -1}, nil}
	moveUltimate(&gameState, 4, 6)
	compare := MoveCoordinate{7, 1}
	if gameState.lastMove != compare {
		t.Errorf("lastMove should be %v, got %v", compare, gameState.lastMove)
	}

	if gameState.self {
		t.Errorf("self should be false")
	}

}


func TestComputeMove(t *testing.T) {
	x, y := computeMove(MoveCoordinate{4, 2})
	if x != 5 && y != 3 {
		t.Errorf("x should be 5 and y should be 3, got %v and %v instead", x, y)
	}

}