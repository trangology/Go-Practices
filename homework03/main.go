package main

import (
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	var elapsed time.Duration

	// init screen
	screen, _ := tcell.NewConsoleScreen()
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	// width, height := screen.Size()
	// width /= 2
	// height /= 2

	// set color for background and foreground
	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorWhite))
	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack))

	// create new list of cells for new game
	cList := createCellList(10, 15, 1)

	startTime := time.Now()
	for {
		// init board of game into screen
		b := Board{cList, cList.width, cList.height}

		b.Init(screen, cList)
		cList.Update()

		// show screen
		screen.Sync()

		// if game can not stop, we need to stop loop
		// to avoid case too many time and exit code 2
		endTime := time.Now()
		elapsed = endTime.Sub(startTime)

		// elapsed has unit is nanosecond
		// and we have ~ 5 seconds to run game
		if elapsed > 5000000000 {
			break
		}
	}
	screen.Clear()
	time.Sleep(3000 * time.Millisecond)
}
