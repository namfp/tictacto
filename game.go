package main

const (
	SELF = iota
	OPPONENT = iota
	EMPTY = iota
)

type gameState struct {
	self bool
	result float32
	board Board
	bestMove *gameState
}

type Board = [9]int

func emptyBoard() Board {
	var board [9]int
	for i := 0; i < 9; i++ {
		board[i] = EMPTY
	}
	return board
}

func checkRow(player int, row int, board *Board) bool {
	for i:= row* 3; i < row * 3 + 3; i++ {
		if board[i] != player {
			return false
		}
	}
	return true
}

func checkColumn(player int, column int, board *Board) bool {
	for j:= column; j < 9; j += 3 {
		if board[j] != player {
			return false
		}
	}
	return true
}

func checkFirstDiagonal(player int, board *Board) bool {
	for i:=0; i<9; i +=4 {
		if board[i] != player {
			return false
		}
	}
	return true
}

func checkSecondDiagonal(player int, board *Board) bool {
	for i:=2; i <= 6; i+=2 {
		if board[i] != player {
			return false
		}
	}
	return true
}



func checkWinner(player int, board *Board) bool {
	for i := 0; i < 3; i++ {
		if checkRow(player, i, board) || checkColumn(player, i, board) {
			return true
		}
	}

	if checkFirstDiagonal(player, board) || checkSecondDiagonal(player, board) {
		return true
	}
	return false
}

func computeBoardResult(board *Board) float32 {
	if checkWinner(SELF, board) {
		return 1.0
	} else if checkWinner(OPPONENT, board) {
		return -1.0
	} else {
		return 0.0
	}
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