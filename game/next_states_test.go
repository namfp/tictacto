package game

import "testing"

func TestConversionCoordinate(t *testing.T) {
	c := toBoardCoordinate(5, 7)
	if c !=( MoveCoordinate{7, 5}) {
		t.Errorf("%v not correct", c)
	}
}

type Cd struct {
	x int
	y int
}

func TestFindNextPossibilities1(t *testing.T) {
	gameData := DataGame{true, EmptyUltimateBoard(),
		EmptyBoard(), MoveCoordinate{-1, -1}}
	played := [...]Cd{
		Cd{4,5 },
		Cd{5, 8},
		Cd{8, 6},
		Cd{8, 0},
		Cd{8, 2},
		Cd{8, 8},
		Cd{6, 6},
		Cd{0, 2},
		Cd{1, 7},
		Cd{5, 5},
		Cd{6, 7},
		Cd{1, 3},
		Cd{5, 1},
		Cd{8, 3},
		Cd{7, 2},
		Cd{5, 6},
		Cd{7, 1},
		Cd{3, 5},
		Cd{1, 8},
		Cd{5, 7}}


	for _, c := range played {
		Move(&gameData, c.y, c.x)
	}
	tests := FindNextPossibilities(&gameData)
	if len(tests) != 8 {
		t.Errorf("lenfth must be 8, got %v", len(tests))
	}
}


func TestFindNextPossibilities2(t *testing.T) {

	gameData := DataGame{true, EmptyUltimateBoard(),
		EmptyBoard(), MoveCoordinate{-1, -1}}
	allPossibilites := FindNextPossibilities(&gameData)
	if len(allPossibilites) != 81 {
		t.Errorf("allPossibilites should be 81, got %v", len(allPossibilites))
	}
}



func TestMoveUltimate2(t *testing.T) {
	gameData :=  DataGame{true, EmptyUltimateBoard(),
		EmptyBoard(), MoveCoordinate{-1, -1}}
	Move(&gameData, 4, 6)
	compare := MoveCoordinate{7, 1}
	if gameData.LastMove != compare {
		t.Errorf("lastMove should be %v, got %v", compare, gameData.LastMove)
	}

	if gameData.Self {
		t.Errorf("self should be false")
	}

}