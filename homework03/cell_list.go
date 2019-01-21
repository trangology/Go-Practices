package main

import (
	"math/rand"
)

type CellList struct {
	grid [][]Cell
	width int
	height int
}


func (cl CellList) Update() [][]Cell {
	for i, row := range cl.grid{
		for j, cell := range row{
			neighbors := cell.GetNeighbors(i, j, cl)
			if cell.state && (neighbors < 2 || neighbors > 3) {
				cell.state = false
			}
			if !cell.state && neighbors > 3{
				cell.state = true
			}
		}
	}
	return cl.grid
}


func randomGrid(nrows, ncols int) [][]Cell {
	// make empty 2d array of Cell struct
	grid := make([][]Cell, nrows)
	for i := 0; i < nrows; i++ {
		grid[i] = make([]Cell, ncols)
	}

	// fill "grid" array
	for i := 0; i < nrows; i++{
		for j := 0; j < ncols; j++{
			grid[i][j] = makeCell(i, j, rand.Intn(2))
		}
	}
	return grid
}

func createCellList(width int, height int, cellSize int) (cList CellList) {
	cellWidth := width / cellSize
	cellHeight := height / cellSize

	grid := randomGrid(cellHeight, cellWidth)
	cList = CellList{grid, width, height}
	return
}
