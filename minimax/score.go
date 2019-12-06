package minimax

import . "tictactoe/game"

func scoreGameState(ultimateBoard *UltimateBoard) float64 {
	scoreSelf := computeMiniBoardScore(ultimateBoard, SELF)
	scoreOpponent := computeMiniBoardScore(ultimateBoard, OPPONENT)
	return scoreSelf - scoreOpponent
}


func computeMiniBoardScore(ultimateBoard *UltimateBoard, player int) float64 {
	globalBoard := EmptyBoard()
	score := 0.0
	for i, board:= range ultimateBoard {
		globalBoard[i] = FindWinner(&board)
		if i == 4 && globalBoard[i] == EMPTY {
			for c := range board {
				if c == player {
					score += 3 // add 3 for each played at the center
				}
			}
		}
		score += computeConsecutiveWinningScore(&board, player) * 2
	}

	for i, v := range globalBoard {
		if v == player {
			score += 5
			if i == 4 { // center board
				score += 10
			}
			if i == 0 || i == 2 || i == 6 || i == 8 {
				score += 3
			}
		}
	}
	score += computeConsecutiveWinningScore(&globalBoard, player) * 4
	return score
}

func computeConsecutiveWinningScoreColumn(board *Board, player int) float64 {
	score := 0.0
	for i := 0; i < 3; i++ { // count columns
		count := 0
		isCount := true
		for j := 0; j < 3; j++ {
			v := board[j * 3 + i]
			if v == player {
				count += 1
			} else if v != EMPTY {
				isCount = false
			}
		}
		if isCount && count == 2{
			score += 1
		}
	}
	return score
}

func computeConsecutiveWinningScoreLine(board *Board, player int) float64 {
	score := 0.0
	for j := 0; j < 3; j++ {
		count := 0
		isCount := true
		for i := 0; i < 3; i++ {
			v := board[j * 3 + i]
			if v == player {
				count += 1
			} else if v != EMPTY {
				isCount = false
			}
		}
		if isCount && count == 2{
			score += 1
		}
	}
	return score
}

func computeConsecutiveWinningScoreFirstDiagonal(board *Board, player int) float64 {
	count := 0
	isCount := true
	score := 0.0
	for i:=0; i<9; i +=4 {
		if board[i] == player {
			count += 1
		} else if board[i] != EMPTY {
			isCount = false
		}
	}
	if isCount && count == 2{
		score += 1
	}
	return score
}


func computeConsecutiveWinningScoreSecondDiagonal(board *Board, player int) float64 {
	count := 0
	isCount := true
	score := 0.0

	for i:=2; i<=6; i += 2 {
		if board[i] == player {
			count += 1
		} else if board[i] != EMPTY {
			isCount = false
		}
	}
	if isCount && count == 2{
		score += 1
	}
	return score
}

func computeConsecutiveWinningScore(board *Board, player int) float64 {
	scoreColumns := computeConsecutiveWinningScoreColumn(board, player)
	scoreLines := computeConsecutiveWinningScoreLine(board, player)
	scoreFirstDiagonal := computeConsecutiveWinningScoreFirstDiagonal(board, player)
	scoreSecondDiagonal := computeConsecutiveWinningScoreSecondDiagonal(board, player)
	return scoreColumns + scoreLines + scoreFirstDiagonal + scoreSecondDiagonal
}