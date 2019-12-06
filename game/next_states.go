package game



type DataGame struct {
	Self        bool
	UBoard      UltimateBoard
	BoardResult Board
	LastMove    MoveCoordinate
}

func nextBoards(player int, board *Board) ([]Board, []int) {
	var boards []Board
	var coordinates []int
	for i, v := range board {
		if v == EMPTY {
			newBoard := *board
			newBoard[i] = player
			boards = append(boards, newBoard)
			coordinates = append(coordinates, i)
		}
	}
	return boards, coordinates
}


func Player(self bool) int {
	if self {
		return SELF
	} else {
		return OPPONENT
	}
}


func possibleUltimateState(state *DataGame, nextBoard int) []*DataGame {
	var result []*DataGame
	possibleBoards, coordinates := nextBoards(Player(state.Self), &state.UBoard[nextBoard])
	for j, p := range possibleBoards {
		cloned := *state
		cloned.UBoard[nextBoard] = p
		cloned.BoardResult[nextBoard] = GetWinner(&cloned.UBoard[nextBoard])
		cloned.Self = !state.Self
		cloned.LastMove = MoveCoordinate{nextBoard, coordinates[j]}
		result = append(result, &cloned)
	}
	return result
}



func FindNextPossibilities(state *DataGame) []*DataGame {
	var result []*DataGame

	if state.LastMove.BoardCoordinate == -1 {
		for i:= range state.UBoard {
			for _, s:= range possibleUltimateState(state, i) {
				result = append(result, s)
			}
		}
	} else if state.BoardResult[state.LastMove.Coordinate] != EMPTY {
		for i:= range state.UBoard {
			if state.BoardResult[i] == EMPTY {
				for _, s:= range possibleUltimateState(state, i) {
					result = append(result, s)
				}
			}
		}
	} else {
		nextBoard := state.LastMove.Coordinate
		for _, s:= range possibleUltimateState(state, nextBoard) {
			result = append(result, s)
		}
	}
	return result
}

func toBoardCoordinate(i int, j int) MoveCoordinate {
	boardi := i / 3
	boardj := j / 3
	boardCoordinate := boardj* 3 + boardi
	cellCoordinate := (j % 3) * 3 + i % 3
	return MoveCoordinate{boardCoordinate, cellCoordinate}
}

func Move(state *DataGame, i int, j int) {
	lastMove := toBoardCoordinate(i, j)
	state.UBoard[lastMove.BoardCoordinate][lastMove.Coordinate] = Player(state.Self)
	state.Self = !state.Self
	state.LastMove = lastMove
	// Check winner to update resultBoard
	state.BoardResult[lastMove.BoardCoordinate] = GetWinner(&state.UBoard[lastMove.BoardCoordinate])
}