package types

type CellData struct {
	Number   int
	Editable bool
}

type BoardData = [][]CellData

type CellPos struct {
	RowIdx int
	ColIdx int
}

type Difficulty int
