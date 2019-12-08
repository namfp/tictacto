package montecarlo

import (
	"fmt"
	"math"
	"math/rand"
	"tictactoe/game"
	"time"
)

type MCTS struct {
	value float64
	nbVisited int
	gameData *game.DataGame
	children []MCTS
	parent *MCTS
}

func uct(node *MCTS) float64 {
	return node.value / float64(node.nbVisited) +
		1.41 * math.Sqrt(math.Log(float64(node.parent.nbVisited)) / float64(node.nbVisited))

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

func runMCTS(node *MCTS, timeLimit time.Time) {
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

func afterMove(node *MCTS, col int, row int) *MCTS {
	lastMove := game.ToBoardCoordinate(col, row)
	var found *MCTS
	for i, c := range node.children {
		if c.gameData.LastMove == lastMove {
			found = &node.children[i]
		}
	}
	if found == nil {
		game.Move(node.gameData, col, row)
		return node
	}
	return found
}

func chooseBestMove(node *MCTS) (int, int) {
	chosen := &node.children[0]
	for i, c := range node.children {
		if c.value > chosen.value {
			chosen = &node.children[i]
		}
	}
	return game.ComputeMove(chosen.gameData.LastMove)
}

func MonteCarloGamePlay() {
	state := MCTS{0.0, 0, &game.DataGame{Self: true, UBoard: game.EmptyUltimateBoard(),
		BoardResult: game.EmptyBoard(), LastMove: game.MoveCoordinate{BoardCoordinate: -1, Coordinate: -1}},
		make([]MCTS, 0), nil}
	turn := 0
	for {
		turn ++
		var opponentRow, opponentCol int
		_, _ = fmt.Scan(&opponentRow, &opponentCol)

		var validActionCount int
		_, _ = fmt.Scan(&validActionCount)

		for i := 0; i < validActionCount; i++ {
			var row, col int
			_, _ = fmt.Scan(&row, &col)
		}

		if opponentRow == -1 && opponentCol == -1 {
			state.gameData.Self = true
		} else {
			state.gameData.Self = false
			afterMove(&state, opponentCol, opponentRow)
		}

		var durationLimit time.Duration
		if turn == 1 {
			durationLimit, _ = time.ParseDuration("0.99s")
		} else {
			durationLimit, _ = time.ParseDuration("99ms")
		}
		runMCTS(&state, time.Now().Add(durationLimit))
		x, y := chooseBestMove(&state)
		afterMove(&state, x, y)
		fmt.Println(y, x)
	}
}


func selecting(node *MCTS) *MCTS {

	if len(node.children) == 0 {
		if node.nbVisited == 0 {
			return node
		} else {
			nextPossibilites := game.FindNextPossibilities(node.gameData)
			for i := range nextPossibilites {
				// check if terminated node
				node.children = append(node.children,
				MCTS{0.0, 0, nextPossibilites[i], make([]MCTS, 0),
				node})
			}
			return &node.children[0]
			
		}
	}

	for _, c := range node.children {
		if c.nbVisited == 0 {
			return &c
		}
	}

	chosen := &node.children[0]
	currentUCT := uct(chosen)

	for i, c := range node.children {
		if currentUCT < uct(&c) {
			chosen = &node.children[i]
			currentUCT = uct(&c)
		}
	}

	return selecting(chosen)
}



func rollout(node *MCTS) float64 {
	dataGame := node.gameData
	for {
		winner := game.FindWinnerUltimate(&dataGame.UBoard)
		if winner == game.SELF {
			return 1.0
		} else if winner == game.OPPONENT {
			return -1.0
		}

		nextStates := game.FindNextPossibilities(dataGame)
		nbValues := len(nextStates)
		if nbValues == 0 {
			return 0.0
		} else {
			dataGame = nextStates[rand.Intn(len(nextStates))]
		}
	}
}