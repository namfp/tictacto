package main

type GameState interface {
		self() bool
		result() float64
		nextMoves() []GameState
		gameFinished() bool
		minimax() *GameState
	}

type gameState struct {
	self bool
	result float32
	board Board
	bestMove *gameState
}

type Board = [9]int

//
//func minimaxGeneric(s GameState) GameState {
//	result := s.result()
//	allMoves := s.nextMoves()
//	if s.gameFinished() {
//
//	}
//
//
//}