package main

import "testing"

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

func TestStateCopy(t *testing.T) {
	board := [9]int{
		EMPTY, EMPTY, OPPONENT,
		SELF, SELF, SELF,
		EMPTY, EMPTY, OPPONENT}
	state := gameState{true, board}
	copiedState := stateCopy(state)
	state.board[0] = SELF
	if copiedState.board[0] == SELF {
		t.Errorf("Copied failed")
	}
}