package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func doGetShort(i int, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("  ", i, ": ", string(contents))
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("   usage: ", args[0], " url [count]")
		fmt.Println("     e.g. ", args[0],
			" http://localhost:8081/v0/read 1000")
		os.Exit(2)
	}

	url := os.Args[1]
	count := 1
	if len(args) > 2 {
		count, _ = strconv.Atoi(os.Args[2])
	}

	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go doGetShort(i, url, &wg)
	}

	fmt.Println("waiting... ")
	wg.Wait()

	fmt.Println("\nterminating")
}
