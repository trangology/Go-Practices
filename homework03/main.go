package main

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)


func main()  {
	cList := createCellList(120, 30, 1)   // cList = list of cells
	encoding.Register()
	for {
		// init screen
		s, _ := tcell.NewScreen()		// s == screen
		s.Resize(cList.width, cList.height, cList.width, cList.height)
		if err := s.Init(); err != nil {
			panic(err)
		}

		defer s.Fini()

		// init board of game into screen
		b := Board{cList, cList.width, cList.height}
		b.Init(s, cList)

		cList.Update()

		// show screen
		s.Show()
	}
}
