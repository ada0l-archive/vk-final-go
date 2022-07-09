package main

import (
	"flag"
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
	parsedDurations := internal.Load(*file)

	// waiter
	var wg sync.WaitGroup
	wg.Add(len(parsedDurations))

	for i, duration := range parsedDurations {
		job := func(i int, duration time.Duration, wg *sync.WaitGroup) {
			log.Print("Job started ", i+1, ", Duration: ", duration, "\n")
			internal.DoSomething(duration)
			log.Print("Job ended ", i+1, ", Duration: ", duration, "\n")
			wg.Done()
		}
		go job(i, duration, &wg)
	}
	wg.Wait()
}
