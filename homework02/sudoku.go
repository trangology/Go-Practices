package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func readSudoku(filename string) ([][]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	grid := group(filter(data), 9)
	return grid, nil
}

func filter(values []byte) []byte {
	filtered_values := make([]byte, 0)
	for _, v := range values {
		if (v >= '1' && v <= '9') || v == '.' {
			filtered_values = append(filtered_values, v)
		}
	}
	return filtered_values
}

func display(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
}

func group(values []byte, n int) [][]byte {
	times := len(values) / n
	result := make([][]byte, 0)
	count := 0

	for i := 0; i < times; i++ {
		var row []byte
		for j := 0; j < times; j++ {
			row = append(row, values[count])
			count++
		}
		result = append(result, row)
	}
	return result
}

func getRow(grid [][]byte, row int) []byte {
	var valueRow []byte
	for i := 0; i < len(grid); i++ {
		valueRow = append(valueRow, grid[row][i])
	}
	return valueRow
}

func getCol(grid [][]byte, col int) []byte {
	var valueCol []byte
	for i := 0; i < len(grid); i++ {
		valueCol = append(valueCol, grid[i][col])
	}
	return valueCol
}

func getBlock(grid [][]byte, row int, col int) []byte {
	indexRow := row / 3 * 3
	indexCol := col / 3 * 3
	var block []byte

	for i := indexRow; i < indexRow+3; i++ {
		for j := indexCol; j < indexCol+3; j++ {
			block = append(block, grid[i][j])
		}
	}
	return block
}

func findEmptyPosition(grid [][]byte) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func contains(values []byte, search byte) bool {
	for _, v := range values {
		if v == search {
			return true
		}
	}
	return false
}

func findPossibleValues(grid [][]byte, row int, col int) []byte {
	valuesRow := getRow(grid, row)
	valuesCol := getCol(grid, col)
	valuesBlock := getBlock(grid, row, col)
	check := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for _, value := range valuesRow {
		for index, item := range check {
			if value == item {
				check[index] = '0'
			}
		}
	}

	for _, value := range valuesCol {
		for index, item := range check {
			if value == item {
				check[index] = '0'
			}
		}
	}

	for _, value := range valuesBlock {
		for index, item := range check {
			if value == item {
				check[index] = '0'
			}
		}
	}

	var result []byte
	for _, value := range check {
		if value != '0' {
			result = append(result, value)
		}
	}
	return result
}

func solve(grid [][]byte) ([][]byte, bool) {
	row, col := findEmptyPosition(grid)
	if row == -1 && col == -1 {
		return grid, true
	}
	arrPos := findPossibleValues(grid, row, col)
	for _, value := range arrPos {
		grid[row][col] = value
		_, result := solve(grid)
		if result == true {
			return solve(grid)
		}
	}
	grid[row][col] = '.'
	return grid, false
}

func checkSolution(grid [][]byte) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '.' {
				return false
			}
		}
	}
	return true
}

func main() {
	puzzles, err := filepath.Glob("puzzle*.txt")
	if err != nil {
		fmt.Printf("Could not find any puzzles")
		return
	}
	for _, fname := range puzzles {
		go func(fname string) {
			grid, _ := readSudoku(fname)
			solution, _ := solve(grid)
			fmt.Println("Solution for", fname)
			display(solution)
		}(fname)
	}
	var input string
	fmt.Scanln(&input)
}
