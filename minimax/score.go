package minimax

import . "tictactoe/game"

type ScoreTable struct {
	PlayedCenterBoard float64
	PlayedCenter float64
	ConsWinning       float64
	ConsWinningBoard  float64
	WinCenterBoard    float64
	WindEdge          float64
	WinBoard          float64
}

func scoreGameState(ultimateBoard *UltimateBoard, scoreTable *ScoreTable) float64 {
	scoreSelf := computeMiniBoardScore(ultimateBoard, SELF, scoreTable)
	scoreOpponent := computeMiniBoardScore(ultimateBoard, OPPONENT, scoreTable)
	return scoreSelf - scoreOpponent
}


func computeMiniBoardScore(ultimateBoard *UltimateBoard, player int, scoreTable *ScoreTable) float64 {
	globalBoard := EmptyBoard()
	score := 0.0
	for i, board:= range ultimateBoard {
		globalBoard[i] = FindWinner(&board)
		if i == 4 && globalBoard[i] == EMPTY {
			for c := range board {
				if c == player {
					score += scoreTable.PlayedCenterBoard // add 3 for each played at the center
				}
			}
		}
		if board[4] == player {
			score += scoreTable.PlayedCenterBoard
		}
		score += computeConsecutiveWinningScore(&board, player) * scoreTable.ConsWinning
	}

	for i, v := range globalBoard {
		if v == player {
			score += scoreTable.WinBoard
			if i == 4 { // center board
				score += scoreTable.WinCenterBoard
			}
			if i == 0 || i == 2 || i == 6 || i == 8 {
				score += scoreTable.WindEdge
			}
		}
	}
	score += computeConsecutiveWinningScore(&globalBoard, player) * scoreTable.ConsWinningBoard
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
		if isCount && count == 2 {
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