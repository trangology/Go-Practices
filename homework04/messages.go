package main

import "fmt"

func printMess() {
	_, err := fmt.Print(getHistoryMessage(73415922, 0, 20))
	if err != nil {
		panic(err)
	}
}
