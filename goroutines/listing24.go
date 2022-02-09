package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    numberOfGoroutines = 4
    taskLoad = 10
)

var wg sync.WaitGroup

func init() {
    rand.Seed(time.Now().Unix())
}

func main() {
    tasks := make(chan string, taskLoad)

    wg.Add(numberOfGoroutines)
    for i := 1; i <= numberOfGoroutines; i++ {
        go worker(tasks, i)
    }

    for j := 1; j <= taskLoad; j++ {
        tasks <- fmt.Sprintf("Task : %d", j)
    }

    close(tasks) // When a channel is closed, goroutines can still perform receives on the channel but can no longer send on the channel.

    wg.Wait()
}

func worker(tasks chan string, worker int) {
    defer wg.Done()

    for {
        task, ok := <- tasks
        if !ok {
            fmt.Printf("Worker: %d : Shutting Down\n", worker)
            return
        }

        fmt.Printf("Worker: %d : Started %s\n", worker, task)
        sleep := rand.Int63n(100)
        time.Sleep(time.Duration(sleep) * time.Millisecond)
        fmt.Printf("Worker: %d : Completed %s\n", worker, task)
    }
}
