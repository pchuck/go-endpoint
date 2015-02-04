package main

import (
	"fmt"
	"runtime"
	"sync"
)

func alpha(wg *sync.WaitGroup) {
	defer wg.Done()
	for char := 'a'; char < 'a'+26; char++ {
		fmt.Printf("%c ", char)
	}
}

func count(wg *sync.WaitGroup) {
	defer wg.Done()
	for number := 1; number < 27; number++ {
		fmt.Printf("%d ", number)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("starting routines")

	// pass a pointer to the wait group to each routine
	go alpha(&wg)
	go count(&wg)

	fmt.Println("waiting")
	wg.Wait()

	fmt.Println("\nterminating")
}
