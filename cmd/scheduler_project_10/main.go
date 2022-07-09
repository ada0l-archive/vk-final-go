package main

import (
	"flag"
	"log"
	"scheduler_project/internal"
)

var (
	file = flag.String("f", "", "file-name")
)

func main() {
	flag.Parse()
	parsedDurations := internal.Load(*file)

	for i, duration := range parsedDurations {
		log.Print("Job started ", i+1, ", Duration: ", duration, "\n")
		internal.DoSomething(duration)
		log.Print("Job ended ", i+1, ", Duration: ", duration, "\n")
	}
}
