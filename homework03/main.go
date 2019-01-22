package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"time"
)

var game = [3][3]int{{30, 30, 1}, {40, 20, 2}, {20, 20, 1}}


func main()  {
	var cList CellList  // cList = list of cells
	var elapsed time.Duration

	// 4 lines below will be use to improve code, but not today
	/**
	width, height := screen.Size()
	width /= 2
	height /= 2
	**/

	for index, value := range game{
		// init screen
		screen, _ := tcell.NewScreen()
		if err := screen.Init(); err != nil {
			panic(err)
		}
		defer screen.Fini()

		// set color for background and foreground
		screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorWhite))
		screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite))

		// create new list of cells for new game
		cList = createCellList(value[0], value[1], value[2])

		startTime := time.Now()
		for {
			// init board of game into screen

			b := Board{cList, cList.width, cList.height}

			b.Init(screen, cList)
			cList.Update()

			// show screen
			screen.Sync()


			// if game can not stop, we need to stop loop to avoid case too many time and exit code 2
			endTime := time.Now()
			elapsed = endTime.Sub(startTime)

			// elapsed has unit is nanosecond and we have 5 seconds to run game
			if elapsed > 5000000000 {
				break
			}
		}

		// need to find why below command not working
		if index < 2 {
			fmt.Printf("Game %d run in %v. Loading game %d...\t", index+1, elapsed, index+2)
		}
	}
	fmt.Print("The end.")
}
