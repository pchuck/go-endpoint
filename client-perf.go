package main

import (
       "fmt"
       "log"
       "net/http"
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
      fmt.Println("content: ", contents)
	//      fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
	fmt.Println("   ", response.StatusCode)
	hdr := response.Header
	for key, value := range hdr {
	    fmt.Println("   ", key, ":", value)
        }
    }
}

func main() {
  // two different ways to execute an http.GET
  doGet("http://localhost:8081/v0/read")
}
