package main

type CellData struct {
	Number uint
}

type BoardData = [][]CellData

type CellPos struct {
	RowIdx uint
	ColIdx uint
}
