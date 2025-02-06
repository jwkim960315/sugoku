package main

type CellData struct {
	Number int
}

type BoardData = [][]CellData

type CellPos struct {
	RowIdx int
	ColIdx int
}

type Difficulty int