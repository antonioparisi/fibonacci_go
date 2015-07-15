package main

import "fmt"
import "net/http"
import "strconv"

func main() {
  http.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
    i := fib_generator()

    index, _ := strconv.Atoi(r.FormValue("index"))
    var result uint64

    for n := 0; n <= index; n++ {
      result = <- i
    }

    fmt.Fprintf(w, "%d", result)
  })

  http.ListenAndServe(":8080", nil)
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
