package main

type Cell struct {
	row int
	col int
	state bool
}


// find all neighbours are living
func (c * Cell) GetNeighbours(row, col int, cl CellList) (neighbours int) {
	around := [8][2]int{{-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}}
	neighbours = 0

	for _, pos := range around {
		x := row + pos[0]
		y := col + pos[1]
		if x >= 0 && y >= 0 && x < len(cl.grid) && y < len(cl.grid) {
			neighbour := cl.grid[x][y].state
			if neighbour {
				neighbours++
			}
		}
	}
	return
}


func makeCell(row int, col int, value int) (cell Cell){
	var state bool
	if value == 1 {
		state = true
	} else {
		state = false
	}
	cell = Cell{row, col, state}
	return
}
