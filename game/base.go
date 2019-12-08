package game
type UltimateBoard = [9]Board

const (
	SELF = iota
	OPPONENT = iota
	EMPTY = iota
	FULL = iota
)

type Board = [9]int

func EmptyBoard() Board {
	var board [9]int
	for i := 0; i < 9; i++ {
		board[i] = EMPTY
	}
	return board
}

type MoveCoordinate struct {
	BoardCoordinate int
	Coordinate int
}

func EmptyUltimateBoard() UltimateBoard {
	var board UltimateBoard
	for i := range board {
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}
	return board
}

func GetWinner(board *Board) int {
	if CheckWinner(SELF, board) {
		return SELF
	} else if CheckWinner(OPPONENT, board) {
		return OPPONENT
	} else if fullBoard(board) {
		return FULL
	} else {
		return EMPTY
	}
}

func fullBoard(board *Board)  bool {
	for i:= range board {
		if board[i] == EMPTY {
			return false
		}
	}
	return true
}

func convertCoordinate(oneDim int) (int, int) {
	return oneDim % 3, oneDim / 3
}

func ComputeMove(move MoveCoordinate) (int, int) { //col then row
	boardX, boardY := convertCoordinate(move.BoardCoordinate)
	i, j := convertCoordinate(move.Coordinate)
	return boardX * 3 + i, boardY * 3 + j
}