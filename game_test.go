package main

import (
	"testing"
)

func TestCheckWinner(t *testing.T) {
	board := [9]int{
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY}
	isSelfWinner := checkWinner(SELF, &board)
	if isSelfWinner {
		t.Errorf("Self is not winner")
	}

	isOpponentWinner := checkWinner(SELF, &board)
	if isOpponentWinner {
		t.Errorf("Opponent is not winner")
	}
}

func TestCheckWinner1(t *testing.T) {
	board := [9]int{
		EMPTY, EMPTY, OPPONENT,
		SELF, SELF, SELF,
		EMPTY, EMPTY, OPPONENT}
	isSelfWinner := checkWinner(SELF, &board)

	if !isSelfWinner {
		t.Errorf("Self must be winner")
	}

	isOpponentWinner := checkWinner(OPPONENT, &board)
	if isOpponentWinner {
		t.Errorf("Opponent is not winner")
	}
}

func TestCheckDiagonal1(t *testing.T) {
	board := [9]int{
		SELF, EMPTY, OPPONENT,
		EMPTY, SELF, OPPONENT,
		EMPTY, EMPTY, SELF}
	isSelfWinner := checkWinner(SELF, &board)

	if !isSelfWinner {
		t.Errorf("Self must be winner")
	}

	isOpponentWinner := checkWinner(OPPONENT, &board)
	if isOpponentWinner {
		t.Errorf("Opponent is not winner")
	}
}

func TestCheckDiagonal2(t *testing.T) {
	board := [9]int{
		SELF, EMPTY, OPPONENT,
		EMPTY, OPPONENT, SELF,
		OPPONENT, EMPTY, SELF}
	isSelfWinner := checkWinner(SELF, &board)

	if isSelfWinner {
		t.Errorf("Self is not winner")
	}

	isOpponentWinner := checkWinner(OPPONENT, &board)
	if !isOpponentWinner {
		t.Errorf("Opponent is winner")
	}
}


func TestStateCopy(t *testing.T) {
	board := [9]int{
		EMPTY, EMPTY, OPPONENT,
		SELF, SELF, SELF,
		EMPTY, EMPTY, OPPONENT}
	state := gameState{true, 0, board, nil}
	copiedState := stateCopy(&state)
	state.board[0] = SELF
	if copiedState.board[0] == SELF {
		t.Errorf("Copied failed")
	}
}

func TestMiniMax1(t *testing.T) {
	board := [9]int{
		OPPONENT, EMPTY, OPPONENT,
		SELF, SELF, EMPTY,
		SELF, OPPONENT, OPPONENT}
	state := gameState{true, 0, board, nil}
	minimax(&state)
	if state.result != 1.0 {
		t.Errorf("Must win")
	}

	if state.bestMove.board[5] != SELF {
		t.Errorf("It must play the position 5")
	}
}

func TestMiniMax2(t *testing.T) {
	board := [9]int{
		EMPTY, EMPTY, OPPONENT,
		SELF, SELF, EMPTY,
		EMPTY, EMPTY, OPPONENT}
	state := gameState{true, 0, board, nil}
	minimax(&state)
	if state.result != 1.0 {
		t.Errorf("Must win")
	}

	if state.bestMove.board[5] != SELF {
		t.Errorf("It must play the position 5")
	}
}


func TestCopy(t *testing.T) {
	type New struct {
		a *int
		b [][]string
	}

	first:=New{new(int), [][]string{[]string{"x"}, []string{"y"}}}
	*first.a = 0
	second := first
	*first.a = 1

	if &second == &first {
		t.Errorf("second and first is not the same")
	}

	if *second.a != *first.a {
		t.Errorf("second and first is not the same")
	}
	second.a = new(int)
	*second.a = 2
	if *second.a == *first.a {
		t.Errorf("second and first is the same)")
	}

	if &first.b == &second.b {
		t.Errorf("should not be equal")
	}

	second.b[0] = make([]string, 0)
	if &first.b[0] == &second.b[0] {
		t.Errorf("should be equal")
	}

}

