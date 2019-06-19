package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

type Result [][]int


func BToMB(b uint64) uint64 {
	return b / 1000 / 1000
}


func FindMaxWorkers(value, minWorkers, maxWorkers int, memUsage float64) int {
	// find maximum memory of compute function
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	Compute(value)
	maxMemory := float64(BToMB(m.Sys))
	runtime.GC()
	fmt.Printf("Maximum memory a function take = %v MB \n", maxMemory)

	// find maximum workers can use
	allWorkers := runtime.NumCPU()

	if allWorkers < minWorkers {
		fmt.Print("Not enough cells to running multiprocessing.")
		return -1
	}

	//maxWorkers = math.MinInt8(sum_workers, maxWorkers) -> not working?

	if allWorkers < maxWorkers{
		maxWorkers = allWorkers
	}

	nWorkers := int(math.Min(float64(memUsage / maxMemory), float64(maxWorkers)))

	fmt.Printf("So we can use %d workers at the same time. \n\n", nWorkers)

	return nWorkers
}


func main() {
	lst := [4]int{0, 1, 2, 3}
	nWorkers := FindMaxWorkers(lst[0], 2, 10, 15)

	// parallel processing using GOMAXPROCS() and WaitGroup
	runtime.GOMAXPROCS(nWorkers)
	var wg sync.WaitGroup
	wg.Add(len(lst))

	results := make([]Result, len(lst))
	for i, value := range lst {
		go func (i, value int){
			results[i] = Compute(value)
			defer wg.Done()
		} (i, value)
	}
	wg.Wait()
	fmt.Print(results)
}
