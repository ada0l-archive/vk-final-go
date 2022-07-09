package main

import (
	"flag"
	"fmt"
	"log"
	"scheduler_project/internal"
	"sync"
	"time"
)

var (
	file = flag.String("f", "", "file-name")
)

func main() {
	flag.Parse()
	var maxCountOfThreads int
	fmt.Print("Write max count of threads: ")
	_, err := fmt.Scanf("%d", &maxCountOfThreads)
	if err != nil {
		fmt.Println(err)
		panic("Not valid max count of threads")
	}
	parsedDurations := internal.Load(*file)

	// waiter
	var wg sync.WaitGroup
	wg.Add(len(parsedDurations))

	// chain
	ch := make(chan int, maxCountOfThreads)

	for i, duration := range parsedDurations {
		job := func(i int, duration time.Duration, ch chan int, wg *sync.WaitGroup) {
			log.Print("Job started ", i+1, ", Duration: ", duration, "\n")
			internal.DoSomething(duration)
			log.Print("Job ended ", i+1, ", Duration: ", duration, "\n")
			<-ch
			wg.Done()
		}
		ch <- 1
		go job(i, duration, ch, &wg)
	}
	wg.Wait()
}
