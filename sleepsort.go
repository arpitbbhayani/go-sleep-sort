package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// prints a number of sleeping for n seconds
func sleepAndPrint(x int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Sleeping for time proportional to value
	time.Sleep(time.Duration(x) * time.Millisecond)

	// Printing the value
	fmt.Println(x)
}

// Sorts given integer slice using sleep sort
func Sort(numbers []int) {
	var wg sync.WaitGroup

	// Creating wait group that waits of len(numbers) of go routines to finish
	wg.Add(len(numbers))

	for _, x := range numbers {
		// Spinning a Go routine
		go sleepAndPrint(x, &wg)
	}

	// Waiting for all go routines to finish
	wg.Wait()
}

func main() {
	args := os.Args[1:]
	var numbers = []int{}
	for _, x := range args {
		xi, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, xi)
	}
	Sort(numbers)
}
