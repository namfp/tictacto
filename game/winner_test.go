package game

import (
	"testing"
)

func TestCheckWinner(t *testing.T) {
	board := [9]int{
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY}
	isSelfWinner := CheckWinner(SELF, &board)
	if isSelfWinner {
		t.Errorf("Self is not winner")
	}

	isOpponentWinner := CheckWinner(SELF, &board)
	if isOpponentWinner {
		t.Errorf("Opponent is not winner")
	}
}

func TestCheckWinner1(t *testing.T) {
	board := [9]int{
		EMPTY, EMPTY, OPPONENT,
		SELF, SELF, SELF,
		EMPTY, EMPTY, OPPONENT}
	isSelfWinner := CheckWinner(SELF, &board)

	if !isSelfWinner {
		t.Errorf("Self must be winner")
	}

	isOpponentWinner := CheckWinner(OPPONENT, &board)
	if isOpponentWinner {
		t.Errorf("Opponent is not winner")
	}
}

func TestCheckDiagonal1(t *testing.T) {
	board := [9]int{
		SELF, EMPTY, OPPONENT,
		EMPTY, SELF, OPPONENT,
		EMPTY, EMPTY, SELF}
	isSelfWinner := CheckWinner(SELF, &board)

	if !isSelfWinner {
		t.Errorf("Self must be winner")
	}

	isOpponentWinner := CheckWinner(OPPONENT, &board)
	if isOpponentWinner {
		t.Errorf("Opponent is not winner")
	}
}

func TestCheckDiagonal2(t *testing.T) {
	board := [9]int{
		SELF, EMPTY, OPPONENT,
		EMPTY, OPPONENT, SELF,
		OPPONENT, EMPTY, SELF}
	isSelfWinner := CheckWinner(SELF, &board)

	if isSelfWinner {
		t.Errorf("Self is not winner")
	}

	isOpponentWinner := CheckWinner(OPPONENT, &board)
	if !isOpponentWinner {
		t.Errorf("Opponent is winner")
	}
}

