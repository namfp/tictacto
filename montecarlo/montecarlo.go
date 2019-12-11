package montecarlo

import (
	"fmt"
	"math"
	"math/rand"
	. "tictactoe/game"
	"time"
)

type MCTS struct {
	value     float64
	nbVisited int
	GameData  *DataGame
	children  []MCTS
	parent    *MCTS
	isLeafNode bool
	isChecked bool
}

func uct(node *MCTS) float64 {
	return node.value/float64(node.nbVisited) +
		5*math.Sqrt(math.Log(float64(node.parent.nbVisited))/float64(node.nbVisited))
}

func backPropagation(node *MCTS, value float64) {
	node.value += value
	node.nbVisited += 1
	if node.parent == nil {
		return
	} else {
		backPropagation(node.parent, value)
	}
}

func RunMCTS(node *MCTS, timeLimit time.Time) {
	for {
		now := time.Now()
		if now.After(timeLimit) {
			return
		} else {
			selected := selecting(node)
			v := rollout(selected)
			backPropagation(selected, v)
		}
	}
}

func AfterMove(node *MCTS, col int, row int) *MCTS {
	lastMove := ToBoardCoordinate(col, row)
	var found *MCTS
	for i, c := range node.children {
		if c.GameData.LastMove == lastMove {
			found = &node.children[i]
		}
	}
	if found == nil {
		Move(node.GameData, col, row)
		return node
	}
	return found
}

func ChooseBestMove(node *MCTS) (int, int) {
	chosen := &node.children[0]
	for i, c := range node.children {
		if c.nbVisited > chosen.nbVisited {
			chosen = &node.children[i]
		}
	}
	return ComputeMove(chosen.GameData.LastMove)
}

func StartStateMCTS() MCTS {
	return MCTS{0.0, 0, &DataGame{Self: true, UBoard: EmptyUltimateBoard(),
		BoardResult: EmptyBoard(), LastMove: MoveCoordinate{BoardCoordinate: -1, Coordinate: -1}},
		make([]MCTS, 0), nil, false, false}
}

func MonteCarloGamePlay() {
	s := StartStateMCTS()
	state := &s
	turn := 0
	for {
		turn++
		var opponentRow, opponentCol int
		_, _ = fmt.Scan(&opponentRow, &opponentCol)

		var validActionCount int
		_, _ = fmt.Scan(&validActionCount)

		for i := 0; i < validActionCount; i++ {
			var row, col int
			_, _ = fmt.Scan(&row, &col)
		}

		if opponentRow == -1 && opponentCol == -1 {
			state.GameData.Self = true
		} else {
			state.GameData.Self = false
			state = AfterMove(state, opponentCol, opponentRow)
		}

		var durationLimit time.Duration
		if turn == 1 {
			durationLimit, _ = time.ParseDuration("0.99s")
		} else {
			durationLimit, _ = time.ParseDuration("99ms")
		}
		RunMCTS(state, time.Now().Add(durationLimit))
		x, y := ChooseBestMove(state)
		state = AfterMove(state, x, y)
		fmt.Println(y, x)
	}
}

func expand(node *MCTS) {
	nextPossibilites := FindNextPossibilities(node.GameData)
	for i := range nextPossibilites {
		// check if terminated node
		child := MCTS{0.0, 0, nextPossibilites[i], make([]MCTS, 0),
			node, false, false}
		updateResult(&child)
		node.children = append(node.children, child)
	}
}

func gameResult(board *UltimateBoard) int {
	winner := FindWinnerUltimate(board)
	if winner == SELF || winner == OPPONENT {
		return winner
	} else {
		gameOver := true
		for _, b := range board {
			for _, c := range b {
				if c == EMPTY {
					gameOver = false
				}
			}
		}
		if gameOver {
			return FULL
		}
	}
	return EMPTY
}

func updateResult(node *MCTS) {
	result := gameResult(&node.GameData.UBoard)
	if result != EMPTY {
		// leaf node
		node.isLeafNode = true
		if result == SELF {
			node.value = 1.0
		} else if result == OPPONENT {
			node.value = -1.0
		} else {
			node.value = 0.0
		}
	}
	node.isChecked = true
}

func selecting(node *MCTS) *MCTS {
	if !node.isChecked {
		updateResult(node)
	}

	if node.isLeafNode {
		return node
	} else if len(node.children) == 0 {
		if node.nbVisited == 0 && node.parent != nil {
			return node
		} else {
			expand(node)
			if len(node.children) == 0 {
				fmt.Println("notpossible")
			}
			return &node.children[0]
		}
	}

	for i, c := range node.children {
		if c.nbVisited == 0 {
			return &node.children[i]
		}
	}

	chosen := &node.children[0]
	currentUCT := uct(chosen)

	for i, c := range node.children {
		uc := uct(&c)
		if currentUCT < uc {
			chosen = &node.children[i]
			currentUCT = uc
		}
	}
	return selecting(chosen)
}

func rollout(node *MCTS) float64 {
	dataGame := node.GameData
	if node.isLeafNode {
		return node.value
	}
	for {
		winner := FindWinnerUltimate(&dataGame.UBoard)
		if winner == SELF {
			return 1.0
		} else if winner == OPPONENT {
			return -1.0
		}

		nextStates := FindNextPossibilities(dataGame)
		nbValues := len(nextStates)
		if nbValues == 0 {
			return 0.0 // Equal
		} else {
			dataGame = nextStates[rand.Intn(len(nextStates))]
		}
	}
}
