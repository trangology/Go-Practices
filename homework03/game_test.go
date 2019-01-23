package main

import (
	//"fmt"
	"io/ioutil"
	"strings"

	"testing"
)

var steps [][]string
var grid [][]Cell


func check(e error)  {
	if e != nil{
		panic(e)
	}
}


func MakeGridFromFile() {
	g, err := ioutil.ReadFile("./grid.txt")
	check(err)
	index := 0
	grid = make([][]Cell, 6)
	for i := 0; i < 6; i++ {
		grid[i] = make([]Cell, 8)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 8; j++{
			state := int(g[index]) - 48
			cell := makeCell(i, j, state)
			grid[i][j] = cell
			index++
		}
		index += 2  // has 2 symbol not 0 or 1
	}
	return
}


func MakeStepsFromFile(content []string) {
	index := 0
	for i := 0; i < 2; i++{
		var step []string
		for j := 0; j < 6; j++{
			step = append(step, content[index])
			index++
		}
		// increase index to avoid string = "\r" in "content" variable
		index++
		steps = append(steps, step)
	}
	return
}


func ReadAllFile() {
	MakeGridFromFile()
	f, _ := ioutil.ReadFile("./steps.txt")
	content := strings.Split(string(f), "\n")
	MakeStepsFromFile(content)
	return
}

func AsserEq(cl [][]Cell, step int) bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 8; j++{
			for pos, chr := range steps[step][i]{
				if pos == j {
					if cl[i][j].state && chr - 48 == 0{
						return false
					}
					if !cl[i][j].state && chr - 48 == 1{
						return false
					}
				}
			}
		}
	}
	return true
}


func TestCell_GetNeighbours(t *testing.T) {
	ReadAllFile()
	cellList := CellList{grid, 8, 6}
	for step := 0; step < len(steps); step++{
		newClist := cellList.Update()
		if !AsserEq(newClist, step){
			t.Fatalf("Got incorrect cell list when running at step %d", step + 1)
		}
	}
}