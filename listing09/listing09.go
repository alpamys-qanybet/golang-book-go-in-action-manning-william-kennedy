package main

import (
  "fmt"
  "runtime"
  "sync"
)

var (
  counter int
  wg sync.WaitGroup
)

func main() {
  wg.Add(8)

  go incCounter(1)
  go incCounter(2)
  go incCounter(3)
  go incCounter(4)
  go incCounter(5)
  go incCounter(6)
  go incCounter(7)
  go incCounter(8)

  wg.Wait()
  fmt.Println("Final counter: ", counter)
}

func incCounter(id int) {
  fmt.Println("incCounter ", id)
  defer wg.Done()

  for i := 0; i < 2; i++ {
    v := counter

    runtime.Gosched() // Yield the thread and be placed back in queue. used to swap goroutine
    v++

    counter = v
  }
}


/*
The counter variable is read and written to four times, twice by each goroutine,
but the value of the counter variable when the program terminates is 2.
*/
