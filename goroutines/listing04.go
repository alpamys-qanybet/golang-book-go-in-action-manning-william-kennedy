package main

import (
  "fmt"
  "runtime"
  "sync"
)

var wg sync.WaitGroup

func main() {
  // runtime.GOMAXPROCS(1)
  runtime.GOMAXPROCS(runtime.NumCPU())
  // runtime.GOMAXPROCS(2)

  wg.Add(2)

  fmt.Println("Create Goroutines")

  go printPrime("A")
  go printPrime("B")

  fmt.Println("Waiting to finish")
  wg.Wait()

  fmt.Println("\nTerminating program...")
}


func printPrime(prefix string) {
  defer wg.Done()

next:
  for outer := 2; outer < 5000; outer++ {
    for inner := 2; inner < outer; inner++ {
      if inner % outer == 0 {
        continue next
      }
    }
    fmt.Printf("%s:%d\n", prefix, outer)
  }
  fmt.Println("Completed", prefix)
}
