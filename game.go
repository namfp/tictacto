package main

import (
	"fmt"
)

const (
	SELF = iota
	OPPONENT = iota
	EMPTY = iota
	FULL = iota
)

func emptyBoard() Board {
	var board [9]int
	for i := 0; i < 9; i++ {
		board[i] = EMPTY
	}
	return board
}



func stateCopy(state *gameState) gameState {
	return gameState{state.self, state.result,state.board, state.bestMove}
}

func nextPossibilities(state *gameState) []gameState {
	var states []gameState
	for i:=0; i<9; i++ {
		if state.board[i] == EMPTY {
			newState := stateCopy(state)
			if state.self {
				newState.board[i] = SELF
			} else {
				newState.board[i] = OPPONENT
			}
			newState.self = !newState.self
			states = append(states, newState)
		}
	}
 	return states
 }

func minimax(state *gameState) {
	result := computeBoardResult(&state.board)
	allPossibilities := nextPossibilities(state)
	if result == 1.0 || result == -1.0 || len(allPossibilities) == 0 {
		state.result = result
		state.bestMove = nil
	} else {
		if state.self {
			evaluate(allPossibilities, state, func(x float32, y float32) bool {return x > y})
		} else {
			evaluate(allPossibilities, state, func(x float32, y float32) bool {return x < y})
		}
	}
}

func evaluate(allPossibilities []gameState, state *gameState, compare func(float32, float32) bool) {
	for i, s := range allPossibilities {
		minimax(&s)
		if i == 0 || compare(s.result, state.result) {
			state.result = s.result
			state.bestMove = &allPossibilities[i]
		}
	}
}

func getMove(state *gameState, nextState *gameState) (int, int) {
	for i:=0; i<9; i++ {
		if state.board[i] != nextState.board[i] {
			return i / 3, i % 3
		}
	}
	panic("two boards are identical")
}

func move(state *gameState, player int, i int, j int) {
	if state.self && player == OPPONENT {
		panic("It is not the opponent turn!")
	}
	pos := i * 3 + j
	state.board[pos] = player
	state.self = !state.self
}


func testMain() {
	board := emptyBoard()
	state := gameState{true, 0, board, nil}
	for {
		var opponentRow, opponentCol int
		_, _ = fmt.Scan(&opponentRow, &opponentCol)

		var validActionCount int
		_, _ = fmt.Scan(&validActionCount)

		for i := 0; i < validActionCount; i++ {
			var row, col int
			_, _ = fmt.Scan(&row, &col)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		if opponentRow == -1 && opponentCol == -1 {
			state.self = true
		} else {
			state.self = false
			move(&state, OPPONENT, opponentRow, opponentCol)
		}

		minimax(&state)
		myMoveX, myMoveY := getMove(&state, state.bestMove)
		move(&state, SELF, myMoveX, myMoveY)
		fmt.Println(myMoveX, myMoveY)// Write action to stdout
	}
}

