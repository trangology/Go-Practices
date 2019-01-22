package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func check(e error)  {
	if e != nil{
		panic(e)
	}
}

func TestCell_GetNeighbours(t *testing.T) {
	g, err := ioutil.ReadFile("./grid.txt")
	check(err)
	fmt.Print(len(g))
}