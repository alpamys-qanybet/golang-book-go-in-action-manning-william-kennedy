package main

import (
  "fmt"
  "runtime"
  "sync"
)

func main() {
  runtime.GOMAXPROCS(1)

  var wg sync.WaitGroup
  wg.Add(2)

  fmt.Println("Start Goroutines")

  go func() {
    defer wg.Done()

    for i := 0; i<3; i++ {
      for c := 'a'; c<'a'+26; c++ {
        fmt.Printf("%c ", c)
      }
    }
  }()

  go func() {
    defer wg.Done()

    for i := 0; i<3; i++ {
      for c := 'A'; c<'A'+26; c++ {
        fmt.Printf("%c ", c)
      }
    }
  }()

  fmt.Println("Waiting to finish")
  wg.Wait()

  fmt.Println("\nTerminating program...")
}
