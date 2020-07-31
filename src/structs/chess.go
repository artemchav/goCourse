package structs

//Можно обойтись и одним типом, конечно, но для понятности их два - ходы описываются
//смещением по x & y, а Coords - точка на доске
type Coords [2]int
type Offset [2]int
type ChessBoard [8][8]*Figure //for further figures positions

type Figure struct {
	position  Offset
	movements []Offset
}

func (*ChessBoard) NewChessKnight(x, y int) *Figure {
	return &Figure{
		position: Offset{x, y},
		movements: []Offset{
			{-1, 2},
			{-1, -2},
			{1, -2},
			{1, 2},
			{-2, 1},
			{-2, -1},
			{2, 1},
			{2, -1},
		},
	}
}

func (f *Figure) SetPosition(x, y int) {
	f.position = Offset{x, y}
}

func (f *Figure) GetAvailableTurns() []Coords {
	availableTurns := make([]Coords, 0)
	for _, move := range f.movements {
		x := f.position[0] + move[0]
		y := f.position[1] + move[1]
		if x < 0 || y < 0 {
			continue
		}
		availableTurns = append(availableTurns, Coords{x, y})
	}

	return availableTurns
}
