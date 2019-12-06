package game


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


func CheckWinner(player int, board *Board) bool {
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

func FindWinner(board *Board) int {
	if CheckWinner(SELF, board) {
		return SELF
	} else if CheckWinner(OPPONENT, board) {
		return OPPONENT
	} else {
		return EMPTY
	}
}

/*
Return SELF, OPPONENT or EMPTY means no winner
*/
func FindWinnerUltimate(ultimateBoard *UltimateBoard) int {
	board := EmptyBoard()
	for i, b:= range ultimateBoard {
		board[i] = FindWinner(&b)
	}
	return FindWinner(&board)
}

func ComputeBoardResult(board *Board) float32 {
	if CheckWinner(SELF, board) {
		return 1.0
	} else if CheckWinner(OPPONENT, board) {
		return -1.0
	} else {
		return 0.0
	}
}