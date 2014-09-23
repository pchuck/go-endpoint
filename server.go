package main

import (
    "fmt"
    "net/http"
    "time"
)

type Hello struct{}
type timeHandler struct {
  zone *time.Location
}

func (h Hello) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "hello")
}

// from http://www.alexedwards.net/blog/a-recap-of-request-handling
func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().In(th.zone).Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}

func newTimeHandler(name string) *timeHandler {
  return &timeHandler{zone: time.FixedZone(name, 0)}
}

func main() {
    var h Hello
    mux := http.NewServeMux()
    mux.Handle("/est", newTimeHandler("EST"))

    http.Handle("/hello", h)
    http.Handle("/time", newTimeHandler("MST"))

    http.ListenAndServe("localhost:4000", nil)
}
