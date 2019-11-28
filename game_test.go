package main

import "testing"

func TestGet(t *testing.T) {
	board := emptyBoard()
	first := get(0, 0, board)
	if first != 2 {
		t.Errorf("first value is incorrect: %d, want: %d.", first, 2)
	}
}