package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func doGet(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("    The calculated length is:",
			len(string(contents)),
			"for the url:", url)
		fmt.Println("    content : ", string(contents))
		fmt.Println("    status : ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("   usage: ", args[0], " url")
		fmt.Println("     e.g. ", args[0],
			" http://localhost:8081/v0/read")
		os.Exit(2)
	}

	url := os.Args[1]
	doGet(url)
}
