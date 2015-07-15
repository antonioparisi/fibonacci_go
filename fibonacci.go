package main

import (
  "fmt"
  "net/http"
  "strconv"
  "os"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Call /fibonacci?index={number}")
  })

  http.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
    i := fib_generator()

    index, _ := strconv.Atoi(r.FormValue("index"))
    var result uint64

    for n := 0; n <= index; n++ {
      result = <- i
    }

    fmt.Fprintf(w, "%d", result)
  })

  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    panic(err)
  }
}

func fib_generator() chan uint64 {
  c := make(chan uint64)
  var i, j uint64

  go func() {
    for i, j = 0, 1; ; i, j = i+j,i {
        c <- i
    }
  }()

  return c
}
