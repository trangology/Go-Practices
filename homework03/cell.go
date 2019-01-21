package main

type Cell struct {
	row int
	col int
	state bool
}

func (c * Cell) GetNeighbors(x, y int, cl CellList) (neighbors int) {
	around := [8][2]int{{-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}}
	neighbors = 0

	for _, pos := range around {
		x += pos[0]
		y += pos[1]
		if x >= 0 && y >= 0 && x < len(cl.grid) && y < len(cl.grid) {
			if c.state {
				neighbors++
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
