package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

var (
    counter int64
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
    fmt.Println("Final counter:", counter)
}

func incCounter(id int) {
  defer wg.Done()

  for i := 0; i < 2; i++ {
    atomic.AddInt64(&counter, 1)

    runtime.Gosched()
  }
}
