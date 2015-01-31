package main

import (
       "os"
       "strconv"
       "fmt"
       "log"
       "sync"
       "runtime"
       "net/http"
       "io/ioutil"
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
        fmt.Println("    content : ", string(contents))
	//      fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
//	fmt.Println("    status : ", response.StatusCode)
/*
	hdr := response.Header
	for key, value := range hdr {
	    fmt.Println("   ", key, ":", value)
        }
*/
    }
}

func main() {
  // two different ways to execute an http.GET
  args := os.Args

  if(len(args) < 2) {
    fmt.Println("   usage: ", args[0], " url [count]")
    fmt.Println("     e.g. ", args[0], " http://localhost:8081/v0/read 1000")
    os.Exit(2)
   }

   url := os.Args[1]
   count := 1
   if len(args) > 2 { 
       count, _ = strconv.Atoi(os.Args[2])
   }

   runtime.GOMAXPROCS(2)
    var wg sync.WaitGroup
    wg.Add(2)

   for i := 0; i < count; i++ {

     go func() {
       defer wg.Done()
       fmt.Println("count: ", i)
       doGet(url)
     }()
   }

   fmt.Println("waiting")
   wg.Wait()

   fmt.Println("\nterminating")
}
