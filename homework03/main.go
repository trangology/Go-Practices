package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"time"
)

var game = [3][3]int{{30, 30, 1}, {20, 10, 1}, {20, 20, 0}}


func main() {
	var cList CellList // cList = list of cells
	var elapsed time.Duration

	// 4 lines below will be use to improve code, but not today
	/**
	width, height := screen.Size()
	width /= 2
	height /= 2
	**/

	fmt.Println("We have 3 games. First game will start after 5s... Please wait!")
	for i := 0; i < 5; i++{
		fmt.Printf("%d...\t", i+1)
		time.Sleep(1000*time.Millisecond)
	}

	for index, value := range game {
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
		if value[2] == 0 {
			fmt.Print("Can not run this game. Size of cell must > 0")
		}
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

			// elapsed has unit is nanosecond and we have ~ 5 seconds to run game
			if elapsed > 5000000000 {
				break
			}
		}
		screen.Clear()
		time.Sleep(3000*time.Millisecond)

		// need to find why below command not exactly working when game is running
		fmt.Printf("Game %d runs in %v", index+1, elapsed)
	}
	fmt.Print("The end.")
}