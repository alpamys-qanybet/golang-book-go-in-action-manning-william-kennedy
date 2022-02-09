package main

import (
    "fmt"
    "runtime"
    "sync"
)

var (
    counter int
    wg sync.WaitGroup

    mutex sync.Mutex
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
        mutex.Lock()
        { // The use of the curly brackets is just to make the critical section easier to see; they’re not necessary.
            v := counter

            runtime.Gosched() // When the thread is yielded the scheduler assigns the same goroutine to continue running
            v++
            counter = v
        }
        mutex.Unlock()
    }
}
